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
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
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

	DB.AutoMigrate(&model.Role{})
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Provinces{})
	DB.AutoMigrate(&model.Cities{})
	DB.AutoMigrate(&model.Hospitals{})
	DB.AutoMigrate(&model.News{})
	DB.AutoMigrate(&model.OtherPerson{})
	DB.AutoMigrate(&model.Vaccines{})
	DB.AutoMigrate(&model.VaccineHospitals{})
	DB.AutoMigrate(&model.Session{})
	DB.AutoMigrate(&model.VaccineStatus{})
	DB.AutoMigrate(&model.Booking{})
	DB.AutoMigrate(&model.BookingDetail{})


	return DB

}
