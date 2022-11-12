# pokemon-battle-royale
Farmacare Recruitment Exercise

## Prerequisites
1. Golang
2. MySQL

## How to run
1. Prepare your database (create new database), then adjust `.env` file for the database settings
2. run
   ```
   make migrate_up
   ```
3. run
   ```
   go mod download
   ```
4. run
   ```
   go run server.go
   ```

## Endpoints
1. [GET] `/pokemons/count`
2. [GET] `/matches`
3. [GET] `/pokemon-average-scores`
4. [POST] `/matches`
   example JSON body
   ```
    {
        "matchDetails": [
            {
                "pokemonId": 1,
                "score": 1
            },
            {
                "pokemonId": 2,
                "score": 2
            },
            {
                "pokemonId": 3,
                "score": 3
            },
            {
                "pokemonId": 4,
                "score": 4
            },
            {
                "pokemonId": 5,
                "score": 5
            }
        ]
    }
   ```
