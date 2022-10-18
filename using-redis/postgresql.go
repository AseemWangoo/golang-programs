package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"

	"github.com/aseemwangoo/golang-programs/structs"
	"github.com/jackc/pgx/v4"
	_ "github.com/lib/pq"
)

const (
	hostProd = "127.0.0.1"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "dbname"
)

var (
	db *sql.DB
)

func ConnectDB() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", hostProd, port, user, password, dbname)

	var err error

	db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalf("Error while connecting db: %v", err)
	}
}

func PGXConnection() (*pgx.Conn, error) {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   hostProd,
		User:   url.UserPassword(user, password),
		Path:   dbname,
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")

	dsn.RawQuery = q.Encode()

	conn, err := pgx.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("pgx.Connect %w", err)
	}

	return conn, nil
}

func GetDBUsers() ([]structs.Users, error) {
	query := `
	select * from users`

	users := []structs.Users{}

	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Failed to run GetDBUsers : %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p structs.Users
		if err := rows.Scan(&p.ID, &p.CreatedTime, &p.Name, &p.UpdatedTime, &p.Age); err != nil {
			return nil, err
		}
		users = append(users, p)
	}

	return users, nil
}

func GetUserByID(id string) (structs.Users, error) {
	query := `
	select * from users
	WHERE id = $1;`

	rows := db.QueryRow(query, id)
	var p structs.Users

	err := rows.Scan(&p.ID, &p.CreatedTime, &p.Name, &p.UpdatedTime, &p.Age)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return p, err
	case nil:
		return p, err
	default:
		panic(err)
	}
}
