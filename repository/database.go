package repository

import (
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"javne_nabavke_back/model"
	"log"
	"os"
)

type DatabaseError struct {
	Err     error
	Message string
}

func (repositoryError *DatabaseError) Error() string {
	return fmt.Sprintf("%s not found in query. %v\n", repositoryError.Message, repositoryError.Err)
}

func QueryNotFoundError(name string) error {
	return &DatabaseError{
		errors.New("Query not found"),
		fmt.Sprintf("%s not found\n", name),
	}
}

func CannotCreateError(name string) error {
	return &DatabaseError{
		errors.New("Cannot create"),
		fmt.Sprintf("Cannot create %s\n", name),
	}
}

func PostgreSQLConnection(l *log.Logger) (*gorm.DB, error) {
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
	return db, nil

}
func setup(db *gorm.DB) {
	db.AutoMigrate(&model.Procurement{})
	db.AutoMigrate(&model.ProcurementPlan{})
	db.AutoMigrate(&model.Offer{})
}
