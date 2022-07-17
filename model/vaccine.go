package model

type Vaccines struct {
	ID   int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name string `json:"name"`
}

func (*Vaccines) TableName() string {
	return "vaccines"
}
