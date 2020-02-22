package mysql

import (
	"database/sql"

	"github.com/muathendirangu/lavida-api/domains"
)

type repository struct {
	conn *sql.DB
}

//New creates a new instance of repository
func New(conn *sql.DB) domains.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Store(category *domains.Category) error {
	query := `INSERT trip_category SET category_name=?`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(category.Name)
	if err != nil {
		return err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	category.ID = lastID
	return nil
}
