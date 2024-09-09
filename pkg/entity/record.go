package entity

type Record struct {
	ID    int    `gorm:"primary_key;AUTO_INCREMENT"`
	Path  string `gorm:"index"`
	Value string
}
