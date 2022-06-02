package model

type User struct {
	ID       int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"`
	RoleId   int    `gorm:"not null" json:"role_id"`
	Role     Role   `gorm:"ForeignKey:RoleId;references:ID;null" json:"-"`
}

func (*User) TableName() string {
	return "users"
}
