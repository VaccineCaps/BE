package model

type User struct {
	ID       int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"`
	Image    string `json:"image"`
	RoleId   int    `gorm:"not null; default:2" json:"role_id"`
	Role     Role   `gorm:"ForeignKey:RoleId;references:ID;null" json:"-"`
}

func (*User) TableName() string {
	return "users"
}
