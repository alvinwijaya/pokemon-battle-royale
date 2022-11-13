#!/bin/bash

# Database
mockgen --source=db/database.go -package=mock_database --destination=mocks/db/database.go

# Match
mockgen --source=api/match/repository.go -package=mock_match --destination=mocks/api/match/repository.go
mockgen --source=api/match/service.go -package=mock_match --destination=mocks/api/match/service.go

# Match Detail
mockgen --source=api/match_detail/repository.go -package=mock_match_detail --destination=mocks/api/match_detail/repository.go
mockgen --source=api/match_detail/service.go -package=mock_match_detail --destination=mocks/api/match_detail/service.go

# Match
mockgen --source=api/match/repository.go -package=mock_match --destination=mocks/api/match/repository.go
mockgen --source=api/match/service.go -package=mock_match --destination=mocks/api/match/service.go

# Pokemon
mockgen --source=api/pokemon/service.go -package=mock_pokemon --destination=mocks/api/pokemon/service.go

# Pokemon Average Score
mockgen --source=api/pokemon_average_score/repository.go -package=mock_pokemon_average_score --destination=mocks/api/pokemon_average_score/repository.go
mockgen --source=api/pokemon_average_score/service.go -package=mock_pokemon_average_score --destination=mocks/api/pokemon_average_score/service.go

# Request
mockgen --source=request/pokeapi/interface.go -package=mock_request_pokeapi --destination=mocks/request/pokeapi/interface.go
