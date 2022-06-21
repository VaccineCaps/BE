package model

type News struct {
	ID       		int    		`json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	Title    		string 		`json:"title" validate:"required"`
	Image    		string 		`json:"image" validate:"required"`
	Context 		string 		`json:"context" validate:"required"`
	HospitalId   	int    		`gorm:"not null" json:"hospital_id"`
	Hospital     	Hospitals   `gorm:"ForeignKey:HospitalId;references:ID;null" json:"-"`
}

func (*News) TableName() string {
	return "news"
}
