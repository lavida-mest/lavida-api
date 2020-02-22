package category

import "github.com/muathendirangu/lavida-api/domains"

//Service is an interface provides categories methods/usecases
type Service interface {
	Store(name string) domains.Response
}
