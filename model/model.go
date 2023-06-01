package model

import (
	"time"
)

type Offer struct {
	Id             string    `json:"id"`
	Price          string    `json:"price"`
	BidderPib      string    `json:"bidder_pib"`
	TermAndPayment string    `json:"term_and_payment"`
	ProcurementId  string    `json:"procurement_id"`
	StartDate      time.Time `json:"start_date"`
}

type OfferRequestDTO struct {
	Price          string `json:"price"`
	BidderPib      string `json:"bidder_pib"`
	TermAndPayment string `json:"term_and_payment"`
	Quantity       int    `json:"quantity"`
	ProcurementId  string `json:"procurement_id"`
}

type ProcurementPlan struct {
	CompanyPib        string `json:"company_pib"`
	ProcurementPlanId string `json:"procurement_plan_id"`
	ProductType       string `json:"product_type"`
	EstimatedValue    string `json:"estimated_value"`
	Quantity          int    `json:"quantity"`
}

type Procurement struct {
	Id                 string    `json:"id"`
	ProcuringEntityPiB string    `json:"procuring_entity_pi_b"`
	ProcurementPlanId  string    `json:"procurement_plan_id"`
	StartDate          time.Time `json:"start_date"`
	EndDate            string    `json:"end_date"`
	ProcurementName    string    `json:"procurement_name"`
	Description        string    `json:"description"`
	winnerId           string    `json:"winner_id"`
}
type ProcurementWithWinnerOffer struct {
	ProcuringEntityPiB string    `json:"procuring_entity_pi_b"`
	StartDate          time.Time `json:"start_date"`
	EndDate            string    `json:"end_date"`
	ProcurementName    string    `json:"procurement_name"`
	Description        string    `json:"description"`
	Price              string    `json:"price"`
	BidderPib          string    `json:"bidder_pib"`
	TermAndPayment     string    `json:"term_and_payment"`
}
