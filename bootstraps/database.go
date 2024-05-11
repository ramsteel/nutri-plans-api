package bootstraps

import (
	"fmt"
	"log"
	fpconst "nutri-plans-api/constants/filepath"
	msgconst "nutri-plans-api/constants/message"
	"nutri-plans-api/entities"
	loggerutil "nutri-plans-api/utils/logger"
	seedutil "nutri-plans-api/utils/seed"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase() *gorm.DB {
	var db *gorm.DB
	var err error

	var (
		DB_HOST      = os.Getenv("DB_HOST")
		DB_USER      = os.Getenv("DB_USER")
		DB_PASSWORD  = os.Getenv("DB_PASSWORD")
		DB_NAME      = os.Getenv("DB_NAME")
		DB_PORT      = os.Getenv("DB_PORT")
		DB_SSL       = os.Getenv("DB_SSL")
		DB_TZ        = os.Getenv("DB_TZ")
		DB_LOG_LEVEL = os.Getenv("DB_LOG_LEVEL")
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

	logLevel := loggerutil.GetDBLogLevel(DB_LOG_LEVEL)
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel:             logLevel,
			ParameterizedQueries: true,
			Colorful:             false,
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger,
	})
	if err != nil {
		log.Fatal(msgconst.MsgFailedConnectDB)
	}

	migrate(db)
	seed(db)

	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.RoleType{},              // role types
		&entities.Auth{},                  // auths
		&entities.Country{},               // countries
		&entities.User{},                  // users
		&entities.FoodType{},              // food types
		&entities.DrinkType{},             // drink types
		&entities.DietaryPreferenceType{}, // dietary preference types
		&entities.UserPreference{},        // user preference
		&entities.DietaryRestriction{},    // dietary restriction
		&entities.MealType{},              // meal types
		&entities.Meal{},                  // meal
		&entities.MealItem{},              // meal item
		&entities.Recommendation{},        // recommendation
		&entities.Admin{},                 // admin
	)

	if err != nil {
		log.Fatal(msgconst.MsgFailedMigrateDB)
	}
}

func seed(db *gorm.DB) {
	seedCountries(db)
	seedRoles(db)
	seedFoodTypes(db)
	seedDrinkTypes(db)
	seedDietaryPreferenceTypes(db)
	seedMealTypes(db)
}

func seedCountries(db *gorm.DB) {
	countries := seedutil.LoadCountryData(fpconst.CountryDataPath)
	if err := db.Save(countries).Error; err != nil {
		log.Fatal(msgconst.MsgSeedFailed)
	}
}

func seedRoles(db *gorm.DB) {
	roleTypes := seedutil.GetRoleTypes()
	if err := db.Save(roleTypes).Error; err != nil {
		log.Fatal(msgconst.MsgSeedFailed)
	}
}

func seedFoodTypes(db *gorm.DB) {
	foodTypes := seedutil.GetFoodTypes()

	for _, foodType := range *foodTypes {
		err := db.FirstOrCreate(&entities.FoodType{}, foodType).Error
		if err != nil {
			log.Fatal(msgconst.MsgSeedFailed)
		}
	}
}

func seedDrinkTypes(db *gorm.DB) {
	drinkTypes := seedutil.GetDrinkTypes()

	for _, drinkType := range *drinkTypes {
		err := db.FirstOrCreate(&entities.DrinkType{}, drinkType).Error
		if err != nil {
			log.Fatal(msgconst.MsgSeedFailed)
		}
	}
}

func seedDietaryPreferenceTypes(db *gorm.DB) {
	dietaryPreferencesTypes := seedutil.GetDietaryPreferenceTypes()

	for _, dietaryPreferenceType := range *dietaryPreferencesTypes {
		err := db.FirstOrCreate(&entities.DietaryPreferenceType{}, dietaryPreferenceType).Error
		if err != nil {
			log.Fatal(msgconst.MsgSeedFailed)
		}
	}
}

func seedMealTypes(db *gorm.DB) {
	mealTypes := seedutil.GetMealTypes()
	if err := db.Save(mealTypes).Error; err != nil {
		log.Fatal(msgconst.MsgSeedFailed)
	}
}
