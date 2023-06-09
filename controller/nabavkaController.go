package controller

import (
	"github.com/gin-gonic/gin"
	"io"
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
	n.l.Println("Procurement_Controller - CreateProcurement")
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
	n.l.Println("Procurement_Controller - CreateProcurementPlan")

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
	n.l.Println("Procurement_Controller - Get procurement plan")

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
	n.l.Println("Procurement_Controller - GET procurements")

	procurements, err := n.service.GetProcurements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)
}
func (n *ProcurementController) GetCompanyFromToken(c *gin.Context) {
	n.l.Println("Procurement_Controller - GET company info")

	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	url := "http://apr-backend:7887/api/company/" + companyPiB.(string)

	// Send the HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.Data(http.StatusOK, "application/json; charset=utf-8", body)
}

func (n *ProcurementController) GetCompanyProcurements(c *gin.Context) {
	n.l.Println("Procurement_Controller - Get company procurements")

	companyId := c.Param("id")
	procurements, err := n.service.GetCompanyProcurements(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)

}

func (n *ProcurementController) DeclareWinner(c *gin.Context) {
	n.l.Println("Procurement_Controller - Declare winner")
	procId := c.Param("procId")

	offerId := c.Param("id")
	err := n.service.DeclareWinner(procId, offerId)
	if err != nil {
		n.l.Printf("Error occurred, Couldn't declare Winner")
		c.JSON(http.StatusInternalServerError, gin.H{"Error creating procurement! :": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Status": "Winner declared"})

}

func (n *ProcurementController) GetProcWithOffer(c *gin.Context) {
	n.l.Println("Procurement_Controller - Get procurements with offer")

	procWithOffers, err := n.service.GetProcurementAndWinningOffer()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procWithOffers)

}

func (n *ProcurementController) GetCompProcurements(c *gin.Context) {
	n.l.Println("Procurement_Controller - Get company Procurements")

	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	procurements, err := n.service.GetCompanyProcurements(companyPiB.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)

}

func (n *ProcurementController) GetAllProcurements(c *gin.Context) {
	n.l.Println("Procurement_Controller - Get ALL Procurements")
	procurements, err := n.service.GetAllFinishedProcurements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return

	}
	c.JSON(http.StatusOK, procurements)

}

func (n *ProcurementController) CheckIfCanPostOffer(c *gin.Context) {
	n.l.Println("Procurement_Controller - Check if can post offer")
	companyPiB, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"errors": "claims not found in context"})
		return
	}
	offerId := c.Param("id")
	if companyPiB.(string) == offerId {
		c.JSON(http.StatusBadRequest, gin.H{"message": "User can't create offer for themself"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "sve ok :)"})

}
