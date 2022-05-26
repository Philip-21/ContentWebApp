package database

import "gorm.io/gorm"

type testDBRepo struct {
	DB *gorm.DB
}
