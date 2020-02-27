package preference

//Service or Usecase is an interface provides preference method definitions
type Service interface {
	AddPreference(preference *Preference) error
}
