package entity

import "time"

type UserRole string

const (
	RoleSeller              UserRole = "seller"
	RolePurchasingStaff     UserRole = "purchasing_staff"
	RoleCustomerService     UserRole = "customer_service"
	RoleManager             UserRole = "manager"
	RoleDriver              UserRole = "driver"
	RoleTransportSupervisor UserRole = "transport_supervisor"
	RoleSalesStaff          UserRole = "sales_staff"
)

type UserStatus string

const (
	UserStatusActive    UserStatus = "active"
	UserStatusSuspended UserStatus = "suspended"
	UserStatusPending   UserStatus = "pending"
)

// User is the shared parent identity used by every module.
type User struct {
	UserID       string     `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	Name         string     `gorm:"column:name;type:text;not null" json:"name"`
	Phone        string     `gorm:"column:phone;type:text" json:"phone"`
	Email        string     `gorm:"column:email;type:text;not null;uniqueIndex" json:"email"`
	PasswordHash string     `gorm:"column:password_hash;type:text;not null" json:"-"`
	Role         UserRole   `gorm:"column:role;type:text;not null;index" json:"role"`
	Status       UserStatus `gorm:"column:status;type:text;not null;default:'active'" json:"status"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"column:updated_at;not null" json:"updated_at"`
}

func (User) TableName() string { return "users" }

type PurchasingStaff struct {
	UserID   string    `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	Position string    `gorm:"column:position;type:text;not null" json:"position"`
	HireDate time.Time `gorm:"column:hire_date;type:date;not null" json:"hire_date"`
	User     User      `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (PurchasingStaff) TableName() string { return "purchasing_staffs" }

type CustomerServiceOfficer struct {
	UserID    string `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	Position  string `gorm:"column:position;type:text;not null" json:"position"`
	ShiftTime string `gorm:"column:shift_time;type:text;not null" json:"shift_time"`
	User      User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (CustomerServiceOfficer) TableName() string { return "customer_service_officers" }

type Manager struct {
	UserID   string `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	Position string `gorm:"column:position;type:text;not null" json:"position"`
	User     User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (Manager) TableName() string { return "managers" }

type Driver struct {
	UserID string `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (Driver) TableName() string { return "drivers" }

type TransportSupervisor struct {
	UserID string `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (TransportSupervisor) TableName() string { return "transport_supervisors" }

type SalesStaff struct {
	UserID string `gorm:"column:user_id;type:text;primaryKey" json:"user_id"`
	User   User   `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user,omitempty"`
}

func (SalesStaff) TableName() string { return "sales_staff" }
