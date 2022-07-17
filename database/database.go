package database

import (
	config "BE/configs"
	"BE/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB(conf config.Config) *gorm.DB {

	conectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_HOST,
		conf.DB_PORT,
		conf.DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(conectionString))
	if err != nil {
		fmt.Println("error open conection : ", err)
	}

	DB.AutoMigrate(&model.Role{}, &model.User{}, &model.Provinces{}, &model.Cities{}, &model.Hospitals{},
		&model.News{}, &model.OtherPerson{}, &model.Vaccines{}, &model.VaccineHospitals{}, &model.Session{},
		&model.Certificate{}, &model.Booking{}, &model.BookingDetail{}, &model.VaccineTransactionsIn{},
		&model.Advertise{}, &model.VaccineTransactionsOut{})

	return DB

}
