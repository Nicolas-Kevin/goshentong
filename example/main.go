package main

import (
	"database/sql"
	"log"

	_ "githu.com/wb253/goshentong"
)

func main() {
	db, err := sql.Open("shentong", "user/pass@127.0.0.1:2003/OSRDB")
	if err != nil {
		log.Fatal(err)
	}

	var id string
	err = db.QueryRow("select id from system_task where id= :1", 1).Scan(&id)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(id)

}
