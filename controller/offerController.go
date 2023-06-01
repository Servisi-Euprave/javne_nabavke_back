package controller

import (
	"github.com/gin-gonic/gin"
	"javne_nabavke_back/model"
	"javne_nabavke_back/service"
	"log"
	"net/http"
)

type OfferController struct {
	l           *log.Logger
	service     service.OfferService
	procService service.ProcurementService
}

func NewOfferController(l *log.Logger, service service.OfferService) *OfferController {
	return &OfferController{
		l:       l,
		service: service,
	}
}
func (n *OfferController) CreateOffer(c *gin.Context) {
	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}

	var offer model.OfferRequestDTO
	offer.BidderPib = companyPiB.(string)
	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err := n.service.InsertOffer(&offer)
	if err != nil {
		n.l.Printf("Error occurred, Couldn't create offer")
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating offer! :": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Offer created"})
}
func (n *OfferController) ProcurementOffers(c *gin.Context) {
	procId := c.Param("id")
	offers, err := n.service.GetOffers(procId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, offers)

}
