package project

import (
	"database/sql"
	"github.com/alx-t/go-playground/models"
)

func ReadOne(db *sql.DB, id int) (models.Project, error) {
	var rec models.Project
	row := db.QueryRow("SELECT id, name, data FROM projects WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.Name, &rec.Data)
}

func Read(db *sql.DB, str string) ([]models.Project, error) {
	var rows *sql.Rows
	var err error
	if str != "" {
		rows, err = db.Query("SELECT * FROM projects WHERE name LIKE $1 ORDER BY id",
			"%"+str+"%")
	} else {
		rows, err = db.Query("SELECT * FROM projects ORDER BY id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs = make([]models.Project, 0)
	var rec models.Project
	for rows.Next() {
		if err = rows.Scan(&rec.Id, &rec.UserId, &rec.Name, &rec.Data, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}
