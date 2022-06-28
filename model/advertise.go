package model

type Advertise struct {
	ID       		int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Image    		string 		`json:"image" validate:"required"`
	HospitalId   	int    		`gorm:"not null" json:"hospital_id"`
	Hospital     	Hospitals   `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
}

func (*Advertise) TableName() string {
	return "advertise"
}
