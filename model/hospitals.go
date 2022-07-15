package model

type Hospitals struct {
	ID   			int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name 			string 		`json:"name"`
	Address 		string 		`json:"address"`
	Email 			string 		`json:"email"`
	UserId			int 		`gorm:"not null" json:"user_id"`
	CitiesId   		int    		`gorm:"not null" json:"cities_id"`
	User     		User	 	`gorm:"ForeignKey:UserId;references:ID;null" json:"-"`
	Cities	     	Cities 		`gorm:"ForeignKey:CitiesId;references:ID;null" json:"-"`
}

func (*Hospitals) TableName() string {
	return "hospitals"
}
