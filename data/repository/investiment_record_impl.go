package repository

import (
	"controle-investimento-back-end/data/database"
	"controle-investimento-back-end/domain/entities"
	"log"
	"time"
	 "database/sql"
)

// InvestmentRecordRepositoryImpl implements persistence for investment records.
type InvestmentRecordRepositoryImpl struct{}

func NewInvestmentRecordRepositoryImpl() *InvestmentRecordRepositoryImpl {
	return &InvestmentRecordRepositoryImpl{}
}

func (r *InvestmentRecordRepositoryImpl) SaveInvestmentRecord(record entities.InvestmentRecord) error {

	query := `
	INSERT INTO investments (
		date,
		emergency,
		fixed_income,
		variable_income,
		contribution,
		variation,
		total
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7
)
	`

	_, err := database.DB.Exec(
		query,
		record.Date,
		record.Emergency,
		record.FixedIncome,
		record.VariableIncome,
		record.Contribution,
		record.Variation,
		record.Total,
	)

	if err != nil {
		log.Printf("❌ Error SaveInvestmentRecord: %v\n", err)
		return err
	}

	return nil
}

func (r *InvestmentRecordRepositoryImpl) UpdateInvestmentRecord(record entities.InvestmentRecord) error {
	query := `
	UPDATE investments SET
		date = $1,
		emergency = $2,
		fixed_income = $3,
		variable_income = $4,
		contribution = $5,
		variation = $6,
		total = $7
	WHERE id = $8
	`

	_, err := database.DB.Exec(
		query,
		record.Date,
		record.Emergency,
		record.FixedIncome,
		record.VariableIncome,
		record.Contribution,
		record.Variation,
		record.Total,
		record.ID,
	)

	if err != nil {
		log.Printf("❌ Error UpdateInvestmentRecord: %v\n", err)
		return err
	}

	return nil
}

func (r *InvestmentRecordRepositoryImpl) GetAllInvestmentRecords(start, end time.Time, ascending bool) ([]entities.InvestmentRecord, error) {
	query := `
	SELECT id, date, emergency, fixed_income, variable_income, contribution, variation, total
	FROM investments

	`

	var args []interface{}

	if !start.IsZero() && !end.IsZero() {
		query += ` WHERE date >= $1 AND date < $2`
		args = append(args, start, end)
	}

	if ascending {
		query += `
	ORDER BY date ASC
	`
	} else {
		query += `
	ORDER BY date DESC
	`
	}

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Printf("❌ Error GetAllInvestmentRecords: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var records []entities.InvestmentRecord
	for rows.Next() {
		var record entities.InvestmentRecord
		err := rows.Scan(&record.ID, &record.Date, &record.Emergency, &record.FixedIncome, &record.VariableIncome, &record.Contribution, &record.Variation, &record.Total)
		if err != nil {
			log.Printf("❌ Error scanning row: %v\n", err)
			return nil, err
		}
		records = append(records, record)
	}

	return records, nil
}

func (r *InvestmentRecordRepositoryImpl) GetLastInvestmentRecord() (*entities.InvestmentRecord, error) {
	query := `
	SELECT date, emergency, fixed_income, variable_income, contribution, variation, total	
	FROM investments
	ORDER BY date DESC
	LIMIT 1
	`
	row := database.DB.QueryRow(query)

	var record entities.InvestmentRecord
	err := row.Scan(&record.Date, &record.Emergency, &record.FixedIncome, &record.VariableIncome, &record.Contribution, &record.Variation, &record.Total)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		log.Printf("❌ Error scanning row: %v\n", err)
		return nil, err
	}

	return &record, nil
}

func (r *InvestmentRecordRepositoryImpl) DataDashboard(start, end time.Time) (entities.InvestmentRecord, error) {

	db := database.DB

	query := `
	SELECT
	 COALESCE(SUM(emergency), 0) AS emergency,
	 COALESCE(SUM(fixed_income), 0) AS fixed_income,
	 COALESCE(SUM(variable_income), 0) AS variable_income,
	 COALESCE(SUM(total), 0) AS total
	FROM investments
	`
	var args []interface{}

	if !start.IsZero() && !end.IsZero() {
		query += ` WHERE date >= $1 AND date < $2`
		args = append(args, start, end)
	}

	row := db.QueryRow(query, args...)

	var record entities.InvestmentRecord
	err := row.Scan(&record.Emergency, &record.FixedIncome, &record.VariableIncome, &record.Total)
	if err != nil {
		log.Printf("❌ Error scanning row: %v\n", err)
		return entities.InvestmentRecord{}, err
	}
	return record, nil
}

func (r *InvestmentRecordRepositoryImpl) AssetGrowth(start, end time.Time) ([]entities.AssetGrowth, error) {
	query := `
	SELECT 
	   DISTINCT EXTRACT(YEAR FROM date) AS year , 
	   SUM(total) 
	FROM investments
	`
	var args []interface{}

	if !start.IsZero() && !end.IsZero() {
		query += ` WHERE date >= $1 AND date < $2`
		args = append(args, start, end)
	}

	query += ` GROUP BY year ORDER BY year ASC`

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Printf("❌ Error AssetGrowth: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var result []entities.AssetGrowth
	for rows.Next() {
		var growth entities.AssetGrowth
		err := rows.Scan(&growth.Year, &growth.Total)
		if err != nil {
			log.Printf("❌ Error scanning row: %v\n", err)
			return nil, err
		}
		result = append(result, growth)
	}

	return result, nil
}

func (r *InvestmentRecordRepositoryImpl) CategoryGrowth(start, end time.Time) ([]entities.CategoryGrowth, error) {
	query := `
	SELECT
       DATE_TRUNC('month', date) AS date,
       SUM(emergency) AS emergency,
       SUM(fixed_income) AS fixed_income,
       SUM(variable_income) AS variable_income
    FROM public.investments
	`
	var args []interface{}

	if !start.IsZero() && !end.IsZero() {
		query += ` WHERE date >= $1 AND date < $2`
		args = append(args, start, end)
	}

	query += ` GROUP BY DATE_TRUNC('month', date)
	ORDER BY date
	`

	rows, err := database.DB.Query(query, args...)
	if err != nil {
		log.Printf("❌ Error CategoryGrowth: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var result []entities.CategoryGrowth
	for rows.Next() {

		var growth entities.CategoryGrowth
		err := rows.Scan(&growth.Date, &growth.Emergency, &growth.FixedIncome, &growth.VariableIncome)
		if err != nil {
			log.Printf("❌ Error scanning row: %v\n", err)
			return nil, err
		}
		result = append(result, growth)
	}
	return result, nil
}

func (r *InvestmentRecordRepositoryImpl) AvailableYears() ([]int, error) {
	query := `
	SELECT DISTINCT EXTRACT(YEAR FROM date) AS year
	FROM investments
	`
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Printf("❌ Error AvailableYears: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var result []int
	for rows.Next() {
		var year int
		err := rows.Scan(&year)
		if err != nil {
			log.Printf("❌ Error scanning row: %v\n", err)
			return nil, err
		}
		result = append(result, year)
	}

	return result, nil
}
