package mysql

import (
	"database/sql"

	"github.com/muathendirangu/lavida-api/guide"
)

type repository struct {
	conn *sql.DB
}

//New creates a new instance of guide mysql repository
func New(conn *sql.DB) guide.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Store(guide *guide.Guide) error {
	query := `INSERT INTO guide (tour_guide_name, tour_guide_email, tour_guide_number, category) VALUES(?,?,?,?)`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(guide.Name, guide.Email, guide.Number, guide.Category)
	if err != nil {
		return err
	}
	lasTG, err := res.LastInsertId()
	guide.ID = lasTG
	return nil
}
