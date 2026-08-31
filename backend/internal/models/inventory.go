package entity

import "time"

type Warehouse struct {
	WarehouseID   string        `gorm:"column:warehouse_id;type:text;primaryKey" json:"warehouse_id"`
	WarehouseName string        `gorm:"column:warehouse_name;type:text;not null" json:"warehouse_name"`
	TotalCapacity float64       `gorm:"column:total_capacity;type:numeric(14,3);not null" json:"total_capacity"`
	LastUpdated   time.Time     `gorm:"column:last_updated;not null" json:"last_updated"`
	StorageZones  []StorageZone `gorm:"foreignKey:WarehouseID;references:WarehouseID" json:"storage_zones,omitempty"`
}

func (Warehouse) TableName() string { return "warehouses" }

// StorageZone is the source of truth for current stock quantity.
type StorageZone struct {
	ZoneID         string                   `gorm:"column:zone_id;type:text;primaryKey" json:"zone_id"`
	ZoneName       string                   `gorm:"column:zone_name;type:text;not null" json:"zone_name"`
	Capacity       float64                  `gorm:"column:capacity;type:numeric(14,3);not null" json:"capacity"`
	SupportedGrade string                   `gorm:"column:supported_grade;type:text" json:"supported_grade"`
	LastUpdated    time.Time                `gorm:"column:last_updated;not null" json:"last_updated"`
	StockStatus    string                   `gorm:"column:stock_status;type:text;not null" json:"stock_status"`
	QuantityOnHand float64                  `gorm:"column:quantity_on_hand;type:numeric(14,3);not null;default:0" json:"quantity_on_hand"`
	WarehouseID    string                   `gorm:"column:warehouse_id;type:text;not null;index" json:"warehouse_id"`
	MaterialID     string                   `gorm:"column:material_id;type:text;not null;index" json:"material_id"`
	Warehouse      *Warehouse               `gorm:"foreignKey:WarehouseID;references:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"warehouse,omitempty"`
	Material       *Material                `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
	Transactions   []StockTransaction       `gorm:"foreignKey:ZoneID;references:ZoneID" json:"transactions,omitempty"`
	Adjustments    []StockAdjustmentRequest `gorm:"foreignKey:ZoneID;references:ZoneID" json:"adjustments,omitempty"`
}

func (StorageZone) TableName() string { return "storage_zones" }

type StockTransactionType string

const (
	StockTransactionReceive StockTransactionType = "receive"
	StockTransactionIssue   StockTransactionType = "issue"
	StockTransactionAdjust  StockTransactionType = "adjust"
)

type StockTransaction struct {
	TransactionID      int                  `gorm:"column:transaction_id;type:integer;primaryKey;autoIncrement" json:"transaction_id"`
	TransactionType    StockTransactionType `gorm:"column:transaction_type;type:text;not null;index" json:"transaction_type"`
	Quantity           float64              `gorm:"column:quantity;type:numeric(14,3);not null" json:"quantity"`
	TransactionDate    time.Time            `gorm:"column:transaction_date;not null" json:"transaction_date"`
	EmployeeID         string               `gorm:"column:employee_id;type:text;not null;index" json:"employee_id"`
	ZoneID             string               `gorm:"column:zone_id;type:text;not null;index" json:"zone_id"`
	Employee           *User                `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
	Zone               *StorageZone         `gorm:"foreignKey:ZoneID;references:ZoneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"zone,omitempty"`
	ReceiveTransaction *ReceiveTransaction  `gorm:"foreignKey:TransactionID;references:TransactionID" json:"receive_transaction,omitempty"`
	IssueTransaction   *IssueTransaction    `gorm:"foreignKey:TransactionID;references:TransactionID" json:"issue_transaction,omitempty"`
}

func (StockTransaction) TableName() string { return "stock_transactions" }

type ReceiveTransaction struct {
	ReceiveNo     string                `gorm:"column:receive_no;type:text;primaryKey" json:"receive_no"`
	TransactionID int                   `gorm:"column:transaction_id;type:integer;not null;uniqueIndex" json:"transaction_id"`
	PendingID     int                   `gorm:"column:pending_id;type:integer;not null;uniqueIndex" json:"pending_id"`
	Transaction   *StockTransaction     `gorm:"foreignKey:TransactionID;references:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"transaction,omitempty"`
	PendingItem   *PendingWarehouseItem `gorm:"foreignKey:PendingID;references:PendingID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"pending_item,omitempty"`
}

func (ReceiveTransaction) TableName() string { return "receive_transactions" }

type IssueTransaction struct {
	IssueNo        string            `gorm:"column:issue_no;type:text;primaryKey" json:"issue_no"`
	ReferenceNo    string            `gorm:"column:reference_no;type:text;not null" json:"reference_no"`
	RequestingUnit string            `gorm:"column:requesting_unit;type:text;not null" json:"requesting_unit"`
	TransactionID  int               `gorm:"column:transaction_id;type:integer;not null;uniqueIndex" json:"transaction_id"`
	Transaction    *StockTransaction `gorm:"foreignKey:TransactionID;references:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"transaction,omitempty"`
}

func (IssueTransaction) TableName() string { return "issue_transactions" }

type StockAdjustmentRequest struct {
	RequestNo       int                 `gorm:"column:request_no;type:integer;primaryKey;autoIncrement" json:"request_no"`
	SystemQuantity  float64             `gorm:"column:system_quantity;type:numeric(14,3);not null" json:"system_quantity"`
	CountedQuantity float64             `gorm:"column:counted_quantity;type:numeric(14,3);not null" json:"counted_quantity"`
	Description     string              `gorm:"column:description;type:text;not null" json:"description"`
	RequestDate     time.Time           `gorm:"column:request_date;not null" json:"request_date"`
	Status          string              `gorm:"column:status;type:text;not null" json:"status"`
	EmployeeID      string              `gorm:"column:employee_id;type:text;not null;index" json:"employee_id"`
	ZoneID          string              `gorm:"column:zone_id;type:text;not null;index" json:"zone_id"`
	Employee        *User               `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
	Zone            *StorageZone        `gorm:"foreignKey:ZoneID;references:ZoneID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"zone,omitempty"`
	Approval        *AdjustmentApproval `gorm:"foreignKey:RequestNo;references:RequestNo" json:"approval,omitempty"`
}

func (StockAdjustmentRequest) TableName() string { return "stock_adjustment_requests" }

type AdjustmentApproval struct {
	ApprovalID       int                     `gorm:"column:approval_id;type:integer;primaryKey;autoIncrement" json:"approval_id"`
	Decision         string                  `gorm:"column:decision;type:text;not null" json:"decision"`
	ApprovedQuantity *float64                `gorm:"column:approved_quantity;type:numeric(14,3)" json:"approved_quantity,omitempty"`
	DecisionReason   *string                 `gorm:"column:decision_reason;type:text" json:"decision_reason,omitempty"`
	ApprovedAt       time.Time               `gorm:"column:approved_at;not null" json:"approved_at"`
	RequestNo        int                     `gorm:"column:request_no;type:integer;not null;uniqueIndex" json:"request_no"`
	EmployeeID       string                  `gorm:"column:employee_id;type:text;not null;index" json:"employee_id"`
	Request          *StockAdjustmentRequest `gorm:"foreignKey:RequestNo;references:RequestNo;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"request,omitempty"`
	Employee         *User                   `gorm:"foreignKey:EmployeeID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"employee,omitempty"`
}

func (AdjustmentApproval) TableName() string { return "adjustment_approvals" }
