package model

type User struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
type Role struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
