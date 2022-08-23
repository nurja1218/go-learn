package logis

import (
	"database/sql"
	"time"
)

type LogisAuditTrail struct {
	ID          int `gorm:"primaryKey"`
	CompanyId   int
	UserId      int
	Type        string
	RelatedKey  sql.NullString
	Description sql.NullString
	CreatedAt   time.Time
}

type LogisCompany struct {
	ID                int `gorm:"primaryKey"`
	Name              string
	Phone             string
	Email             string
	PlaceName         sql.NullString
	Address           string
	DepartureLocation sql.NullString
	DepartureAddress  sql.NullString
	CustomerCode      sql.NullString
	ClientIp          sql.NullString
	IsDeleted         int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type LogisDevice struct {
	Serial             string `gorm:"primaryKey"`
	CompanyId          int
	CurrentTransportId int
	DeviceStatus       string
	UsingCount         int
	RecentUptime       sql.NullInt64
	RecentUsedAt       sql.NullTime
	Banned             sql.NullInt16
	IsDeleted          int
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type LogisTransportMappingLog struct {
	ID        int `gorm:"primaryKey"`
	UserId    int
	CompanyId int
	Serial    sql.NullString
	LabelCode sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LogisTransport struct {
	ID                      int `gorm:"primaryKey"`
	Serial                  sql.NullString
	LabelCode               string
	CompanyId               int
	ProductName             string
	ProductQuantity         int
	BuyerName               string
	BuyerAddress            string
	TransportStatus         string
	DepartureTime           sql.NullTime
	DepartureLocation       sql.NullString
	DepartureAddress        sql.NullString
	ArrivalTime             sql.NullTime
	ArrivalLocation         sql.NullString
	ArrivalAddress          sql.NullString
	ManualArrivalTime       sql.NullTime
	Description             sql.NullString
	Temperature             sql.NullString
	TemperatureAfterArrival sql.NullString
	ReturnTemperature       sql.NullString
	ReturnArrivalTime       sql.NullTime
	CompletionTime          sql.NullTime
	Interval                int
	Mkt                     sql.NullFloat64
	RawData                 sql.NullString
	CreatedAt               time.Time
	UpdatedAt               time.Time
}

type LogisUser struct {
	ID                int `gorm:"primaryKey"`
	Name              string
	Phone             string
	Email             string
	Password          string
	Role              string
	CompanyId         int
	Department        sql.NullString
	IsDeleted         int
	PasswordUpdatedAt sql.NullTime
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
