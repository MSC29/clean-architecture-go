package tests

import (
	"clean-architecture/clean-architecture-go/src/adapter/spi/db"
	"clean-architecture/clean-architecture-go/src/infrastructure"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	configuration := infrastructure.NewConfigurationMapper()
	dbConnection := db.NewDbConnection(configuration)

	dog1 := db.DogFact{
		ID:   1,
		Fact: "Forty-five percent of U.S. dogs sleep in their owner's bed",
	}
	dog2 := db.DogFact{
		ID:   2,
		Fact: "Seventy percent of people sign their dog's name on their holiday cards",
	}
	dog3 := db.DogFact{
		ID:   3,
		Fact: "Dogs have about 1,700 taste buds. We humans have between 2,000-10,000",
	}

	dbConnection.GetManager().Create(&dog1)
	dbConnection.GetManager().Create(&dog2)
	dbConnection.GetManager().Create(&dog3)

	os.Exit(m.Run())
}
