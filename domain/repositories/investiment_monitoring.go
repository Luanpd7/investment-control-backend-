package repositories

import (
	"controle-investimento-back-end/domain/entities"
	"time"
)

type InvestmentRecordRepository interface {
	SaveInvestmentRecord(record entities.InvestmentRecord) error

	UpdateInvestmentRecord(record entities.InvestmentRecord) error

	GetAllInvestmentRecords(start time.Time, end time.Time, ascending bool) ([]entities.InvestmentRecord, error)

	GetLastInvestmentRecord(date string) (*entities.InvestmentRecord, error)

	DataDashboard(start time.Time, end time.Time) (entities.InvestmentRecord, error)

	AssetGrowth(start, end time.Time) ([]entities.AssetGrowth, error)

	CategoryGrowth(start, end time.Time) ([]entities.CategoryGrowth, error)

	AvailableYears() ([]int, error)
}
