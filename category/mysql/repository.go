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

func (r *repository) Store(category *domains.Category) domains.Response {
	query := `INSERT category SET category_name=?`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return domains.Response{
			Success: false,
			Message: "failed to prepare the SQL query statement of inserting category",
			Errors:  err.Error(),
			Payload: err.Error(),
		}
	}
	res, err := stmt.Exec(category.Name)
	if err != nil {
		return domains.Response{
			Success: false,
			Message: "failed to insert the category to the database",
			Errors:  err.Error(),
			Payload: err.Error(),
		}
	}
	return domains.Response{
		Success: true,
		Message: "successfully created a category",
		Errors:  nil,
		Payload: res,
	}
}
