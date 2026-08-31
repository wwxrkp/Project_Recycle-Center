package entity

import "time"

// MaterialAssessmentItem links Ping's seller to Filme's quality workflow.
type MaterialAssessmentItem struct {
	AssessmentItemID  int                `gorm:"column:assessment_item_id;type:integer;primaryKey;autoIncrement" json:"assessment_item_id"`
	Quantity          float64            `gorm:"column:quantity;type:numeric(14,3);not null" json:"quantity"`
	SubmittedAt       time.Time          `gorm:"column:submitted_at;not null" json:"submitted_at"`
	Status            string             `gorm:"column:status;type:text;not null" json:"status"`
	SellerCode        string             `gorm:"column:seller_code;type:text;not null;index" json:"seller_code"`
	MaterialID        string             `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	Seller            *Seller            `gorm:"foreignKey:SellerCode;references:SellerCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"seller,omitempty"`
	Material          *Material          `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
	QualityAssessment *QualityAssessment `gorm:"foreignKey:AssessmentItemID;references:AssessmentItemID" json:"quality_assessment,omitempty"`
}

func (MaterialAssessmentItem) TableName() string { return "material_assessment_items" }

type QualityAssessment struct {
	AssessmentID     int                     `gorm:"column:assessment_id;type:integer;primaryKey;autoIncrement" json:"assessment_id"`
	AssessedGrade    string                  `gorm:"column:assessed_grade;type:text;not null" json:"assessed_grade"`
	CleanlinessLevel string                  `gorm:"column:cleanliness_level;type:text;not null" json:"cleanliness_level"`
	Result           string                  `gorm:"column:result;type:text;not null" json:"result"`
	Detail           *string                 `gorm:"column:detail;type:text" json:"detail,omitempty"`
	AssessedQuantity float64                 `gorm:"column:assessed_quantity;type:numeric(14,3);not null" json:"assessed_quantity"`
	AssessedAt       time.Time               `gorm:"column:assessed_at;not null" json:"assessed_at"`
	EmployeeID       string                  `gorm:"column:employee_id;type:text;not null;index" json:"employee_id"`
	AssessmentItemID int                     `gorm:"column:assessment_item_id;type:integer;not null;uniqueIndex" json:"assessment_item_id"`
	Employee         *User                   `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
	AssessmentItem   *MaterialAssessmentItem `gorm:"foreignKey:AssessmentItemID;references:AssessmentItemID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"assessment_item,omitempty"`
	ScrapPurchase    *ScrapPurchaseItem      `gorm:"foreignKey:AssessmentID;references:AssessmentID" json:"scrap_purchase,omitempty"`
}

func (QualityAssessment) TableName() string { return "quality_assessments" }

// ScrapPurchaseItem is the single purchase record shared by Ping and Filme.
type ScrapPurchaseItem struct {
	PurchaseID   string                `gorm:"column:purchase_id;type:text;primaryKey" json:"purchase_id"`
	PaymentID    *string               `gorm:"column:payment_id;type:text" json:"payment_id,omitempty"`
	PurchaseDate time.Time             `gorm:"column:purchase_date;not null" json:"purchase_date"`
	SellerCode   string                `gorm:"column:seller_code;type:text;not null;index" json:"seller_code"`
	MaterialID   string                `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	AssessmentID int                   `gorm:"column:assessment_id;type:integer;not null;uniqueIndex" json:"assessment_id"`
	EmployeeID   string                `gorm:"column:employee_id;type:text;not null;index" json:"employee_id"`
	Weight       float64               `gorm:"column:weight;type:numeric(14,3);not null" json:"weight"`
	PricePerKg   float64               `gorm:"column:price_per_kg;type:numeric(12,2);not null" json:"price_per_kg"`
	TotalAmount  float64               `gorm:"column:total_amount;type:numeric(14,2);not null" json:"total_amount"`
	Seller       *Seller               `gorm:"foreignKey:SellerCode;references:SellerCode;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"seller,omitempty"`
	Material     *Material             `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
	Assessment   *QualityAssessment    `gorm:"foreignKey:AssessmentID;references:AssessmentID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"assessment,omitempty"`
	Employee     *PurchasingStaff      `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
	PendingItem  *PendingWarehouseItem `gorm:"foreignKey:PurchaseID;references:PurchaseID" json:"pending_item,omitempty"`
}

func (ScrapPurchaseItem) TableName() string { return "scrap_purchase_items" }

type PendingWarehouseItem struct {
	PendingID          int                 `gorm:"column:pending_id;type:integer;primaryKey;autoIncrement" json:"pending_id"`
	Quantity           float64             `gorm:"column:quantity;type:numeric(14,3);not null" json:"quantity"`
	AssessedGrade      string              `gorm:"column:assessed_grade;type:text;not null" json:"assessed_grade"`
	AssessedByUserID   string              `gorm:"column:assessed_by_user_id;type:text;not null;index" json:"assessed_by_user_id"`
	TransferredDate    time.Time           `gorm:"column:transferred_date;not null" json:"transferred_date"`
	StockRouteType     string              `gorm:"column:stock_route_type;type:text;not null" json:"stock_route_type"`
	PurchaseID         string              `gorm:"column:purchase_id;type:text;not null;uniqueIndex" json:"purchase_id"`
	ReceivingStatus    string              `gorm:"column:receiving_status;type:text;not null" json:"receiving_status"`
	MaterialID         string              `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	AssessedBy         *User               `gorm:"foreignKey:AssessedByUserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"assessed_by,omitempty"`
	Purchase           *ScrapPurchaseItem  `gorm:"foreignKey:PurchaseID;references:PurchaseID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"purchase,omitempty"`
	Material           *Material           `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
	ReceiveTransaction *ReceiveTransaction `gorm:"foreignKey:PendingID;references:PendingID" json:"receive_transaction,omitempty"`
}

func (PendingWarehouseItem) TableName() string { return "pending_warehouse_items" }
