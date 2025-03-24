package schemas

// An example schematic for database tables
type Product struct {
	ID          int64   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name        string  `json:"name" gorm:"size:255;not null"`
	Description string  `json:"description" gorm:"type:text"`
	Price       float64 `json:"price" gorm:"not null"`
	Stock       int64   `json:"stock" gorm:"not null"`
	CreatedAt   string  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   string  `json:"updated_at" gorm:"autoUpdateTime"`
}
