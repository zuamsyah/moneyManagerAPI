package models

type Home struct {
	ID           int `json:"id" gorm:"primaryKey"`
	CurrentSaldo int `json:"current_saldo"`
	TotalExpense int `json:"total_expense"`
	TotalIncome  int `json:"total_income"`
}
