package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
	r.Use(cors.Default())
	nabavkeRepo, err := repository.PostgreSQLConnection(l)
	if err != nil {
		l.Println("Error connecting to postgres")
	}
	procurementService := service.NewProcurementService(l, nabavkeRepo)
	procurementController := controller.NewProcurementController(l, *procurementService)

	//	publicKey, err := client.ReadRSAPublicKeyFromFile("./public.pem")
	if err != nil {
		l.Println(err.Error())
		return
	}

	authorized := r.Group("/api")
	//authorized.Use(client.CheckAuthWithPublicKey(publicKey))
	//{
	//
	//}
	authorized.POST("/createProcurement", procurementController.CreateProcurement)
	authorized.POST("/createProcurementPlan", procurementController.CreateProcurementPlan)
	authorized.GET("/getProcurements", procurementController.GetProcurements)

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
