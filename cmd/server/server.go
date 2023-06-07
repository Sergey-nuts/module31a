package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgreSQL.
	postgrUser := os.Getenv("dbuser")
	postgrPwd := os.Getenv("dbpass")
	dbhost := os.Getenv("dbhost")
	fmt.Printf("conncting to postgresql... ")
	db2, err := postgres.New("postgres://" + postgrUser + ":" + postgrPwd + "@" + dbhost + "/GoNews")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success...\n")

	// Документная БД MongoDB.
	mongoDB := os.Getenv("mongodb")
	collection := "posts"
	fmt.Printf("connecting to MongoDB... ")
	db3, err := mongo.New("mongodb://"+dbhost+":27017/", mongoDB, collection)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("success...\n")

	_, _, _ = db, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db3

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	fmt.Printf("start server... listening on :8080\n")
	http.ListenAndServe(":8080", srv.api.Router())
}
