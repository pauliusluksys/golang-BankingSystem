package domain

import (
	"gorm.io/gorm"
)

type DbGormConn struct {
	db *gorm.DB
}

func DbGormConnection(db *gorm.DB) DbGormConn {
	return DbGormConn{db}
}
