package logis

import (
	"database/sql"
	"log"
	"time"
)

// **TABLE 컬럼 순서와 struct구조가 완전히 동일해야함**
type Device struct {
	Serial             string        `db:"serial" json:"serial"`
	CompanyId          int           `db:"company_id" json:"company_id"`
	CurrentTransportId int           `db:"current_transport_id" json:"current_transport_id"`
	DeviceStatus       string        `db:"device_status" json:"device_status"`
	UsingCount         int           `db:"using_count" json:"using_count"`
	RecentUptime       sql.NullInt64 `db:"recent_uptime" json:"recent_uptime"`
	RecentUsedAt       sql.NullTime  `db:"recent_used_at" json:"recent_used_at"`
	Banned             sql.NullInt16 `db:"banned" json:"banned"`
	IsDeleted          int           `db:"is_deleted" json:"is_deleted"`
	CreatedAt          time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time     `db:"updated_at" json:"updated_at"`
}

func GetAllDevice(db *sql.DB) []Device {
	rows, err := db.Query("SELECT * FROM device")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	devices := []Device{}
	for rows.Next() {
		var row Device
		err := rows.Scan(&row.Serial, &row.CompanyId, &row.CurrentTransportId,
			&row.DeviceStatus, &row.UsingCount, &row.RecentUptime,
			&row.RecentUsedAt, &row.Banned, &row.IsDeleted,
			&row.CreatedAt, &row.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		devices = append(devices, row)
	}
	return devices
}
