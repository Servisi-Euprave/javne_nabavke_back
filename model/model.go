package model

import "time"

type Offer struct {
	Price          float64 `json:"price"`
	BidderPib      string
	TermAndPayment string
	ProcurementId  string `json:"procurement_id"`
}
type ProcurementPlan struct {
	ProcurementPlanId string `gorm:"default:uuid_generate_v3()"`
	ProductType       string `json:"product_type"`
	EstimatedValue    string `json:"estimated_value"`
	Quantity          int    `json:"quantity"`
}

type Procurement struct {
	Id                 string    `gorm:"default:uuid_generate_v3()"`
	ProcuringEntityPiB string    `json:"procuring_entity_pi_b"`
	ProcurementPlanId  int       `json:"procurement_plan_id"`
	StartDate          time.Time `json:"start_date"`
	EndDate            string    `json:"end_date"`
	ProcurementName    string    `json:"procurement_name"`
	Description        string    `json:"description"`
}
