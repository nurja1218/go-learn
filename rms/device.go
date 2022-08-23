package rms

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type Device struct {
	Serial            string `gorm:"primaryKey"`
	CurrentCompanyId  sql.NullInt64
	DeviceTypeId      sql.NullInt64
	CalibrationDate   sql.NullTime
	CalibrationNumber sql.NullString
	ArrivalUptime     sql.NullInt64
	BatteryLevel      sql.NullInt64
	Description       sql.NullString
	IsDeleted         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type DeviceSupply struct {
	ID           uint `gorm:"primaryKey"`
	CompanyId    uint
	DeviceSerial string
	SupplyDate   sql.NullTime
	ReturnDate   sql.NullTime
	IsDeleted    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DeviceType struct {
	ID              uint `gorm:"primaryKey"`
	DeviceName      sql.NullString
	SensingType     sql.NullString
	FirmwareVersion sql.NullString
	IsDeleted       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type DeviceTable struct {
	Serial            string `gorm:"primaryKey"`
	CurrentCompanyId  sql.NullInt64
	DeviceTypeId      sql.NullInt64
	CalibrationDate   sql.NullTime
	CalibrationNumber sql.NullString
	ArrivalUptime     sql.NullInt64
	BatteryLevel      sql.NullInt64
	Description       sql.NullString
	IsDeleted         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeviceType        DeviceType `gorm:"embedded"`
}

func GetAllDeviceBySearch(db *gorm.DB, search Search) []DeviceTable {
	// textSearch := "true"
	// if search.Text != "" {
	// 	textSearch = fmt.Sprintf("(serial like concat('%','%s','%') or name like concat('%','%s','%') or calibration_number like concat('%','%s','%') or firmware_version like concat('%','%s','%'))", search.Text, search.Text, search.Text, search.Text)
	// }
	// dateRangeSearch := "true"
	// if search.DateRange.from != "" && search.DateRange.to != "" {
	// 	dateRangeSearch = fmt.Sprintf("DATE_FORMAT(calibration_date, '%Y-%m-%d %H:%i:%s') between %s and %s", search.DateRange.from, search.DateRange.to)
	// }

	deviceTables := []DeviceTable{}
	err := db.Limit(5).Table("device").Select("device.*, device_type.*").Joins("left join device_type on device.device_type_id = device_type.id").Find(&deviceTables).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}

	return deviceTables

}
