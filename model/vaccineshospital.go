package model

type VaccineHospitals struct {
	ID         int       `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	HospitalId int       `gorm:"not null" json:"hospital_id"`
	Hospital   Hospitals `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	VaccineId  int       `gorm:"not null" json:"vaccine_id"`
	Vaccine    Vaccines  `gorm:"ForeignKey:VaccineId;references:ID;null" json:"-"`
	Stok       int       `json:"stok"`
}

func (*VaccineHospitals) TableName() string {
	return "vaccine_hospitals"
}
