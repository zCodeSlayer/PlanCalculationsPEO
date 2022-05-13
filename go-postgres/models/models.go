package models

// Roles

type Group struct {
	ID          int64  `json:"id,string"`
	Name        string `json:"name" validate:"required,min=4"`
	Permissions string `json:"permissions" validate:"required,min=4"`
}

type User struct {
	ID       int64  `json:"id,string"`
	Login    string `json:"login" validate:"required,min=6"`
	Password string `json:"password"`
	Role     int64  `json:"id_role,string"`
}

// Calculations (subject area)

type Calculation struct {
	ID        int64  `json:"id,string"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	ProductID int64  `json:"id_product,string"`
}

type CostItem struct {
	ID     int64  `json:"id,string"`
	Name   string `json:"name"`
	Cipher string `json:"cipher"`
}

type PeriodForItem struct {
	Cost          float64 `json:"cost,string"`
	CostItemID    int64   `json:"id_cost_item,string"`
	CalculationID int64   `json:"id_calc,string"`
}

type Product struct {
	ID     int64  `json:"id,string"`
	Name   string `json:"name"`
	Cipher string `json:"cipher"`
}

// "right" branch

type Material struct {
	ID     int64  `json:"id,string"`
	Cipher string `json:"cipher"`
	Name   string `json:"name"`
}

type MaterialsNeed struct {
	Consumption float64 `json:"consumption,string"`
	NormID      int64   `json:"id_norm,string"`
	MaterialID  int64   `json:"id_material,string"`
	ProductID   int64   `json:"id_product,string"`
}

type MaterialCost struct {
	ID         int64   `json:"id,string"`
	ActualDate string  `json:"actual_date"`
	Cost       float64 `json:"cost,string"`
	NormID     int64   `json:"id_norm,string"`
	MaterialID int64   `json:"id_material,string"`
}

// "left" branch
type Profession struct {
	ID          int64  `json:"id,string"`
	Description string `json:"description"`
	Cipher      string `json:"cipher"`
}

type ProfessionsNeed struct {
	ProductID    int64 `json:"id_product,string"`
	ProfessionID int64 `json:"id_profession,string"`
}

type LaborCost struct {
	ID           int64   `json:"id,string"`
	ActualDate   string  `json:"actual_date"`
	Cost         float64 `json:"cost,string"`
	NormID       int64   `json:"id_norm,string"`
	ProfessionID int64   `json:"id_profession,string"`
}

// norms
type Norm struct {
	ID     int64  `json:"id,string"`
	Cipher string `json:"cipher"`
	Name   string `json:"name"`
}
