package model

import "time"

type VaccineTransactions struct {
	ID            int                `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	HospitalId    int                `gorm:"not null" json:"hospital_id"`
	Hospital      Hospitals          `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	VaccineId     int                `gorm:"not null" json:"vaccine_id"`
	Vaccine       Vaccines           `gorm:"ForeignKey:VaccineId;references:ID;null" json:"-"`
	Status        StatusTransactions `gorm:"not null" json:"status"`
	Tanggal       time.Time          `json:"tanggal"`
	NoTransaction int                `gorm:"not null" json:"no_transaction"`
	Distributor   string             `gorm:"not null" json:"distributor"`
}

type StatusTransactions string

const (
	Masuk  StatusTransactions = "Masuk"
	Keluat StatusTransactions = "Keluar"
)

func (*VaccineTransactions) TableName() string {
	return "vaccine_transactions"
}
