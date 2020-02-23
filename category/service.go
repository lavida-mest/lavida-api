package category

//Service or Usecase is an interface provides categories methods
type Service interface {
	AddCategory(category *Category) error
	GetCategories() []*Category
	GetCategory(ID int) *Category
}
