package model

import "time"

type VaccineTransactionsIn struct {
	ID                int              `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	HospitalId        int              `gorm:"not null" json:"hospital_id"`
	Hospital          Hospitals        `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	VaccineHospitalId int              `gorm:"not null" json:"vaccinehospital_id"`
	VaccineHospital   VaccineHospitals `gorm:"ForeignKey:VaccineHospitalId;references:ID;null" json:"-"`
	AsalVaccine       string           `json:"asalvaccine"`
	Status            bool             `gorm:"default: false" json:"status"`
	Tanggal           time.Time        `json:"tanggal"`
	NoTransaction     int              `json:"no_transaction"`
	Distributor       string           `json:"distributor"`
}

func (*VaccineTransactionsIn) TableName() string {
	return "vaccine_transactions_in"
}
