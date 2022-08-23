package rms

import (
	"fmt"
	"log"
	"testing"

	"github.com/nurja1218/go-learn/util"
	"github.com/stretchr/testify/assert"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	database = "rms_db"
	user     = "jayden"
	password = "jayden1"
)

func TestDevice(t *testing.T) {
	search := NewSearch()

	assert := assert.New(t)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, database)
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	devices := GetAllDeviceBySearch(db, search)
	assert.GreaterOrEqual(len(devices), 0)

	for _, device := range devices {
		util.PrintJson(device)
	}
}
