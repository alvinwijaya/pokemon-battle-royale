package service

import (
	"testing"

	"github.com/alvinwijaya/pokemon-battle-royale/api/match"
	mock_match "github.com/alvinwijaya/pokemon-battle-royale/mocks/api/match"
	mock_match_detail "github.com/alvinwijaya/pokemon-battle-royale/mocks/api/match_detail"
	mock_pokemon_average_score "github.com/alvinwijaya/pokemon-battle-royale/mocks/api/pokemon_average_score"
	mock_database "github.com/alvinwijaya/pokemon-battle-royale/mocks/db"
	mock_request_pokeapi "github.com/alvinwijaya/pokemon-battle-royale/mocks/request/pokeapi"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type MatchSuite struct {
	suite.Suite
	*require.Assertions
	ctrl                          *gomock.Controller
	dbManager                     *mock_database.MockDatabaseManager
	matchRepository               *mock_match.MockRepository
	matchDetailRepository         *mock_match_detail.MockRepository
	pokemonAverageScoreRepository *mock_pokemon_average_score.MockRepository
	pokeapiReq                    *mock_request_pokeapi.MockInterface

	matchService match.Service
}

func TestMatchSuite(t *testing.T) {
	suite.Run(t, new(MatchSuite))
}

func (s *MatchSuite) SetupTest() {
	s.Assertions = require.New(s.T())
	s.ctrl = gomock.NewController(s.T())

	s.dbManager = mock_database.NewMockDatabaseManager(s.ctrl)
	s.matchRepository = mock_match.NewMockRepository(s.ctrl)
	s.matchDetailRepository = mock_match_detail.NewMockRepository(s.ctrl)
	s.pokemonAverageScoreRepository = mock_pokemon_average_score.NewMockRepository(s.ctrl)
	s.pokeapiReq = mock_request_pokeapi.NewMockInterface(s.ctrl)

	s.matchService = NewService(s.dbManager, s.matchRepository, s.matchDetailRepository, s.pokemonAverageScoreRepository, s.pokeapiReq)
}

func (s *MatchSuite) TearDownTest() {
	s.ctrl.Finish()
}
