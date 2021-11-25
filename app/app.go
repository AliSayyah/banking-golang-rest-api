package app

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/logger"
	"github.com/AliSayyah/banking/service"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {
	logger.Info("Starting server on port 8000")

	router := mux.NewRouter()

	// database client
	dbClient := getDBClient()
	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}
	ah := AccountHandler{service: service.NewAccountService(domain.NewAccountRepositoryDB(dbClient))}
	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)

	// Start the server
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func getDBClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", GetDSN())
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}

func GetDSN() string {
	// load .env file and get environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// get database connection string
	mysqlUser := os.Getenv("MYSQL_USER")
	MysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	return mysqlUser + ":" + MysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB
}
