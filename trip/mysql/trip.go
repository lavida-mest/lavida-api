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

func (r *repository) Search(Traveler, Month, Activity string) []*trip.Trip {
	query := `SELECT tr.trip_id, tr.trip_name, tr.trip_location, tr.trip_month, tr.trip_year, tr.trip_price , tr.tour_guide, guide.tour_guide_name FROM trip AS tr
		 INNER JOIN guide ON tr.tour_guide=guide.tour_guide_id WHERE traveler_type=? OR trip_month=? OR trip_activity=?`
	stmt, err := r.conn.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(Traveler, Month, Activity)
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
			&trips.Month,
			&trips.Year,
			&trips.Price,
			&trips.Guide,
			&trips.Activity,
		)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, trips)
	}
	return result
}

func (r *repository) View(ID, Guide int64) *trip.Trip {
	var trip = trip.Trip{}
	query := `SELECT trip_id, trip_name, trip_location, trip_description, trip_price, tour_guide_name, trip.tour_guide, trip_category.category_name, 
	trip_activity FROM trip INNER JOIN guide ON trip.tour_guide = guide.tour_guide_id 
	INNER JOIN trip_category ON guide.category_id=trip_category.category_id WHERE trip_id=? AND guide.tour_guide_id=?;`
	err := r.conn.QueryRow(query, ID, Guide).Scan(
		&trip.ID,
		&trip.Name,
		&trip.Location,
		&trip.Description,
		&trip.Price,
		&trip.Month,    //references tour_guide_name
		&trip.Capacity, //references tour_guide
		&trip.Status,   //references category_name
		&trip.Activity,
	)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("the criteria you choose does not exist with ID %v", ID)
	case err != nil:
		log.Fatalf("an error %v occurred", err)

	}
	return &trip
}

func (r *repository) Get() []*trip.Trip {
	query := `SELECT trip_id, trip_name,trip_description, trip_duration, trip_month,
	trip_year, trip_location, trip_activity, tour_guide, trip_type FROM trip`
	rows, err := r.conn.Query(query)
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
			&trips.Description,
			&trips.Duration,
			&trips.Month,
			&trips.Year,
			&trips.Location,
			&trips.Activity,
			&trips.Guide,
			&trips.Type,
		)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, trips)
	}
	return result
}
