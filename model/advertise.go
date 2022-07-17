package model

type Advertise struct {
	ID       		int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Image    		string 		`json:"image" validate:"required"`
}

func (*Advertise) TableName() string {
	return "advertise"
}
