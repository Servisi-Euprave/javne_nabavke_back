package controller

import (
	"github.com/gin-gonic/gin"
	"javne_nabavke_back/model"
	"javne_nabavke_back/service"
	"log"
	"net/http"
)

type ProcurementController struct {
	l       *log.Logger
	service service.ProcurementService
}

func NewProcurementController(l *log.Logger, service service.ProcurementService) *ProcurementController {
	return &ProcurementController{
		l:       l,
		service: service,
	}
}
func (n *ProcurementController) CreateProcurement(c *gin.Context) {
	companyPiB, exists := c.Get("claims")
	if !exists {
		// Handle errors case where claims are not set in the context
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	var procurement model.Procurement
	if err := c.ShouldBindJSON(&procurement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	procurement.ProcuringEntityPiB = companyPiB.(string)
	err := n.service.SaveProcurement(&procurement)
	if err != nil {
		n.l.Printf("Error occurred, Couldn't create procurement")
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating procurement! :": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Procurement created"})

}
func (n *ProcurementController) CreateProcurementPlan(c *gin.Context) {
	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}

	var procurementPlan model.ProcurementPlan
	if err := c.ShouldBindJSON(&procurementPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	procurementPlan.CompanyPib = companyPiB.(string)
	err := n.service.SaveProcurementPlan(&procurementPlan)
	if err != nil {
		n.l.Printf("Error occurred, Couldn't create procurement plan")
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating procurement plan! :": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Procurement plan created"})

}
func (n *ProcurementController) GetProcurementPlans(c *gin.Context) {
	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	plans, err := n.service.GetProcurementPlans(companyPiB.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, plans)

}

func (n *ProcurementController) GetProcurements(c *gin.Context) {

	procurements, err := n.service.GetProcurements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)
}

func (n *ProcurementController) GetCompanyProcurements(c *gin.Context) {
	companyId := c.Param("id")
	procurements, err := n.service.GetCompanyProcurements(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)

}

func (n *ProcurementController) DeclareWinner(c *gin.Context) {
	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	offerId := c.Param("id")
	err := n.service.DeclareWinner(companyPiB.(string), offerId)
	if err != nil {
		n.l.Printf("Error occurred, Couldn't declare Winner")
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating procurement! :": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Winner declared"})

}

func (n *ProcurementController) GetProcWithOffer(c *gin.Context) {
	procWithOffers, err := n.service.GetProcurementAndWinningOffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procWithOffers)

}
