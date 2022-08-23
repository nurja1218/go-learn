package copydb

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func CreateAuditTrails(localdb *gorm.DB, auditTrails []RMSAuditTrail) {
	if len(auditTrails) != 0 {
		localdb.AutoMigrate(&RMSAuditTrail{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(auditTrails)

	}
}

func CreateCompanies(localdb *gorm.DB, companies []RMSCompany) {
	if len(companies) != 0 {
		localdb.AutoMigrate(&RMSCompany{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(companies)
	}
}

func CreateCompanyMeetings(localdb *gorm.DB, companyMeetings []RMSCompanyMeeting) {
	if len(companyMeetings) != 0 {
		localdb.AutoMigrate(&RMSCompanyMeeting{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(companyMeetings)
	}
}

func CreateDevices(localdb *gorm.DB, devices []RMSDevice) {
	if len(devices) != 0 {
		localdb.AutoMigrate(&RMSDevice{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "serial"}},
			UpdateAll: true,
		}).Create(devices)
	}
}

func CreateDeviceSupplies(localdb *gorm.DB, deviceSupplies []RMSDeviceSupply) {
	if len(deviceSupplies) != 0 {
		localdb.AutoMigrate(&RMSDeviceSupply{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(deviceSupplies)

	}
}

func CreateDeviceTypes(localdb *gorm.DB, deviceTypes []RMSDeviceType) {
	if len(deviceTypes) != 0 {
		localdb.AutoMigrate(&RMSDeviceType{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(deviceTypes)
	}
}

func CreateUsers(localdb *gorm.DB, users []RMSUser) {
	if len(users) != 0 {
		localdb.AutoMigrate(&RMSUser{})
		localdb.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).Create(users)
	}
}
