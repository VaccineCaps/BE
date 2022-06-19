package model

type Hospitals struct {
	ID   			int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name 			string 		`json:"name"`
	Address 		string 		`json:"address"`
	UserId			int 		`gorm:"not null" json:"user_id"`
	ProvincesId   	int    		`gorm:"not null" json:"province_id"`
	User     		User	 	`gorm:"ForeignKey:UserId;references:ID;null" json:"-"`
	Province     	Provinces   `gorm:"ForeignKey:ProvincesId;references:ID;null" json:"-"`
}

func (*Hospitals) TableName() string {
	return "hospitals"
}
