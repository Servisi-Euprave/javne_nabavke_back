package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"javne_nabavke_back/client"
	"javne_nabavke_back/controller"
	"javne_nabavke_back/repository"
	"javne_nabavke_back/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	port := os.Getenv("PROCUREMENTS_PORT")
	if len(port) == 0 {
		port = "8082"
	}
	l := log.New(os.Stdout, "Javne_nabavke", log.LstdFlags)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:4200", "http://localhost:4201", "http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//nabavkeRepo, err := repository.PostgreSQLConnection(l)

	dbConnection, err := repository.PostgreSQLConnection(l)
	if err != nil {
		l.Println("Error connecting to postgres")
	}

	if err != nil {
		l.Println("Error creating PostgreSQL connection")
	}

	procurementRepo := repository.CreateProcurementRepository(l, dbConnection)
	offerRepo := repository.CreateOfferRepository(l, dbConnection)

	procurementService := service.NewProcurementService(l, procurementRepo)
	procurementController := controller.NewProcurementController(l, *procurementService)

	offerService := service.NewOfferOfferService(l, offerRepo)
	offerController := controller.NewOfferController(l, *offerService)

	publicKey, err := client.ReadRSAPublicKeyFromFile("./public.pem")
	if err != nil {
		l.Println(err.Error())
		return
	}

	authorized := r.Group("/authorizedApi")
	authorized.Use(client.CheckAuthWithPublicKey(publicKey))
	{
		authorized.POST("/createProcurement", procurementController.CreateProcurement)
		authorized.POST("/createProcurementPlan", procurementController.CreateProcurementPlan)
		authorized.GET("/getProcurementPlans", procurementController.GetProcurementPlans)
		authorized.POST("/postOffer", offerController.CreateOffer)
		authorized.GET("/getProcurementOffers/:id", offerController.ProcurementOffers)
		authorized.PUT("/declareWinner/:procId/:id", procurementController.DeclareWinner)
		authorized.GET("/getCompProcurements", procurementController.GetCompProcurements)
		authorized.GET("/checkIfCanPostOffer/:id", procurementController.CheckIfCanPostOffer)
		authorized.GET("/companyProfile", procurementController.GetCompanyFromToken)

	}

	open := r.Group("/api")
	open.GET("/getProcurements", procurementController.GetProcurements)
	open.GET("/getAllProcurements", procurementController.GetAllProcurements)
	open.GET("/getCompanyProcurements/:id", procurementController.GetCompanyProcurements)
	open.GET("/getProcurementAndOfferList", procurementController.GetProcWithOffer)

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	l.Println("Server listening on port" + port)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	l.Println("Graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
