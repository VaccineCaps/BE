package model

import (
	"time"
)

type OtherPerson struct {
	ID       				int    			`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Name    				string 			`json:"name" validate:"required"`
	Placeofbirth    		string 			`json:"placeofbirth" validate:"required"`
	Dateofbirth 			time.Time 		`json:"dateofbirth" validate:"required"`
	Address		    		string 			`json:"address" validate:"required"`
	Phonenumber	    		string 			`json:"phonenumber" validate:"required"`
	NIK			    		int 			`json:"nik" validate:"required"`
	Email		    		string 			`json:"email" validate:"required"`
	VaccineStatus			int 			`json:"vaccinestatus" validate:"required"`
	UserId   				int    			`gorm:"not null" json:"user_id"`
	User     				User   			`gorm:"ForeignKey:UserId;references:ID;null" json:"-"`
}

func (*OtherPerson) TableName() string {
	return "otherperson"
}
