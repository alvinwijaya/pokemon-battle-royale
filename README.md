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
4. [POST] `/matches`<br />
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
2. [DELETE] `/matches/match-details/:matchDetailId`

## Assumptions
1. Response https://pokeapi.co/ merupakan list pokemon yang dimiliki
2. Setiap pokemon dapat memiliki jumlah pertandingan yang berbeda, sehingga
   untuk perhitungan scorenya bisa dari average (total score semua pertandingan/jumlah pertandingan)

## Story
1. IWLT see pokemon stock number.
   endpoint get count pokemon (done)
2. IWLT record match result. For the score, the fastest to be out get 1 score, and so on until for
   the longest to stay get 5 score.
   endpoint store match result (done)
3. IWLT see all the matches within a time period.
   endpoint get all matches in pagination (additional param startDate and endDate) (done)
4. IWLT see overall pokemon score
   endpoint get all pokemon average score (done)
5. BONUS: IWLT annul certain pokemon from certain match
   endpoint to annul (done)