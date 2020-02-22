package category

import (
	"github.com/muathendirangu/lavida-api/domains"
)

//Service is an interface provides categories methods/usecases
type Service interface {
	AddCategory(category *domains.Category) error
}
