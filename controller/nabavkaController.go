package controller

import (
	"github.com/gin-gonic/gin"
	"javne_nabavke_back/service"
	"log"
)

type NabavkaController struct {
	l       *log.Logger
	service service.NabavkaService
}

func NewNabavkaController(l *log.Logger, service service.NabavkaService) *NabavkaController {
	return &NabavkaController{
		l:       l,
		service: service,
	}
}
func (n *NabavkaController) CreateNabavka(c *gin.Context) {

}

func (n *NabavkaController) CreatePonuda(context *gin.Context) {

}
