package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "chessvger"
	password = "chessvger"
	dbname   = "chessvger"
)

type Player struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ListPlayers() []Player {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	baseQuery := "SELECT id, name FROM common_player limit 100"

	rows, err := db.Query(baseQuery)
	if err != nil {
		fmt.Print(fmt.Errorf("error scanning row: %v", err))
		return nil
	}
	defer rows.Close()

	// Parse the results into a slice of Product structs
	var players []Player
	for rows.Next() {
		var player Player
		if err := rows.Scan(&player.ID, &player.Name); err != nil {
			fmt.Print(fmt.Errorf("error scanning row: %v", err))
			return nil
		}
		players = append(players, player)
	}

	if err = rows.Err(); err != nil {
		fmt.Print(fmt.Errorf("error scanning row: %v", err))
		return nil
	}

	return players
}
