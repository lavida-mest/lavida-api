package mysql

import (
	"database/sql"

	"github.com/muathendirangu/lavida-api/guide"
)

type guideRepo struct {
	conn *sql.DB
}

//NewGuideRepository creates a new instance of guide mysql repository
func NewGuideRepository(conn *sql.DB) guide.Repository {
	return &guideRepo{
		conn: conn,
	}
}

func (r *guideRepo) Store(guide *guide.Guide) error {
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
