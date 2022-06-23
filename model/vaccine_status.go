package model

type VaccineStatus struct {
	ID   			int   	`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Vaccine_number 	string 	`json:"vaccinenumber"`
}

func (*VaccineStatus) TableName() string {
	return "vaccinestatus"
}
