package trip

//Service or Usecase is an interface provides trips methods
type Service interface {
	AddTrip(trip *Trip) error
	SearchTrip(Traveler, Month, Activity, Category string) []*Trip
	ViewTrip(ID, Guide int64) *Trip
	GetTrips() []*Trip
}
