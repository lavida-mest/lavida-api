package guide

//Service or Usecase is an interface provides guides methods
type Service interface {
	AddGuide(guide *Guide) error
}
