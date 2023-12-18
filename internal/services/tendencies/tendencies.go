package tendencies

import (
	"fmt"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"gorm.io/gorm"
)

type TendencyService struct {
	db *gorm.DB
}

func NewTendencyService(db *gorm.DB) *TendencyService {
	return &TendencyService{
		db,
	}
}

func (ts *TendencyService) GetTendency(dto *dtos.GetTendencyDTO) (*charts.Bar, error) {
	patient, err := ts.getPatientByName(dto.PatientName)
	if err != nil {
		return nil, err
	}

	indicators, err := patient.IndicatorInteraction.GetDynamicIndicators(dto.IndicatorName)
	if err != nil {
		return nil, err
	}

	bar := charts.NewBar()

	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: fmt.Sprintf("Tendency of %d for %s", dto.IndicatorName, dto.PatientName),
	}))
	bar.AddSeries("Category B", ts.convertToBarItems(indicators))

	return bar, nil
}

func (ts *TendencyService) getPatientByName(name string) (*models.Patient, error) {
	patient := models.Patient{}
	if err := ts.db.Preload("IndicatorInteraction").Preload("MedicalWorker").First(&patient, "name = ?", name).Error; err != nil {
		return nil, err
	}

	return &patient, nil
}

func (ts *TendencyService) convertToBarItems(values []float64) []opts.BarData {
	items := make([]opts.BarData, len(values))
	for _, v := range values {
		items = append(items, opts.BarData{
			Value: v,
		})
	}

	return items
}
