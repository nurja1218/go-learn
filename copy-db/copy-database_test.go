package copydb

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Name                  string   `json: name`
	Type                  string   `json: type`
	Host                  string   `json: host`
	Port                  int      `json: port`
	Username              string   `json: username`
	Password              string   `json: password`
	Database              string   `json: database`
	Entities              []string `json: entities`
	Synchronize           bool     `json: synchronize`
	Logging               []string `json: logging`
	MaxQueryExecutionTime int      `json: maxQueryExecutionTime`
}

func TestCopyDatabase(t *testing.T) {
	production_db, err := os.Open("../rmsconfig.json")
	if err != nil {
		fmt.Println(err)
	}
	defer production_db.Close()

	byteValue, _ := ioutil.ReadAll(production_db)

	var config Config
	json.Unmarshal(byteValue, &config)

	assert := assert.New(t)

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.Username, config.Password, config.Host, config.Database)
	db, err := gorm.Open(gorm_mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	copy_data := CopyDatabase(db)

	assert.GreaterOrEqual(len(copy_data.users), 0)
	assert.GreaterOrEqual(len(copy_data.companies), 0)
	assert.GreaterOrEqual(len(copy_data.devices), 0)
	assert.GreaterOrEqual(len(copy_data.companyMeetings), 0)
	assert.GreaterOrEqual(len(copy_data.deviceSupplies), 0)
	assert.GreaterOrEqual(len(copy_data.deviceTypes), 0)
	assert.GreaterOrEqual(len(copy_data.auditTrails), 0)

	fmt.Print("\n\n\n\nSuccess!!!\n\n\n\n")
}
