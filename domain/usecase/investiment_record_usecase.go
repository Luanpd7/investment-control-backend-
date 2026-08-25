package usecase

import (
	"controle-investimento-back-end/domain/entities"
	domainRepo "controle-investimento-back-end/domain/repositories"
	"errors"
	"time"
)

// InvestmentRecordUseCase contains business rules for investment record management.
type InvestmentRecordUseCase struct {
	repository domainRepo.InvestmentRecordRepository
}

func NewInvestmentRecordUseCase(repo domainRepo.InvestmentRecordRepository) *InvestmentRecordUseCase {
	return &InvestmentRecordUseCase{repository: repo}
}

// SaveInvestmentRecord executes the simulation (business rules) and persists the record.
func (uc *InvestmentRecordUseCase) SaveInvestmentRecord(input *entities.InvestmentRecord) error {

	total := uc.totalInvestiment(input)

	variation := uc.variationInvestment(input)

	// Build record to persist
	record := entities.InvestmentRecord{
		ID:             input.ID,
		Date:           input.Date,
		Emergency:      input.Emergency,
		FixedIncome:    input.FixedIncome,
		VariableIncome: input.VariableIncome,
		Contribution:   input.Contribution,
		Variation:      variation,
		Total:          total,
	}

	if record.ID == 0 {
		// New record, save it
		if err := uc.repository.SaveInvestmentRecord(record); err != nil {
			return err
		}
	} else {
		// Existing record, update it
		if err := uc.repository.UpdateInvestmentRecord(record); err != nil {
			return err
		}
	}

	return nil
}

func (uc *InvestmentRecordUseCase) totalInvestiment(record *entities.InvestmentRecord) float64 {
	return record.Emergency + record.FixedIncome + record.VariableIncome
}

func (uc *InvestmentRecordUseCase) variationInvestment(record *entities.InvestmentRecord) float64 {

	/// Get the last investment record to calculate the variation

	totalLastInvestment, err := uc.repository.GetLastInvestmentRecord()
	if err != nil {
		return 0.0
	}

	  if totalLastInvestment == nil {
        return 0.0
    }

	variation := ((record.Emergency + record.FixedIncome + record.VariableIncome) - totalLastInvestment.Total)

	return variation
}

func (uc *InvestmentRecordUseCase) GetAllInvestmentRecords(filter string, ascending bool) ([]entities.InvestmentRecord, error) {
	start, end, err := uc.parseDateFilter(filter)

	if err != nil {
		return nil, err
	}

	records, err := uc.repository.GetAllInvestmentRecords(start, end, ascending)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (uc *InvestmentRecordUseCase) DataDashboard(filter string) (entities.InvestmentRecord, error) {
	start, end, err := uc.parseDateFilter(filter)
	if err != nil {
		return entities.InvestmentRecord{}, err
	}

	data, err := uc.repository.DataDashboard(start, end)
	if err != nil {
		return entities.InvestmentRecord{}, err
	}
	return data, nil
}

func (uc *InvestmentRecordUseCase) parseDateFilter(filter string) (time.Time, time.Time, error) {
	if filter == "" {
		return time.Time{}, time.Time{}, nil
	}

	switch len(filter) {
	case 4:
		date, err := time.Parse("2006", filter)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}

		start := time.Date(
			date.Year(),
			1,
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		)

		end := start.AddDate(1, 0, 0)

		return start, end, nil

	case 7:
		date, err := time.Parse("2006-01", filter)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}

		start := time.Date(
			date.Year(),
			date.Month(),
			1,
			0,
			0,
			0,
			0,
			time.UTC,
		)

		end := start.AddDate(0, 1, 0)

		return start, end, nil

	default:
		return time.Time{}, time.Time{}, errors.New("invalid date filter format")
	}
}

func (uc *InvestmentRecordUseCase) GetLastInvestmentRecord() (*entities.InvestmentRecord, error) {

	var lastRecords *entities.InvestmentRecord
	var err error

	if lastRecords, err = uc.repository.GetLastInvestmentRecord(); err != nil {
		return nil, err
	}

	return lastRecords, nil
}

func (uc *InvestmentRecordUseCase) AssetGrowth(filter string) ([]entities.AssetGrowth, error) {

	start, end, err := uc.parseDateFilter(filter)
	if err != nil {
		return nil, err
	}
	result, err := uc.repository.AssetGrowth(start, end)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *InvestmentRecordUseCase) CategoryGrowth(filter string) ([]entities.CategoryGrowth, error) {

	start, end, err := uc.parseDateFilter(filter)
	if err != nil {
		return nil, err
	}

	result, err := uc.repository.CategoryGrowth(start, end)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *InvestmentRecordUseCase) AvailableYears() ([]int, error) {

	result, err := uc.repository.AvailableYears()
	if err != nil {
		return nil, err
	}
	return result, nil

}
