package model

type User struct {
	ID       int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Username string `json:"Username" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"-" validate:"required"`
	Image    string `json:"image"; gorm: "default:https://w1.pngwing.com/pngs/726/597/png-transparent-graphic-design-icon-customer-service-avatar-icon-design-call-centre-yellow-smile-forehead.png"`
	RoleId   int    `gorm:"not null; default:2" json:"role_id"`
	Role     Role   `gorm:"ForeignKey:RoleId;references:ID;null" json:"-"`
}

func (*User) TableName() string {
	return "users"
}
