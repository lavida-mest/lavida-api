package trip

//Service or Usecase is an interface provides trips methods
type Service interface {
	AddTrip(trip *Trip) error
}
