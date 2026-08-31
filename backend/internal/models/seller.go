package entity

import "time"

type AccountStatus string

const (
	AccountStatusActive    AccountStatus = "active"
	AccountStatusSuspended AccountStatus = "suspended"
	AccountStatusPending   AccountStatus = "pending"
)

// Seller is the approved seller profile owned by Ping's module.
type Seller struct {
	SellerCode       string        `gorm:"column:seller_code;type:text;primaryKey" json:"seller_code"`
	UserID           string        `gorm:"column:user_id;type:text;not null;uniqueIndex" json:"user_id"`
	NationalID       string        `gorm:"column:national_id;type:text;not null;uniqueIndex" json:"national_id"`
	Address          string        `gorm:"column:address;type:text;not null" json:"address"`
	AccountStatus    AccountStatus `gorm:"column:account_status;type:text;not null;default:'pending'" json:"account_status"`
	RegistrationDate time.Time     `gorm:"column:registration_date;type:date;not null" json:"registration_date"`
	User             User          `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"user,omitempty"`
}

func (Seller) TableName() string { return "sellers" }

// RegistrationForm exists before a seller account is approved.
type RegistrationForm struct {
	RegistrationID   string    `gorm:"column:registration_id;type:text;primaryKey" json:"registration_id"`
	NationalID       string    `gorm:"column:national_id;type:text;not null;index" json:"national_id"`
	Name             string    `gorm:"column:name;type:text;not null" json:"name"`
	Phone            string    `gorm:"column:phone;type:text;not null" json:"phone"`
	Email            string    `gorm:"column:email;type:text" json:"email"`
	Address          string    `gorm:"column:address;type:text;not null" json:"address"`
	RegistrationDate time.Time `gorm:"column:registration_date;not null" json:"registration_date"`
	IsDuplicate      bool      `gorm:"column:is_duplicate;not null;default:false" json:"is_duplicate"`
	Status           string    `gorm:"column:status;type:text;not null" json:"status"`
	ProcessedBy      *string   `gorm:"column:processed_by;type:text;index" json:"processed_by,omitempty"`
	Processor        *User     `gorm:"foreignKey:ProcessedBy;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"processor,omitempty"`
}

func (RegistrationForm) TableName() string { return "registration_forms" }
