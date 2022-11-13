package main

import (
	"github.com/labstack/echo"

	"github.com/alvinwijaya/pokemon-battle-royale/config"
	"github.com/alvinwijaya/pokemon-battle-royale/db"

	_matchController "github.com/alvinwijaya/pokemon-battle-royale/api/match/controller"
	_matchRepository "github.com/alvinwijaya/pokemon-battle-royale/api/match/repository"
	_matchService "github.com/alvinwijaya/pokemon-battle-royale/api/match/service"

	_matchDetailController "github.com/alvinwijaya/pokemon-battle-royale/api/match_detail/controller"
	_matchDetailRepository "github.com/alvinwijaya/pokemon-battle-royale/api/match_detail/repository"
	_matchDetailService "github.com/alvinwijaya/pokemon-battle-royale/api/match_detail/service"

	_pokemonAverageScoreController "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score/controller"
	_pokemonAverageScoreRepository "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score/repository"
	_pokemonAverageScoreService "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon_average_score/service"

	_pokemonController "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon/controller"
	_pokemonService "github.com/alvinwijaya/pokemon-battle-royale/api/pokemon/service"

	_pokeapiReq "github.com/alvinwijaya/pokemon-battle-royale/request/pokeapi"
)

func main() {
	conf := config.GetConfig()

	dbManager := db.NewDatabaseManager()
	dsn := config.GetConfig().DbUser + ":" + config.GetConfig().DbPassword + "@tcp(" + config.GetConfig().DbHost + ":" + config.GetConfig().DbPort + ")/" + config.GetConfig().DbName + "?parseTime=true"
	dbManager.Initialize(
		dsn,
		config.GetConfig().DbConnection,
	)

	pokeapiReq := _pokeapiReq.NewRequest()

	matchRepository := _matchRepository.NewRepository()
	matchDetailRepository := _matchDetailRepository.NewRepository()
	pokemonAverageScoreRepository := _pokemonAverageScoreRepository.NewRepository()

	matchService := _matchService.NewService(
		dbManager,
		matchRepository,
		matchDetailRepository,
		pokemonAverageScoreRepository,
		pokeapiReq,
	)

	matchDetailService := _matchDetailService.NewService(
		dbManager,
		matchDetailRepository,
		pokemonAverageScoreRepository,
	)

	pokemonService := _pokemonService.NewService(
		pokeapiReq,
	)

	pokemonAverageScoreService := _pokemonAverageScoreService.NewService(
		dbManager,
		pokemonAverageScoreRepository,
	)

	matchController := _matchController.NewController(matchService)
	matchDetailController := _matchDetailController.NewController(matchDetailService)
	pokemonController := _pokemonController.NewController(pokemonService)
	pokemonAverageScoreController := _pokemonAverageScoreController.NewController(pokemonAverageScoreService)

	e := echo.New()

	// Init routes
	e.GET("/pokemons/count", pokemonController.GetCount)
	e.GET("/matches", matchController.GetAllInPagination)
	e.POST("/matches", matchController.Store)
	e.GET("/pokemon-average-scores", pokemonAverageScoreController.GetAllInPagination)
	e.DELETE("/matches/match-details/:matchDetailId", matchDetailController.Delete)

	e.Logger.Fatal(e.Start(":" + conf.AppPort))
}
