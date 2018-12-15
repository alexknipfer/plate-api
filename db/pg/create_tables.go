package pg

import (
	"github.com/knipferrc/plate-api/app/models"
	"github.com/knipferrc/plate-api/db/gorm"
)

// CreateTables sets up DB tables
func CreateTables() {
	gorm.DBCon().AutoMigrate(&models.User{})
}
