package person

import (
	"errors"
	"fmt"
	"os"
	"poc-auto-read-ec/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type IPersonRepository interface {
	InsertPersonToDB(list []models.GormPerson) error
}

type personRepository struct{}

func NewPersonRepository() IPersonRepository {
	return &personRepository{}
}

func (repo *personRepository) ConnectMSSQL() (*gorm.DB, error) {
	conn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		// table specific ??
	)

	db, err := gorm.Open(sqlserver.Open(conn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	fmt.Println("Connected to MSSQL DB successfully")
	return db, nil
}

func (repo *personRepository) InsertPersonToDB(list []models.GormPerson) error {

	if len(list) == 0 {
		return errors.New("list is empty")
	}
	db, err := repo.ConnectMSSQL()
	if err != nil {
		return err
	}

	if err := db.AutoMigrate(&models.GormPerson{}); err != nil {
		return fmt.Errorf("failed to migrate schema: %v", err)
	}

	if err := db.Create(&list).Error; err != nil {
		return fmt.Errorf("failed to insert data: %v", err)
	}

	fmt.Println("Inserted persons successfully")
	return nil

}
