package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"auzy-api/config"
	"auzy-api/model"
)

func main() {
	config.AutoReadConfig("local.env")
	workerConfig := config.NewAPIConfig()

	db, _ := gorm.Open(mysql.Open(workerConfig.MySQLDSN))
	dbMigrator := db.Migrator()
	dbMigrator.CreateTable(model.Role{})
	dbMigrator.CreateTable(model.Permission{})
	dbMigrator.CreateTable(model.RolesHasPermission{})
	dbMigrator.CreateTable(model.Staff{})
}
