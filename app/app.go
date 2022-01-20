package app

import (
	"bankingV2/domain"
	"bankingV2/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
)

//func sanityCheck() {
//	if os.Getenv("SERVER_ADDRESS") == "" ||
//		os.Getenv("SERVER_PORT") == "" {
//		log.Fatal("Environment var is missing")
//	}
//}
var DbUser string
var DbPasswd string
var DbAddress string
var DbPort string
var DbName string
var ServerAddress string
var ServerPort string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	DbUser = os.Getenv("DB_USER")
	DbPasswd = os.Getenv("DB_PASSWD")
	DbAddress = os.Getenv("DB_ADDR")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	ServerAddress = os.Getenv("SERVER_ADDRESS")
	ServerPort = os.Getenv("SERVER_PORT")

}
func Start() {
	//sanityCheck()
	timezoneTask()
}
func getDbClient() *sqlx.DB {

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUser, DbPasswd, DbAddress, DbPort, DbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		log.Println("FATAL ERROR SQL DOES NOT WORK PROPERLY")
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
func timezoneTask() {
	router := mux.NewRouter()
	//wiring
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	authRepositoryDb := domain.NewAuthRepository(dbClient)
	jobRepositoryDb := domain.NewJobRepositoryDb(getDbClient())
	jh := JobHandler{service.NewJobService(jobRepositoryDb)}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	ach := AuthHandlers{service.NewLoginService(authRepositoryDb)}
	//define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost).Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost).Name("NewTransaction")
	router.HandleFunc("/career/career-at-seb", jh.GetAllJobs).Methods(http.MethodGet).Name("GetAllJobs")
	router.HandleFunc("/career/career-at-seb/{job_id:[0-9]+}", jh.GetById).Methods(http.MethodGet).Name("JobById")
	router.HandleFunc("/career/career-at-seb/new", jh.NewJob).Methods(http.MethodPost).Name("NewJob")
	router.HandleFunc("/api/time", GetTime)

	am := AuthMiddleware{domain.NewAuthRepository(getDbClient())}
	router.Use(am.authorizationHandler())
	router.HandleFunc("/auth/login", ach.Login).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", ServerAddress, ServerPort), router))
	//if err != nil {
	//	log.Println("something went wrong", err)
	//	return
	//}

}
