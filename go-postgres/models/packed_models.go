package models

type MaterialForPeriod struct {
	Material  Material `json:"material"`
	StartDate string   `json:"start_date"`
	EndDate   string   `json:"end_date"`
}

type ProfessionForPeriod struct {
	Profession Profession `json:"profession"`
	Product    Product    `json:"product"`
	StartDate  string     `json:"start_date"`
	EndDate    string     `json:"end_date"`
}

type LaborCostForProduct struct {
	ProfessionName string  `json:"profession"`
	ProductName    string  `json:"product"`
	ActualDate     string  `json:"actual_date"`
	Cost           float64 `json:"cost,string"`
}

type PackedMaterialCost struct {
	MaterialName string  `json:"material"`
	ActualDate   string  `json:"actual_date"`
	Cost         float64 `json:"cost,string"`
}

type PackedProduct struct {
	ID          int64        `json:"id,string"`
	Name        string       `json:"name"`
	Materials   []Material   `json:"materials"`
	Professions []Profession `json:"professions"`
}

type PackedProductWithActualCosts struct {
	ID          int64                      `json:"id,string"`
	Name        string                     `json:"name"`
	Materials   []MaterialWithActualCost   `json:"materials"`
	Professions []ProfessionWithActualCost `json:"professions"`
}

type PackedCalculation struct {
	ID                  int64                        `json:"id,string"`
	StartDate           string                       `json:"start_date"`
	EndDate             string                       `json:"end_date"`
	Product             PackedProductWithActualCosts `json:"product"`
	CostItems           []PackedCostItem             `json:"cost_items"`
	CalculatedCostItems []CalculatedCostItem         `json:"calculated_cost_items"`
	FullCost            float64                      `json:"full_cost,string"`
}

type MaterialWithActualCost struct {
	ID   int64   `json:"id,string"`
	Name string  `json:"name"`
	Cost float64 `json:"cost,string"`
}

type ProfessionWithActualCost struct {
	ID          int64   `json:"id,string"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost,string"`
}

type PackedCostItem struct {
	ID   int64   `json:"id,string"`
	Name string  `json:"name"`
	Cost float64 `json:"cost,string"`
}

type CalculatedCostItem struct {
	Name string  `json:"name"`
	Cost float64 `json:"cost,string"`
}
