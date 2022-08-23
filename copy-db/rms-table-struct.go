package copydb

import (
	"database/sql"
	"time"
)

type RMSAuditTrail struct {
	ID          int `gorm:"primaryKey"`
	UserId      int
	Type        string
	RelatedKey  sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
}

type RMSCompany struct {
	ID               int `gorm:"primaryKey"`
	ServiceType      sql.NullString
	ServiceCompanyId sql.NullInt64
	WillogManagerId  int
	Name             string
	Address          string
	Phone            string
	ManagerName      string
	CompanyType      string
	TotalUptime      sql.NullInt64
	IsDeleted        int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type RMSCompanyMeeting struct {
	ID             int `gorm:"primaryKey"`
	CompanyId      int
	ManagerId      int
	ProgressedDate sql.NullTime
	MinutesUrl     sql.NullString
	IsDeleted      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type RMSDevice struct {
	Serial            string `gorm:"primaryKey"`
	CurrentCompanyId  sql.NullInt64
	DeviceTypeId      int
	CalibrationDate   sql.NullTime
	CalibrationNumber sql.NullString
	ArrivalUptime     sql.NullInt64
	BatteryLevel      sql.NullInt64
	Description       sql.NullString
	IsDeleted         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type RMSDeviceSupply struct {
	ID           int `gorm:"primaryKey"`
	CompanyId    int
	DeviceSerial string
	SupplyDate   sql.NullTime
	ReturnDate   sql.NullTime
	IsDeleted    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type RMSDeviceType struct {
	ID              int `gorm:"primaryKey"`
	DeviceName      sql.NullString
	SensingType     sql.NullString
	FirmwareVersion sql.NullString
	IsDeleted       int
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type RMSUser struct {
	IsDeleted    int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ID           int `gorm:"primaryKey"`
	Name         string
	Phone        string
	Email        string
	Password     string
	Department   string
	Role         string
	IsSuperAdmin int
}

type Tabler interface {
	TableName() string
}

func (RMSAuditTrail) TableName() string {
	return "audit-trail"
}

func (RMSCompany) TableName() string {
	return "company"
}

func (RMSCompanyMeeting) TableName() string {
	return "company-meeting"
}

func (RMSDevice) TableName() string {
	return "device"
}

func (RMSDeviceSupply) TableName() string {
	return "device-supply"
}

func (RMSDeviceType) TableName() string {
	return "device-type"
}

func (RMSUser) TableName() string {
	return "user"
}
