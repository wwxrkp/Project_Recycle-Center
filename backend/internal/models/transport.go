package entity

import "time"

type Truck struct {
	TruckID      string  `gorm:"column:truck_id;type:text;primaryKey" json:"truck_id"`
	LicensePlate string  `gorm:"column:license_plate;type:text;not null;uniqueIndex" json:"license_plate"`
	Status       string  `gorm:"column:status;type:text;not null;index" json:"status"`
	Capacity     float64 `gorm:"column:capacity;type:numeric(14,3);not null" json:"capacity"`
	DriverID     *string `gorm:"column:driver_id;type:text;uniqueIndex" json:"driver_id,omitempty"`
	Driver       *Driver `gorm:"foreignKey:DriverID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"driver,omitempty"`
}

func (Truck) TableName() string { return "trucks" }

type DeliveryRequest struct {
	RequestID            string                    `gorm:"column:request_id;type:text;primaryKey" json:"request_id"`
	CustomerName         string                    `gorm:"column:customer_name;type:text;not null" json:"customer_name"`
	Address              string                    `gorm:"column:address;type:text;not null" json:"address"`
	Phone                string                    `gorm:"column:phone;type:text" json:"phone"`
	RequestDate          time.Time                 `gorm:"column:request_date;not null" json:"request_date"`
	Status               string                    `gorm:"column:status;type:text;not null;index" json:"status"`
	SupervisorID         string                    `gorm:"column:supervisor_id;type:text;not null;index" json:"supervisor_id"`
	TruckID              *string                   `gorm:"column:truck_id;type:text;index" json:"truck_id,omitempty"`
	DestinationLatitude  float64                   `gorm:"column:destination_latitude;type:double precision" json:"destination_latitude"`
	DestinationLongitude float64                   `gorm:"column:destination_longitude;type:double precision" json:"destination_longitude"`
	Supervisor           *TransportSupervisor      `gorm:"foreignKey:SupervisorID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"supervisor,omitempty"`
	Truck                *Truck                    `gorm:"foreignKey:TruckID;references:TruckID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"truck,omitempty"`
	Materials            []DeliveryRequestMaterial `gorm:"foreignKey:RequestID;references:RequestID" json:"materials,omitempty"`
}

func (DeliveryRequest) TableName() string { return "delivery_requests" }

type DeliveryRequestMaterial struct {
	RequestID        string           `gorm:"column:request_id;type:text;primaryKey" json:"request_id"`
	MaterialID       string           `gorm:"column:material_id;type:text;primaryKey" json:"material_id"`
	DeliveryQuantity float64          `gorm:"column:delivery_quantity;type:numeric(14,3);not null" json:"delivery_quantity"`
	DeliveryRequest  *DeliveryRequest `gorm:"foreignKey:RequestID;references:RequestID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"delivery_request,omitempty"`
	Material         *Material        `gorm:"foreignKey:MaterialID;references:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"material,omitempty"`
}

func (DeliveryRequestMaterial) TableName() string { return "delivery_request_materials" }

type CancelRequest struct {
	CancelID        string           `gorm:"column:cancel_id;type:text;primaryKey" json:"cancel_id"`
	TruckID         string           `gorm:"column:truck_id;type:text;not null;index" json:"truck_id"`
	CancelDate      time.Time        `gorm:"column:cancel_date;not null" json:"cancel_date"`
	RequestID       string           `gorm:"column:request_id;type:text;not null;uniqueIndex" json:"request_id"`
	Reason          string           `gorm:"column:reason;type:text;not null" json:"reason"`
	Truck           *Truck           `gorm:"foreignKey:TruckID;references:TruckID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"truck,omitempty"`
	DeliveryRequest *DeliveryRequest `gorm:"foreignKey:RequestID;references:RequestID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT" json:"delivery_request,omitempty"`
}

func (CancelRequest) TableName() string { return "cancel_requests" }
