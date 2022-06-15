package model

type Cities struct {
	ID   			int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name 			string 		`json:"name"`
	ProvincesId   	int    		`gorm:"not null; default:2" json:"province_id"`
	Province     	Provinces   `gorm:"ForeignKey:ProvincesId;references:ID;null" json:"-"`
}

func (*Cities) TableName() string {
	return "cities"
}
