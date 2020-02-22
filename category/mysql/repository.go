package mysql

import (
	"database/sql"
	"log"

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

func (r *repository) Get() []*domains.Category {
	query := `SELECT * FROM trip_category`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := make([]*domains.Category, 0)
	for rows.Next() {
		categories := new(domains.Category)
		err := rows.Scan(
			&categories.ID,
			&categories.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, categories)
	}
	return result
}
