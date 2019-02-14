package project

import (
	"database/sql"
	"github.com/alx-t/go-playground/models"
)

func ReadOne(db *sql.DB, id int) (models.Project, error) {
	var rec models.Project
	row := db.QueryRow("SELECT id, user_id, name, data, created_at, updated_at FROM projects WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.UserId, &rec.Name, &rec.Data, &rec.CreatedAt, &rec.UpdatedAt)
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

func Insert(db *sql.DB, prj *models.Project) (sql.Result, error) {
	return db.Exec("INSERT INTO projects VALUES (default, $1, $2, $3)",
		prj.UserId, prj.Name, prj.Data)
}

func Remove(db *sql.DB, id int) (sql.Result, error) {
	return db.Exec("DELETE FROM projects WHERE id=$1", id)
}

func Update(db *sql.DB, id int, prj *models.Project) (sql.Result, error) {
	return db.Exec("UPDATE projects SET user_id = $1, name = $2, data = $3, updated_at = now() WHERE id=$4",
		prj.UserId, prj.Name, prj.Data)
}
