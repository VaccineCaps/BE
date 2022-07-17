package model

type News struct {
	ID       		int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Title    		string 		`json:"title" validate:"required"`
	Context 		string 		`json:"context" validate:"required"`
}

func (*News) TableName() string {
	return "news"
}
