package entities

// Book represents itself
type Book struct {
	ID    int64   `json:"id,omitempty"`
	Name  string  `json:"name"`
	ISBN  string  `json:"isbn,omitempty"`
	Price float64 `json:"price"`
}
