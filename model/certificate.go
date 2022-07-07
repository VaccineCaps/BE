package model

type Certificate struct {
	ID   			int   	`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Image		 	string 	`json:"image"`
}

func (*Certificate) TableName() string {
	return "certificate"
}
