package main

import (
	"database/sql"
	"fmt"
	// "log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age int32
}

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/simple_site")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func main() {
	fmt.Println("work with MySQL")

	// db, err := ConnectToDB()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer db.Close()
	
	// _, err = db.Exec("INSERT INTO `users` (`name`, `age`) VALUES ('Alex', 21)")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// rows, err := db.Query("SELECT `name`, `age` FROM `users`")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()

	// var user User
	// for rows.Next() {
    //     err := rows.Scan(&user.Name, &user.Age)
    //     if err != nil {
    //         log.Fatal(err)
	// 	}
	// 	fmt.Printf("User with Name = %s and Age = %d\n", user.Name, user.Age)
    // }
    // if err = rows.Err(); err != nil {
    //     log.Fatal(err)
    // }
}
