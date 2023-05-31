package service

import (
	"javne_nabavke_back/model"
	"javne_nabavke_back/repository"
	"log"
)

type OfferService struct {
	l               *log.Logger
	offerRepository repository.IOfferRepo
}

func NewOfferOfferService(l *log.Logger, repo repository.IOfferRepo) *OfferService {
	return &OfferService{l, repo}
}

func (s *OfferService) InsertOffer(offerDTO *model.OfferRequestDTO) error {
	s.l.Printf("InsertOffer")
	return s.offerRepository.InsertOffer(offerDTO)
}

func (s *OfferService) GetOffers(procurementId string) ([]*model.Offer, error) {
	return s.offerRepository.GetOffers(procurementId)
}
