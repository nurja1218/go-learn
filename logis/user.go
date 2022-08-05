package logis

import (
	"database/sql"
	"log"
	"time"
)

type User struct {
	Id                int            `db:"id" json:"id"`
	Name              string         `db:"name" json:"name"`
	Phone             string         `db:"phone" json:"phone"`
	Email             string         `db:"email" json:"email"`
	Password          string         `db:"password" json:"password"`
	Role              string         `db:"role" json:"role"`
	CompanyId         int            `db:"company_id" json:"company_id"`
	Department        sql.NullString `db:"department" json:"department"`
	IsDeleted         int            `db:"is_deleted" json:"is_deleted"`
	PasswordUpdatedAt sql.NullTime   `db:"password_updated_at" json:"password_updated_at"`
	CreatedAt         time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt         time.Time      `db:"updated_at" json:"updated_at"`
}

func GetAllUser(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM user")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	users := []User{}
	for rows.Next() {
		var row User
		err := rows.Scan(&row.Id, &row.Name, &row.Phone, &row.Email,
			&row.Password, &row.Role, &row.CompanyId, &row.Department,
			&row.IsDeleted, &row.PasswordUpdatedAt, &row.CreatedAt, &row.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		users = append(users, row)
	}
	return users
}
