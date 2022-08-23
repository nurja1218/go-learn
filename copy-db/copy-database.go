package copydb

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type TableInDB struct {
	auditTrails     []RMSAuditTrail
	companies       []RMSCompany
	companyMeetings []RMSCompanyMeeting
	devices         []RMSDevice
	deviceSupplies  []RMSDeviceSupply
	deviceTypes     []RMSDeviceType
	users           []RMSUser
}

const (
	host     = "127.0.0.1"
	database = "copy_db"
	user     = "jayden"
	password = "jayden1"
)

func CopyDatabase(db *gorm.DB) TableInDB {
	serverdb_data := TableInDB{}
	serverdb_data.auditTrails = append(serverdb_data.auditTrails, RMS_GetAllAuditTrail(db)...)
	serverdb_data.companies = append(serverdb_data.companies, RMS_GetAllCompany(db)...)
	serverdb_data.companyMeetings = append(serverdb_data.companyMeetings, RMS_GetAllCompanyMeeting(db)...)
	serverdb_data.devices = append(serverdb_data.devices, RMS_GetAllDevice(db)...)
	serverdb_data.deviceSupplies = append(serverdb_data.deviceSupplies, RMS_GetAllDeviceSupply(db)...)
	serverdb_data.deviceTypes = append(serverdb_data.deviceTypes, RMS_GetAllDeviceType(db)...)
	serverdb_data.users = append(serverdb_data.users, RMS_GetAllUser(db)...)

	lacaldb_data := PasteDatabase(&serverdb_data, db)

	return lacaldb_data
}

func PasteDatabase(serverdb_data *TableInDB, serverdb *gorm.DB) TableInDB {
	localdb_data := TableInDB{}

	var dsn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, database)
	localdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	CreateAuditTrails(localdb, serverdb_data.auditTrails)
	CreateCompanies(localdb, serverdb_data.companies)
	CreateCompanyMeetings(localdb, serverdb_data.companyMeetings)
	CreateDevices(localdb, serverdb_data.devices)
	CreateDeviceSupplies(localdb, serverdb_data.deviceSupplies)
	CreateDeviceTypes(localdb, serverdb_data.deviceTypes)
	CreateUsers(localdb, serverdb_data.users)

	return localdb_data
}
