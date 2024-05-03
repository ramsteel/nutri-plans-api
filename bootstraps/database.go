package bootstraps

import (
	"fmt"
	"log"
	fpconst "nutri-plans-api/constants/filepath"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/entities"
	countryutil "nutri-plans-api/utils/country"
	roleutil "nutri-plans-api/utils/role"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	var (
		DB_HOST     = os.Getenv("DB_HOST")
		DB_USER     = os.Getenv("DB_USER")
		DB_PASSWORD = os.Getenv("DB_PASSWORD")
		DB_NAME     = os.Getenv("DB_NAME")
		DB_PORT     = os.Getenv("DB_PORT")
		DB_SSL      = os.Getenv("DB_SSL")
		DB_TZ       = os.Getenv("DB_TZ")
	)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DB_HOST,
		DB_USER,
		DB_PASSWORD,
		DB_NAME,
		DB_PORT,
		DB_SSL,
		DB_TZ,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(msgconst.MsgFailedConnectDB)
	}

	migrate(db)
	seed(db)

	return db
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&entities.RoleType{}, &entities.Auth{}) // auth
	db.AutoMigrate(&entities.User{}, &entities.Country{})  // user
}

func seed(db *gorm.DB) {
	seedCountries(db)
	seedRoles(db)
}

func seedCountries(db *gorm.DB) {
	countries := countryutil.LoadData(fpconst.CountryDataPath)
	if err := db.Save(countries).Error; err != nil {
		log.Fatal(msgconst.MsgSeedFailed)
	}
}

func seedRoles(db *gorm.DB) {
	roleTypes := roleutil.GetRoleTypes()
	if err := db.Save(roleTypes).Error; err != nil {
		log.Fatal(msgconst.MsgSeedFailed)
	}
}
