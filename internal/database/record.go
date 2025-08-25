package database

// Record 表示系统中的记录实体
type Record struct {
	Path  string `gorm:"primaryKey"`
	Value string
}
