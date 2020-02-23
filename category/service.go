package category

//Service is an interface provides categories methods/usecases
type Service interface {
	AddCategory(category *Category) error
	GetCategories() []*Category
	GetCategory(ID int) *Category
}
