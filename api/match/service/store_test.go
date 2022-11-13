package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/alvinwijaya/pokemon-battle-royale/config"
	"github.com/alvinwijaya/pokemon-battle-royale/custom_error"
	model "github.com/alvinwijaya/pokemon-battle-royale/model/match"
	"github.com/alvinwijaya/pokemon-battle-royale/model/match_detail"
	pas_model "github.com/alvinwijaya/pokemon-battle-royale/model/pokemon_average_score"
	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

func (s *MatchSuite) TestStore() {
	var (
		db      = gorm.DB{}
		payload = model.StoreMatchDTO{
			MatchDetails: []match_detail.StoreMatchDetailDTO{
				{
					PokemonID: 1,
					Score:     1,
				},
				{
					PokemonID: 2,
					Score:     2,
				},
				{
					PokemonID: 3,
					Score:     3,
				},
				{
					PokemonID: 4,
					Score:     4,
				},
				{
					PokemonID: 5,
					Score:     5,
				},
			},
		}
		pokemonAverageScores = []pas_model.PokemonAverageScore{
			{
				ID:           1,
				PokemonID:    1,
				AverageScore: 1,
			},
		}
		match = model.Match{
			ID: 1,
		}
		forceErr = errors.New("Force error")
	)

	config.SetConfig(config.Config{
		MatchParticipantCount: 5,
	})

	s.dbManager.EXPECT().GetDB().Return(&db).AnyTimes()
	s.dbManager.EXPECT().StartTransaction().Return(&db).AnyTimes()
	tx := s.dbManager.StartTransaction()
	s.dbManager.EXPECT().CommitTransaction(tx).Return(&db).AnyTimes()
	s.dbManager.EXPECT().RollbackTransaction(tx).Return(&db).AnyTimes()

	s.Run("Case Success", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(&pokemonAverageScores, nil)
		s.matchRepository.EXPECT().Store(&db, gomock.Any()).Return(&match, nil)
		s.matchDetailRepository.EXPECT().BulkStore(&db, gomock.Any()).Return(nil)
		s.pokemonAverageScoreRepository.EXPECT().UpdateByID(&db, gomock.Any(), gomock.Any()).Return(nil, nil)
		s.pokemonAverageScoreRepository.EXPECT().Store(&db, gomock.Any()).Return(nil, nil).Times(5)

		actualResult, gotErr := s.matchService.Store(&payload)

		s.Nil(gotErr)
		s.Equal(actualResult, "Successfully created.")
	})

	s.Run("Case Failed, match detail count is not valid", func() {
		payloadCaseFailed := model.StoreMatchDTO{
			MatchDetails: []match_detail.StoreMatchDetailDTO{
				{
					PokemonID: 1,
					Score:     1,
				},
			},
		}

		_, gotErr := s.matchService.Store(&payloadCaseFailed)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusBadRequest,
			Message: fmt.Sprintf("match detail count != %v", config.GetConfig().MatchParticipantCount),
		})
	})

	s.Run("Case Failed, match detail contains duplicate pokemonId", func() {
		payloadCaseFailed := model.StoreMatchDTO{
			MatchDetails: []match_detail.StoreMatchDetailDTO{
				{
					PokemonID: 1,
					Score:     1,
				},
				{
					PokemonID: 1,
					Score:     1,
				},
				{
					PokemonID: 1,
					Score:     1,
				},
				{
					PokemonID: 1,
					Score:     1,
				},
				{
					PokemonID: 1,
					Score:     1,
				},
			},
		}

		s.pokeapiReq.EXPECT().GetByID(gomock.Any())

		_, gotErr := s.matchService.Store(&payloadCaseFailed)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusConflict,
			Message: fmt.Sprintf("error duplicate pokemonId %v", payloadCaseFailed.MatchDetails[0].PokemonID),
		})
	})

	s.Run("Case Failed, error when call pokeapiReq.GetByID", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Return(nil, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})

	s.Run("Case Failed, error when call pokemonAverageScoreRepository.GetByPokemonIDs", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(nil, forceErr)

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})

	s.Run("Case Failed, error when call matchRepository.Store", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(&pokemonAverageScores, nil)
		s.matchRepository.EXPECT().Store(&db, gomock.Any()).Return(nil, forceErr)

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})

	s.Run("Case Failed, error when call matchDetailRepository.BulkStore", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(&pokemonAverageScores, nil)
		s.matchRepository.EXPECT().Store(&db, gomock.Any()).Return(&match, nil)
		s.matchDetailRepository.EXPECT().BulkStore(&db, gomock.Any()).Return(forceErr)

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})

	s.Run("Case Failed, error when call pokemonAverageScoreRepository.UpdateByID", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(&pokemonAverageScores, nil)
		s.matchRepository.EXPECT().Store(&db, gomock.Any()).Return(&match, nil)
		s.matchDetailRepository.EXPECT().BulkStore(&db, gomock.Any()).Return(nil)
		s.pokemonAverageScoreRepository.EXPECT().UpdateByID(&db, gomock.Any(), gomock.Any()).Return(nil, forceErr)

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})

	s.Run("Case Failed, error when call pokemonAverageScoreRepository.Store", func() {
		s.pokeapiReq.EXPECT().GetByID(gomock.Any()).Times(5)
		s.pokemonAverageScoreRepository.EXPECT().GetByPokemonIDs(&db, gomock.Any()).Return(&pokemonAverageScores, nil)
		s.matchRepository.EXPECT().Store(&db, gomock.Any()).Return(&match, nil)
		s.matchDetailRepository.EXPECT().BulkStore(&db, gomock.Any()).Return(nil)
		s.pokemonAverageScoreRepository.EXPECT().UpdateByID(&db, gomock.Any(), gomock.Any()).Return(nil, nil)
		s.pokemonAverageScoreRepository.EXPECT().Store(&db, gomock.Any()).Return(nil, forceErr)

		_, gotErr := s.matchService.Store(&payload)

		s.NotNil(gotErr)
		s.Equal(gotErr, &custom_error.ServiceError{
			Code:    http.StatusInternalServerError,
			Message: forceErr.Error(),
		})
	})
}
