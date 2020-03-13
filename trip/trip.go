package trip

// Trip Entity
type Trip struct {
	ID          int64   `json:"trip_id"`
	Name        string  `json:"trip_name"`
	Location    string  `json:"trip_location"`
	Description string  `json:"trip_description"`
	Activity    string  `json:"trip_activity"`
	Price       float64 `json:"trip_price"`
	Capacity    int     `json:"trip_capacity"`
	Month       string  `json:"trip_month"`
	Year        string  `json:"trip_year"`
	Duration    string  `json:"trip_duration"`
	Type        string  `json:"trip_type"`
	Traveler    string  `json:"traveler_type"`
	IsPriceOn   bool    `json:"price_visibilty"`
	IsFull      bool    `json:"trip_availability"`
	Status      string  `json:"trip_status"`
	Guide       int64   `json:"tour_guide"`
}

//Repository defines how to store Trip
type Repository interface {
	Store(trip *Trip) error
	Search(Traveler, Month, Activity string) []*Trip
	View(ID, Guide int64) *Trip
	Get() []*Trip
}
