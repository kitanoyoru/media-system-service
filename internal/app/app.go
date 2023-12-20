package app

import (
	v0 "github.com/kitanoyoru/media-system-service/internal/api/http/v0"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"github.com/kitanoyoru/media-system-service/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Run() error {
	dsn, err := utils.ConnectionURLBuilder("postgres")
	if err != nil {
		return err
	}

	driver := postgres.New(postgres.Config{
		DSN: dsn,
	})

	db, err := gorm.Open(driver, &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		return err
	}

	db.AutoMigrate(models.Height{}, models.Weight{}, models.HeartRate{}, models.BloodPressure{})
	db.AutoMigrate(models.IndicatorInteraction{})
	db.AutoMigrate(models.MedicalWorker{}, models.Patient{})

	r, err := v0.NewRouter(db)
	if err != nil {
		return err
	}

	utils.StartServerWithGracefulShutdown(r)

	return nil
}
