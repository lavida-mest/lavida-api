package mysql

import (
	"database/sql"
	"log"

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

func (r *repository) Search(Location, Duration, Traveler, Month, Year string) []*trip.Trip {
	query := `SELECT * FROM trip WHERE trip_location=? OR trip_duration=? OR traveler_type=? OR trip_month=? OR trip_year=?`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(Location, Duration, Traveler, Month, Year)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := make([]*trip.Trip, 0)
	for rows.Next() {
		trips := new(trip.Trip)
		err := rows.Scan(
			&trips.ID,
			&trips.Name,
			&trips.Location,
			&trips.Description,
			&trips.Activity,
			&trips.Price,
			&trips.Capacity,
			&trips.Month,
			&trips.Year,
			&trips.Duration,
			&trips.Type,
			&trips.Traveler,
			&trips.IsPriceOn,
			&trips.IsFull,
			&trips.Status,
			&trips.Guide,
		)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, trips)
	}
	return result
}
