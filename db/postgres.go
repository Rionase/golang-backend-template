package db

import (
	"fmt"
	"golang-backend-template/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct{}

func NewPostgresDB() *Postgres {
	return &Postgres{}
}

func (p *Postgres) Connect(creds *model.DatabaseCredential) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", creds.Host, creds.Username, creds.Password, creds.DatabaseName, creds.Port)

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return dbConn, nil
}

func (p *Postgres) Reset(db *gorm.DB, table string) error {
	if err := db.Exec("TRUNCATE " + table).Error; err != nil {
		return err
	}
	return nil
}
