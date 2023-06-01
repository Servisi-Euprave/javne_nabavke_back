package repository

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
	"time"
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
	repo.l.Println("Offer_Repo: Insert Offers")

	startDateStr := time.Now().Format("2006-01-02")
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		repo.l.Println("Unable to parse date.", err)
	}
	createdOffer := repo.db.Create(&model.Offer{
		Id:             uuid.NewV4().String(),
		Price:          offerDTO.Price,
		BidderPib:      offerDTO.BidderPib,
		TermAndPayment: offerDTO.TermAndPayment,
		ProcurementId:  offerDTO.ProcurementId,
		Quantity:       offerDTO.Quantity,
		StartDate:      startDate,
	})
	errorMessage := createdOffer.Error
	if errorMessage != nil {
		repo.l.Printf("Unable to create offer: %s", errorMessage)
		return CannotCreateError("Offer")
	}

	return nil
}

func (repo *OfferRepository) GetOffers(procurementId string) ([]*model.Offer, error) {
	repo.l.Println("Offer_Repo: Get Offers")

	var offers []*model.Offer
	if err := repo.db.Table("offers").Where("procurement_id = ?", procurementId).Find(&offers).Error; err != nil {
		return nil, err
	}
	return offers, nil
}
func (repo *OfferRepository) GetResults(procurementId string) (*model.Offer, error) {
	repo.l.Println("Offer_repo: GetResults")

	var result *model.Offer
	if err := repo.db.First("offers").Where("procurement_id = ?", procurementId).Error; err != nil {
		return nil, err
	}

	return result, nil
}
