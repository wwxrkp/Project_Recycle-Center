package entity

import "time"

type Factory struct {
	FactoryID     string `gorm:"column:factory_id;type:text;primaryKey" json:"factory_id"`
	CompanyName   string `gorm:"column:company_name;type:text;not null" json:"company_name"`
	ContactPerson string `gorm:"column:contact_person;type:text;not null" json:"contact_person"`
	Phone         string `gorm:"column:phone;type:text;not null" json:"phone"`
	Address       string `gorm:"column:address;type:text;not null" json:"address"`
}

func (Factory) TableName() string { return "factories" }

type SalesContractStatus string

const (
	SalesContractActive    SalesContractStatus = "active"
	SalesContractExpired   SalesContractStatus = "expired"
	SalesContractSuspended SalesContractStatus = "suspended"
)

type SalesContract struct {
	ContractID   string              `gorm:"column:contract_id;type:text;primaryKey" json:"contract_id"`
	FactoryID    string              `gorm:"column:factory_id;type:text;not null;index" json:"factory_id"`
	MaterialID   string              `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	ContractDate time.Time           `gorm:"column:contract_date;not null" json:"contract_date"`
	Quantity     float64             `gorm:"column:quantity;type:numeric(14,3);not null" json:"quantity"`
	UnitPrice    float64             `gorm:"column:unit_price;type:numeric(12,2);not null" json:"unit_price"`
	ValidFrom    time.Time           `gorm:"column:valid_from;type:date;not null" json:"valid_from"`
	ValidTo      time.Time           `gorm:"column:valid_to;type:date;not null" json:"valid_to"`
	Status       SalesContractStatus `gorm:"column:status;type:text;not null" json:"status"`
	Factory      *Factory            `gorm:"foreignKey:FactoryID;references:FactoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"factory,omitempty"`
	Material     *Material           `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
}

func (SalesContract) TableName() string { return "sales_contracts" }

type ScrapSaleToFactory struct {
	SaleID      string         `gorm:"column:sale_id;type:text;primaryKey" json:"sale_id"`
	FactoryID   string         `gorm:"column:factory_id;type:text;not null;index" json:"factory_id"`
	ContractID  string         `gorm:"column:contract_id;type:text;not null;index" json:"contract_id"`
	MaterialID  string         `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	SaleDate    time.Time      `gorm:"column:sale_date;not null" json:"sale_date"`
	Quantity    float64        `gorm:"column:quantity;type:numeric(14,3);not null" json:"quantity"`
	UnitPrice   float64        `gorm:"column:unit_price;type:numeric(12,2);not null" json:"unit_price"`
	TotalAmount float64        `gorm:"column:total_amount;type:numeric(14,2);not null" json:"total_amount"`
	Factory     *Factory       `gorm:"foreignKey:FactoryID;references:FactoryID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"factory,omitempty"`
	Contract    *SalesContract `gorm:"foreignKey:ContractID;references:ContractID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"contract,omitempty"`
	Material    *Material      `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
}

func (ScrapSaleToFactory) TableName() string { return "scrap_sales_to_factory" }

type SalesBill struct {
	BillID        string             `gorm:"column:bill_id;type:text;primaryKey" json:"bill_id"`
	PurchaseID    string             `gorm:"column:purchase_id;type:text;not null;index" json:"purchase_id"`
	BillDate      time.Time          `gorm:"column:bill_date;not null" json:"bill_date"`
	TotalAmount   float64            `gorm:"column:total_amount;type:numeric(14,2);not null" json:"total_amount"`
	PaymentStatus string             `gorm:"column:payment_status;type:text;not null" json:"payment_status"`
	IssuedBy      string             `gorm:"column:issued_by;type:text;not null;index" json:"issued_by"`
	Purchase      *ScrapPurchaseItem `gorm:"foreignKey:PurchaseID;references:PurchaseID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"purchase,omitempty"`
	Issuer        *User              `gorm:"foreignKey:IssuedBy;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"issuer,omitempty"`
}

func (SalesBill) TableName() string { return "sales_bills" }

type ComplaintStatus string

const (
	ComplaintStatusPending  ComplaintStatus = "pending"
	ComplaintStatusApproved ComplaintStatus = "approved"
	ComplaintStatusRejected ComplaintStatus = "rejected"
)

type Complaint struct {
	ComplaintID        string              `gorm:"column:complaint_id;type:text;primaryKey" json:"complaint_id"`
	PurchaseID         *string             `gorm:"column:purchase_id;type:text;index" json:"purchase_id,omitempty"`
	ManagerID          *string             `gorm:"column:manager_id;type:text;index" json:"manager_id,omitempty"`
	FactoryID          *string             `gorm:"column:factory_id;type:text;index" json:"factory_id,omitempty"`
	EmployeeID         *string             `gorm:"column:employee_id;type:text;index" json:"employee_id,omitempty"`
	SaleID             *string             `gorm:"column:sale_id;type:text;index" json:"sale_id,omitempty"`
	ProblemDescription string              `gorm:"column:problem_description;type:text;not null" json:"problem_description"`
	ComplaintDate      time.Time           `gorm:"column:complaint_date;not null" json:"complaint_date"`
	EvidenceFile       string              `gorm:"column:evidence_file;type:text" json:"evidence_file"`
	Status             ComplaintStatus     `gorm:"column:status;type:text;not null" json:"status"`
	ReviewedBy         *string             `gorm:"column:reviewed_by;type:text;index" json:"reviewed_by,omitempty"`
	Result             *string             `gorm:"column:result;type:text" json:"result,omitempty"`
	RejectionReason    *string             `gorm:"column:rejection_reason;type:text" json:"rejection_reason,omitempty"`
	Purchase           *ScrapPurchaseItem  `gorm:"foreignKey:PurchaseID;references:PurchaseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"purchase,omitempty"`
	Manager            *Manager            `gorm:"foreignKey:ManagerID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"manager,omitempty"`
	Factory            *Factory            `gorm:"foreignKey:FactoryID;references:FactoryID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"factory,omitempty"`
	Employee           *User               `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"employee,omitempty"`
	Sale               *ScrapSaleToFactory `gorm:"foreignKey:SaleID;references:SaleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"sale,omitempty"`
	Reviewer           *User               `gorm:"foreignKey:ReviewedBy;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"reviewer,omitempty"`
}

func (Complaint) TableName() string { return "complaints" }
