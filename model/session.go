package model

type Session struct {
	ID         int       `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	HospitalId int       `gorm:"not null" json:"hospital_id"`
	Hospital   Hospitals `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	VaccineId  int       `gorm:"not null" json:"vaccine_id"`
	Vaccine    Vaccines  `gorm:"ForeignKey:VaccineId;references:ID;null" json:"-"`
	Waktu      string    `gorm:"not null" json:"waktu"`
	MaxSession int       `gorm:"not null" json:"max_session"`
	Jadwal     string    `json:"jadwal"`
}

func (*Session) TableName() string {
	return "session"
}
