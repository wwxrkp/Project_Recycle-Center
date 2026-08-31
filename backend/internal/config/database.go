package config

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	entity "github.com/SA-1-69/T20/backend/internal/models"
)

var DB *gorm.DB

func ConnectDatabase() error {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return errors.New("DATABASE_URL is required")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		return fmt.Errorf("connect to PostgreSQL: %w", err)
	}

	if err := db.AutoMigrate(
		// Shared parent identity and role profiles.
		&entity.User{},
		&entity.PurchasingStaff{},
		&entity.CustomerServiceOfficer{},
		&entity.Manager{},
		&entity.Driver{},
		&entity.TransportSupervisor{},
		&entity.SalesStaff{},

		// Ping: seller registration and external parties.
		&entity.Seller{},
		&entity.RegistrationForm{},
		&entity.Factory{},

		// Filme: material master, quality assessment, and inventory.
		&entity.MaterialType{},
		&entity.Material{},
		&entity.MaterialAssessmentItem{},
		&entity.QualityAssessment{},
		&entity.ScrapPurchaseItem{},
		&entity.PendingWarehouseItem{},
		&entity.Warehouse{},
		&entity.StorageZone{},
		&entity.StockTransaction{},
		&entity.ReceiveTransaction{},
		&entity.IssueTransaction{},
		&entity.StockAdjustmentRequest{},
		&entity.AdjustmentApproval{},

		// Mark: transport and procurement.
		&entity.Truck{},
		&entity.DeliveryRequest{},
		&entity.DeliveryRequestMaterial{},
		&entity.CancelRequest{},
		&entity.Supplier{},
		&entity.PurchaseOrder{},
		&entity.PurchaseOrderMaterial{},
		&entity.PurchaseOrderSupplier{},

		// Ping: factory sales, billing, and complaints.
		&entity.SalesContract{},
		&entity.ScrapSaleToFactory{},
		&entity.SalesBill{},
		&entity.Complaint{},
	); err != nil {
		return fmt.Errorf("auto migrate database: %w", err)
	}

	DB = db
	return nil
}
