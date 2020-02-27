package preference

//Preference Entity
type Preference struct {
	ID    int64 `json:"preference_id"`
	Trip  int64 `json:"trip_id"`
	Guide int64 `json:"tour_guide_id"`
}

//Repository defines methods of how to store a trip
type Repository interface {
	Add(preference *Preference) error
}
