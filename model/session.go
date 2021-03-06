package model

import "time"

type Session struct {
	ID            int       `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	HospitalId    int       `gorm:"not null" json:"hospital_id"`
	Hospital      Hospitals `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	VaccineId     int       `gorm:"not null" json:"vaccine_id"`
	Vaccine       Vaccines  `gorm:"ForeignKey:VaccineId;references:ID;null" json:"-"`
	Sesi          string    `gorm:"not null" json:"sesi"`
	MaxSession    int       `gorm:"not null" json:"max_session"`
	NumberBooking int       `json:"number_booking"`
	Tanggal       time.Time `json:"tanggal"`
}

func (*Session) TableName() string {
	return "session"
}
