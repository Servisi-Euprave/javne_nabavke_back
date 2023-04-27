package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
	"os"
	"time"
)

type ProcurementPostgreSQL struct {
	l  *log.Logger
	db *gorm.DB
}

const (
	errorMsgProc = "Can't create procurement right now."
	errorMsgPlan = "Can't create procurement plan right now."
)

func PostgreSQLConnection(l *log.Logger) (*ProcurementPostgreSQL, error) {
	l.Println("PostrgeSQL_Repo")
	USERNAME := os.Getenv("USER")
	dbHost := os.Getenv("HOST")
	PASSWORD := os.Getenv("PASSWORD")
	dbName := os.Getenv("DB")
	PORT := os.Getenv("PORT")
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, USERNAME, dbName, PASSWORD, PORT)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		log.Println("Error establishing a database connection")
		panic("Failed to connect to database")
	}
	setup(db)
	l.Println("Successfully connected to postgres database")
	return &ProcurementPostgreSQL{l, db}, nil

}
func setup(db *gorm.DB) {
	db.AutoMigrate(&model.Procurement{})
	db.AutoMigrate(&model.ProcurementPlan{})
}
func QueryError(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (n *ProcurementPostgreSQL) AddProcurement(procurement *model.Procurement) error {
	n.l.Println("Procurement_Repository_Postgres")
	createdProcurement := n.db.Create(procurement)
	var errMessage = createdProcurement.Error
	if createdProcurement.Error != nil {
		n.l.Println("Unable to Create procurement.", errMessage)
		return QueryError(errorMsgProc)
	}
	return nil
}
func (n *ProcurementPostgreSQL) AddProcurementPlan(procurementPlan *model.ProcurementPlan) error {
	n.l.Println("Procurement_Plan_Repository_Postgres")
	createdProcurementPlan := n.db.Create(procurementPlan)
	var errMessage = createdProcurementPlan.Error
	if createdProcurementPlan.Error != nil {
		n.l.Println("Unable to Create procurement plan.", errMessage)
		return QueryError(errorMsgPlan)
	}
	return nil
}
func (n *ProcurementPostgreSQL) GetProcurements() ([]model.Procurement, error) {
	var procurements []model.Procurement
	now := time.Now()
	lastMonthStart := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.Local)
	lastMonthEnd := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local).Add(-time.Second)

	// Retrieve all procurements from the "procurements" table
	if err := n.db.Table("procurement").Where("start_date  BETWEEN ? AND ?", lastMonthStart, lastMonthEnd).Find(&procurements).Error; err != nil {
		return nil, err
	}

	return procurements, nil
}
