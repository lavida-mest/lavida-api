package guide

//Guide of a trip
type Guide struct {
	ID       int64  `json:"tour_guide_id"`
	Name     string `json:"tour_guide_name"`
	Email    string `json:"tour_guide_email"`
	Number   string `json:"tour_guide_number"`
	Category int64  `json:"category_id"`
}

//Repository defines how to store Guide
type Repository interface {
	Store(guide *Guide) error
	Get() []*Guide
}
