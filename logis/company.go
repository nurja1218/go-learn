package logis

import (
	"database/sql"
	"log"
	"time"
)

type Company struct {
	Id                int            `db:"id" json:"id"`
	Name              string         `db:"name" json:"name"`
	Phone             string         `db:"phone" json:"phone"`
	Email             string         `db:"email" json:"email"`
	PlaceName         sql.NullString `db:"place_name" json:"place_name"`
	Address           string         `db:"address" json:"address"`
	DepartureLocation sql.NullString `db:"departure_location" json:"departure_location"`
	DepartureAddress  sql.NullString `db:"departure_address" json:"departure_address"`
	CustomerCode      sql.NullString `db:"customer_code" json:"customer_code"`
	ClientIp          sql.NullString `db:"client_ip" json:"client_ip"`
	IsDeleted         int            `db:"is_deleted" json:"is_deleted"`
	CreatedAt         time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updated_at"`
}

func GetAllCompany(db *sql.DB) []Company {
	rows, err := db.Query("SELECT * FROM company")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	companies := []Company{}
	for rows.Next() {
		var row Company
		err := rows.Scan(&row.Id, &row.Name, &row.Phone, &row.Email,
			&row.PlaceName, &row.Address, &row.DepartureLocation,
			&row.DepartureAddress, &row.CustomerCode, &row.ClientIp,
			&row.IsDeleted, &row.CreatedAt, &row.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		companies = append(companies, row)
	}
	return companies
}
