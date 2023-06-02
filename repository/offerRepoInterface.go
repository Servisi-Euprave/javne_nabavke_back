package repository

import "javne_nabavke_back/model"

type IOfferRepo interface {
	InsertOffer(offerDTO *model.OfferRequestDTO) error
	GetOffers(procurementId string) ([]*model.Offer, error)
	GetResults(procurementId string) (*model.Offer, error)
}
