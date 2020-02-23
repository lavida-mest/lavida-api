package mysql

import (
	"database/sql"
	"log"

	"github.com/muathendirangu/lavida-api/category"
)

type repository struct {
	conn *sql.DB
}

//New creates a new instance of repository
func New(conn *sql.DB) category.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Store(category *category.Category) error {
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

func (r *repository) Get() []*category.Category {
	query := `SELECT * FROM trip_category`
	rows, err := r.conn.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := make([]*category.Category, 0)
	for rows.Next() {
		categories := new(category.Category)
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

func (r *repository) GetByID(ID int) *category.Category {
	var category = category.Category{}
	query := `SELECT * FROM trip_category WHERE category_id=?`
	err := r.conn.QueryRow(query, ID).Scan(&category.ID, &category.Name)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no category with ID %v", ID)
	case err != nil:
		log.Fatalf("an error %v occurred", err)

	}
	return &category
}
