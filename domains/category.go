package domains

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Response formats the returned payload
type Response struct {
	Success   bool        `json:"success:omitempty"`
	Message   string      `json:"message,omitempty"`
	Timestamp time.Time   `json:"timestamp,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	Payload   interface{} `json:"payload,omitempty"`
}

//Category of trip
type Category struct {
	ID   int64  `json:"category_id"`
	Name string `json:"category_name"`
	gorm.Model
}

//Repository defines how to store category of
type Repository interface {
	Store(name string) Response
}
