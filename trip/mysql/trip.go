package mysql

import (
	"database/sql"

	"github.com/muathendirangu/lavida-api/trip"
)

type repository struct {
	conn *sql.DB
}

//New creates a new instance of trip mysql repository
func New(conn *sql.DB) trip.Repository {
	return &repository{
		conn: conn,
	}
}

func (r *repository) Store(trip *trip.Trip) error {
	query := `INSERT INTO trip (trip_name, trip_location, trip_description, trip_activity, trip_price, trip_capacity, 
		trip_month, trip_year, trip_duration, trip_type, traveler_type, price_visibilty, trip_availability, 
		trip_status,tour_guide) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(trip.Name, trip.Location, trip.Description, trip.Activity, trip.Price,
		trip.Capacity, trip.Month, trip.Year, trip.Duration, trip.Type, trip.Traveler, trip.IsPriceOn,
		trip.IsPriceOn, trip.Status, trip.Guide)
	if err != nil {
		return err
	}
	lasTP, err := res.LastInsertId()
	trip.ID = lasTP
	return nil
}
