package models

import "time"

type CardDTO struct {
	Id int64 `json:"id"`
	Issuer string `json:"issuer"`
	Type string `json:"type"`
	Number string `json:"number"`
}
type Card struct {
	Id int64 `json:"id"`
	Issuer string `json:"issuer"`
	Type string `json:"type"`
	OwnerId int64 `json:"owner_id"`
	Number string `json:"number"`
	Balance int64 `json:"balance"`
}
type SimpleDTO struct{
	Name string
}
type TransactionsDTO struct{
	Id int64 `json:"id"`
	Mcc string `json:"mcc"`
	IconId int64 `json:"icon_id"`
	Amount int64 `json:"amount"`
	Date time.Time `json:"date"`
	CardId int64 `json:"card_id"`
}
type Result struct{
	Result string `json:"result"`
	ErrorDescription string `json:"errorDesc,omitempty"`

}
type StatsDTO struct {
	MccVisited string `json:"mcc_visited"`
	MccSpent string `json:"mcc_spent"`
	ValueVisited int64 `json:"value_visited"`
	ValueSpent int64 `json:"value_spent"`
}
type MostSpentDTO struct{
	Mcc string `json:"mcc"`
	Value int64 `json:"value"`
}