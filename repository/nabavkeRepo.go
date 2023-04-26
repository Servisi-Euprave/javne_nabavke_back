package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
	"os"
)

type NabavkePostgreSQL struct {
	l  *log.Logger
	db *gorm.DB
}

func PostgreSQLConnection(l *log.Logger) (*NabavkePostgreSQL, error) {
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
	return &NabavkePostgreSQL{l, db}, nil

}
func setup(db *gorm.DB) {
	db.AutoMigrate(&model.Nabavka{})
	db.AutoMigrate(&model.PlanJavneNabavke{})
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func (n NabavkePostgreSQL) AddNabavka(nabavka *model.Nabavka) error {
	//TODO implement me
	panic("implement me")
}
