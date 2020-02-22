package domains

//Category of trip
type Category struct {
	ID   int64  `json:"category_id"`
	Name string `json:"category_name"`
}

//Repository defines how to store category of
type Repository interface {
	Store(c *Category) error
}
