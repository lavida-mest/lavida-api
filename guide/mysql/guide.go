package mysql

import (
	"database/sql"
	"log"

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
	query := `INSERT INTO guide(
						tour_guide_name,
						tour_guide_email,
						tour_guide_number,
						category_id
					)
					VALUES(?, ?, ?, ?)`
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

func (r *repository) Get() []*guide.Guide {
	query := `SELECT * FROM guide`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	var guides = make([]*guide.Guide, 0)

	for rows.Next() {
		guide := new(guide.Guide)
		err := rows.Scan(
			&guide.ID,
			&guide.Name,
			&guide.Email,
			&guide.Number,
			&guide.Category,
		)
		if err != nil {
			log.Fatal(err)
		}
		guides = append(guides, guide)
	}
	defer rows.Close()
	return guides
}
