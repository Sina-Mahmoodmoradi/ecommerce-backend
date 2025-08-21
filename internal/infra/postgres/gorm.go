package postgres

import (
	"fmt"

	"github.com/Sina-Mahmoodmoradi/ecommerce-backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)






func NewGorm(cfg config.Config)(*gorm.DB, error){
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode) 


	return gorm.Open(postgres.Open(dsn),&gorm.Config{})

}


