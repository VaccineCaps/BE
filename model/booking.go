package model

type Booking struct {
	ID              int           `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	UserId          int           `gorm:"not null" json:"user_id"`
	User            User          `gorm:"ForeignKey:UserId;references:ID;null" json:"-"`
	HospitalId      int           `gorm:"not null" json:"hospital_id"`
	Hospital        Hospitals     `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
	SessionId       int           `gorm:"not null" json:"session_id"`
	Session         Session       `gorm:"ForeignKey:SessionId;references:ID;null" json:"-"`
	CertificateId 	int           `gorm:"not null" json:"certificate_id"`
	Certificate   	Certificate   `gorm:"ForeignKey:CertificateId;references:ID;null" json:"-"`
	BookedCode      string        `json:"bookedcode"`
	StatusBook      bool	      `json:"statusbook"`
	VaccineNumber	int			  `json:"vaccinenumber"`
	StatusNumber	bool		  `json:"statusnumber"`
}

func (*Booking) TableName() string {
	return "booking"
}
