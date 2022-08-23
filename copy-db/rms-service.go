package copydb

import "gorm.io/gorm"

func RMS_GetAllAuditTrail(db *gorm.DB) []RMSAuditTrail {
	var auditTrails []RMSAuditTrail
	db.Table("audit_trail").Find(&auditTrails)
	return auditTrails
}

func RMS_GetAllCompany(db *gorm.DB) []RMSCompany {
	var companies []RMSCompany
	db.Table("company").Find(&companies)
	return companies
}

func RMS_GetAllCompanyMeeting(db *gorm.DB) []RMSCompanyMeeting {
	var companyMeetings []RMSCompanyMeeting
	db.Table("company_meeting").Find(&companyMeetings)
	return companyMeetings
}

func RMS_GetAllDevice(db *gorm.DB) []RMSDevice {
	var devices []RMSDevice
	db.Table("device").Find(&devices)
	return devices
}

func RMS_GetAllDeviceSupply(db *gorm.DB) []RMSDeviceSupply {
	var deviceSupplies []RMSDeviceSupply
	db.Table("device_supply").Find(&deviceSupplies)
	return deviceSupplies
}

func RMS_GetAllDeviceType(db *gorm.DB) []RMSDeviceType {
	var deviceTypes []RMSDeviceType
	db.Table("device_type").Find(&deviceTypes)
	return deviceTypes
}

func RMS_GetAllUser(db *gorm.DB) []RMSUser {
	var users []RMSUser
	db.Table("user").Find(&users)
	return users
}
