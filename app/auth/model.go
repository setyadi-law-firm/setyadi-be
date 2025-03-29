package auth

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Email    string `gorm:"uniqueIndex" json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}
