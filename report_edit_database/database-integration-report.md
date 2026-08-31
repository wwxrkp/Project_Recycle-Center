# รายงานการรวม Database Models ของทีม T20

## เป้าหมาย

รวม database models ของ Film, Ping และ Mark ให้ใช้ตารางกลางร่วมกัน ไม่มีชื่อตารางซ้ำ และมี foreign key เชื่อมข้ามระบบอย่างชัดเจน

## ปัญหาก่อนรวม

- `users` ถูกนิยาม 3 ครั้ง และใช้ primary key ไม่ตรงกัน (`user_id` กับ `employee_id`)
- `materials` ถูกนิยาม 2 ครั้ง และมีฟิลด์คนละชุด
- `scrap_purchase_items` ถูกนิยาม 2 ครั้ง
- Foreign keys ของชุด Ping หลายจุดมีเพียง comment แต่ยังไม่ได้ประกาศเป็น GORM association
- `cmd/main.go` ว่าง, ไม่มี `go.mod` และ config เดิมมีรหัสผ่านฐานข้อมูลใน source code

### Shared

- `users` เป็น parent identity กลางเพียงตารางเดียว
- ตารางบทบาทใช้ `user_id` เป็นทั้ง primary key และ foreign key ไปยัง `users.user_id`
- บทบาทประกอบด้วย seller, purchasing staff, customer service, manager, driver, transport supervisor และ sales staff

### Ping — ลงทะเบียนและตรวจสอบผู้ขาย

- เป็นเจ้าของ `sellers`, `registration_forms`, `factories`, `sales_contracts`, `sales_bills`, `complaints`
- `sellers.user_id` เชื่อมกับ `users.user_id`
- Registration สามารถมีอยู่ก่อนสร้างบัญชี seller และเก็บผู้ตรวจสอบผ่าน `processed_by`
- Bill และ complaint อ้างรายการรับซื้อกลางแทนการสร้างรายการรับซื้อซ้ำ

### Filme — คุณภาพและสินค้าคงคลัง

- เป็นเจ้าของ `material_types`, `materials`, quality assessment และ inventory tables
- `materials` เป็น material master กลางเพียงตารางเดียว
- เกรดและสภาพจริงของวัสดุอยู่ในขั้นตอน assessment/stock ไม่เก็บซ้ำใน material master
- `material_assessment_items.seller_code` อ้างผู้ขายของ Ping
- `scrap_purchase_items` เป็นรายการรับซื้อกลางเพียงตารางเดียว เชื่อม seller, material, assessment และ purchasing staff
- ปริมาณคงเหลือใช้ `storage_zones.quantity_on_hand` เป็น source of truth

### Mark — รถขนส่งและคำสั่งซื้อ

- เป็นเจ้าของ `trucks`, `delivery_requests`, `cancel_requests`, `suppliers`, `purchase_orders`
- Delivery และ purchase-order item อ้าง `materials.material_id` ของ Filme
- Mark อ่านยอดคงเหลือจาก storage zone ของ Filme ก่อนทำคำสั่งซื้อ ไม่สร้าง material หรือ stock table ซ้ำ
- Delivery request เชื่อม transport supervisor, truck และรายการวัสดุด้วย foreign keys

## ตารางที่ถูกรวม

| ตารางกลาง | แหล่งเดิม | ผลลัพธ์ |
| --- | --- | --- |
| `users` | Mark, Filme, Ping | รวมเป็น parent identity และแยก role profiles |
| `materials` | Mark, Filme | ใช้ material master ของ Filme และให้ Mark อ้างด้วย FK |
| `scrap_purchase_items` | Filme, Ping | รวมฟิลด์การซื้อ การประเมิน ผู้ขาย วัสดุ และพนักงาน |

หลังรวมมี 35 models และ 35 ชื่อตารางที่ไม่ซ้ำกัน

## การปรับมาตรฐานข้อมูล

- JSON fields ใช้ `snake_case` เหมือนกัน
- Password เปลี่ยนเป็น `PasswordHash` และไม่ถูกส่งออกทาง JSON
- ปริมาณใช้ PostgreSQL `numeric(14,3)`
- ค่าเงินใช้ PostgreSQL `numeric(...,2)` และอนุญาตให้ GORM บันทึก `total_amount`
- ฟิลด์วันที่ทำรายการเก็บเวลา ไม่บังคับเป็น `date` โดยไม่จำเป็น
- เพิ่ม `transaction_type` ใน stock transaction เพื่อแยก receive, issue และ adjust
- Complaint ใช้ `complaint_id` เป็น PK จึงรองรับหลาย complaint ต่อ purchase

## การตั้งค่าและการตรวจสอบ

- เพิ่ม Go module: `github.com/SA-1-69/T20/backend`
- Database config อ่าน `DATABASE_URL` เท่านั้น ไม่มีรหัสผ่านจริงใน source code
- `AutoMigrate` ใช้ canonical models ทั้ง 35 ตัว โดยแบ่ง comment ตามเจ้าของระบบ
- เพิ่ม test ตรวจว่าชื่อตารางไม่ซ้ำ
- เพิ่ม test ให้ GORM parse schema และ associations ทุก model
- `go test ./...` ผ่าน

## ข้อควรระวังเมื่อนำไปใช้กับฐานข้อมูลเดิม

`AutoMigrate` ไม่ลบหรือ rename คอลัมน์เก่าให้อัตโนมัติ หากเคยสร้างตารางจาก models เดิมแล้ว ต้องสำรองข้อมูลและเลือกอย่างใดอย่างหนึ่ง:

1. ฐานข้อมูลสำหรับ development ที่ไม่มีข้อมูลสำคัญ: สร้าง database ใหม่แล้ว migrate canonical schema
2. ฐานข้อมูลที่มีข้อมูล: เขียน versioned SQL migration สำหรับ rename/copy/drop หลังสำรองข้อมูล

ห้าม drop ตาราง production โดยไม่มี backup และ data-migration plan

