package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
)

type OfferRepository struct {
	l  *log.Logger
	db *gorm.DB
}

func CreateOfferRepository(log *log.Logger, dbConn *gorm.DB) *OfferRepository {
	repository := &OfferRepository{log, dbConn}
	return repository
}

func (repo *OfferRepository) InsertOffer(offerDTO *model.OfferRequestDTO) error {
	createdOffer := repo.db.Create(&model.Offer{
		Id:             uuid.NewV4().String(),
		Price:          offerDTO.Price,
		BidderPib:      offerDTO.BidderPib,
		TermAndPayment: offerDTO.TermAndPayment,
		ProcurementId:  offerDTO.ProcurementId,
	})
	errorMessage := createdOffer.Error
	if errorMessage != nil {
		repo.l.Printf("Unable to create offer: %s", errorMessage)
		return CannotCreateError("Offer")
	}

	return nil
}

func (repo *OfferRepository) GetOffers(procurementId string) ([]*model.Offer, error) {
	var offers []*model.Offer
	if err := repo.db.Table("offers").Where("procurement_id = ?", procurementId).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}
func (repo *OfferRepository) GetResults(procurementId string) (*model.Offer, error) {
	var result *model.Offer
	if err := repo.db.First("offers").Where("procurement_id = ?", procurementId).Error; err != nil {
		return nil, err
	}

	return result, nil
}
