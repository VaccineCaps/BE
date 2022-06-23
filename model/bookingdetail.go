package model

type BookingDetail struct {
	ID            int         `json:"id" gorm:"PrimaryKey;AUTO_INCREMENT;column:id"`
	UserId        int         `gorm:"not null" json:"user_id"`
	User          User        `gorm:"ForeignKey:UserId;references:ID;null" json:"-"`
	BookingId     int         `gorm:"not null" json:"booking_id"`
	Booking       Booking     `gorm:"ForeignKey:BookingId;references:ID;null" json:"-"`
	OtherPersonId int         `gorm:"not null" json:"otherperson_id"`
	OtherPerson   OtherPerson `gorm:"ForeignKey:OtherPersonId;references:ID;null" json:"-"`
}

func (*BookingDetail) TableName() string {
	return "bookingdetail"
}
