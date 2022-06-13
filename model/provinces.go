package model

type Provinces struct {
	ID   int    `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name string `json:"name"`
}

func (*Provinces) TableName() string {
	return "provinces"
}
