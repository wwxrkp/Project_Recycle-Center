package entity

import (
	"sync"
	"testing"

	"gorm.io/gorm/schema"
)

type tableNamer interface {
	TableName() string
}

func TestModelTableNamesAreUnique(t *testing.T) {
	models := allModels()

	seen := make(map[string]struct{}, len(models))
	for _, model := range models {
		tableName := model.TableName()
		if _, duplicate := seen[tableName]; duplicate {
			t.Errorf("duplicate table name: %s", tableName)
		}
		seen[tableName] = struct{}{}
	}

	if got, want := len(seen), len(models); got != want {
		t.Fatalf("unique table count = %d, want %d", got, want)
	}
}

func TestModelsHaveValidGormSchema(t *testing.T) {
	var cache sync.Map
	namingStrategy := schema.NamingStrategy{}

	for _, model := range allModels() {
		if _, err := schema.Parse(model, &cache, namingStrategy); err != nil {
			t.Errorf("parse GORM schema for %T: %v", model, err)
		}
	}
}

func allModels() []tableNamer {
	return []tableNamer{
		User{}, PurchasingStaff{}, CustomerServiceOfficer{}, Manager{},
		Driver{}, TransportSupervisor{}, SalesStaff{}, Seller{},
		RegistrationForm{}, Factory{}, MaterialType{}, Material{},
		MaterialAssessmentItem{}, QualityAssessment{}, ScrapPurchaseItem{},
		PendingWarehouseItem{}, Warehouse{}, StorageZone{}, StockTransaction{},
		ReceiveTransaction{}, IssueTransaction{}, StockAdjustmentRequest{},
		AdjustmentApproval{}, Truck{}, DeliveryRequest{}, DeliveryRequestMaterial{},
		CancelRequest{}, Supplier{}, PurchaseOrder{}, PurchaseOrderMaterial{},
		PurchaseOrderSupplier{}, SalesContract{}, ScrapSaleToFactory{}, SalesBill{},
		Complaint{},
	}
}
