package db

import (
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseManager interface {
	Initialize(dsn string, connection string) error
	GetDB() *gorm.DB
	StartTransaction() *gorm.DB
	CommitTransaction(tx *gorm.DB) *gorm.DB
	RollbackTransaction(tx *gorm.DB) *gorm.DB
}

func NewDatabaseManager() DatabaseManager {
	return &databaseManager{}
}

type databaseManager struct {
	DB *gorm.DB
}

func createDBInstance(dsn string, connection string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.S().Error(err.Error())
	}

	return db, nil
}

func (dbManager *databaseManager) Initialize(dsn string, connection string) error {
	var err error

	if dbManager.DB, err = createDBInstance(dsn, connection); err != nil {
		return err
	}

	return nil
}

func (dbManager *databaseManager) GetDB() *gorm.DB {
	if dbManager.DB == nil {
		return nil
	}
	return dbManager.DB
}

func (dbManager *databaseManager) StartTransaction() *gorm.DB {
	return dbManager.DB.Begin()
}

func (dbManager *databaseManager) CommitTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Commit()
}

func (dbManager *databaseManager) RollbackTransaction(tx *gorm.DB) *gorm.DB {
	return tx.Rollback()
}
