package entity

import "time"

// Supplier is an organization used by Mark's procurement workflow.
// It is intentionally separate from Ping's Seller, who sells scrap to the business.
type Supplier struct {
	SupplierID  string `gorm:"column:supplier_id;type:text;primaryKey" json:"supplier_id"`
	CompanyName string `gorm:"column:company_name;type:text;not null" json:"company_name"`
	Phone       string `gorm:"column:phone;type:text" json:"phone"`
	Address     string `gorm:"column:address;type:text" json:"address"`
}

func (Supplier) TableName() string { return "suppliers" }

type PurchaseOrder struct {
	OrderID      string                  `gorm:"column:order_id;type:text;primaryKey" json:"order_id"`
	SalesStaffID string                  `gorm:"column:sales_staff_id;type:text;not null;index" json:"sales_staff_id"`
	OrderDate    time.Time               `gorm:"column:order_date;not null" json:"order_date"`
	Status       string                  `gorm:"column:status;type:text;not null;index" json:"status"`
	SalesStaff   *SalesStaff             `gorm:"foreignKey:SalesStaffID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"sales_staff,omitempty"`
	Materials    []PurchaseOrderMaterial `gorm:"foreignKey:OrderID;references:OrderID" json:"materials,omitempty"`
	Suppliers    []PurchaseOrderSupplier `gorm:"foreignKey:OrderID;references:OrderID" json:"suppliers,omitempty"`
}

func (PurchaseOrder) TableName() string { return "purchase_orders" }

type PurchaseOrderMaterial struct {
	OrderID           string         `gorm:"column:order_id;type:text;primaryKey" json:"order_id"`
	MaterialID        string         `gorm:"column:material_id;type:text;primaryKey" json:"material_id"`
	RequestedQuantity float64        `gorm:"column:requested_quantity;type:numeric(14,3);not null" json:"requested_quantity"`
	PurchaseOrder     *PurchaseOrder `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"purchase_order,omitempty"`
	Material          *Material      `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
}

func (PurchaseOrderMaterial) TableName() string { return "purchase_order_materials" }

type PurchaseOrderSupplier struct {
	OrderID       string         `gorm:"column:order_id;type:text;primaryKey" json:"order_id"`
	SupplierID    string         `gorm:"column:supplier_id;type:text;primaryKey" json:"supplier_id"`
	PurchaseOrder *PurchaseOrder `gorm:"foreignKey:OrderID;references:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"purchase_order,omitempty"`
	Supplier      *Supplier      `gorm:"foreignKey:SupplierID;references:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"supplier,omitempty"`
}

func (PurchaseOrderSupplier) TableName() string { return "purchase_order_suppliers" }
