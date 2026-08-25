package entities

type InvestmentRecord struct {
	ID             int     `json:"id"`
	Date           string  `json:"date"`
	Emergency      float64 `json:"emergency"`
	FixedIncome    float64 `json:"fixed_income"`
	VariableIncome float64 `json:"variable_income"`
	Contribution   float64 `json:"contribution"`
	Variation      float64 `json:"variation"`
	Total          float64 `json:"total"`
}

type AssetGrowth struct {
	Year  int     `json:"year"`
	Total float64 `json:"total"`
}

type CategoryGrowth struct {
	Date           string  `json:"date"`
	Emergency      float64 `json:"emergency"`
	FixedIncome    float64 `json:"fixed_income"`
	VariableIncome float64 `json:"variable_income"`
}
