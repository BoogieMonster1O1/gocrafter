package data

type User struct {
	DisplayName string `db:"display_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Role        string `db:"role"`
	ID          int64  `db:"id"`
}
