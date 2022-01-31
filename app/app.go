package app

import (
	"bankingV2/app/investmentHandler"
	"bankingV2/domain"
	"bankingV2/service"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	driverMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", DbUser, DbPasswd, DbAddress, DbPort, DbName)
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
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local", DbUser, DbPasswd, DbAddress, DbPort, DbName)
	db, err := gorm.Open(driverMysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("FATAL ERROR Gorm SQL DOES NOT WORK PROPERLY")
		panic(err)
	}
	//migrateAll(db)
	//migrations.MigrateInvestments(db)
	dbClient := getDbClient()

	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	authRepositoryDb := domain.NewAuthRepository(dbClient)
	jobRepositoryDb := domain.NewJobRepositoryDb(getDbClient())
	InvestmentRepositoryDb := domain.NewInvestmentRepositoryDb(getDbClient())
	jobRepositoryDbGorm := domain.NewJobRepositoryDbGorm(db)
	InvestmentRepositoryDbGorm := domain.NewInvestmentRepositoryDbGorm(db)

	jh := JobHandler{service.NewJobService(jobRepositoryDb)}
	jhgorm := JobHandlerGorm{service.NewJobServiceGorm(jobRepositoryDbGorm)}
	//ihgorm := InvestmentHandlerGorm{service.NewInvestmentServiceGorm(InvestmentRepositoryDbGorm)}
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	ach := AuthHandlers{service.NewLoginService(authRepositoryDb)}
	ihgorm := investmentHandler.InvestmentHandlerGorm{S: service.NewInvestmentServiceGorm(InvestmentRepositoryDbGorm)}
	ih := investmentHandler.InvestmentHandler{S: service.NewInvestmentService(InvestmentRepositoryDb)}
	//define routes

	router := mux.NewRouter()
	//wiring

	router.HandleFunc("/api/time", GetTime)

	customers := router.PathPrefix("/customers").Subrouter()
	customerById := customers.PathPrefix("/{customer_id:[0-9]+}").Subrouter()
	investments := router.PathPrefix("/investments").Subrouter()
	career := router.PathPrefix("/career").Subrouter()
	customers.HandleFunc("", ch.getAllCustomers).Methods(http.MethodGet).Name("GetAllCustomers")

	customerById.HandleFunc("/investments/create", ih.CustomerInvestmentCreate).Methods(http.MethodPost).Name("CreateCustomerInvestment")
	customerById.HandleFunc("/investments", ih.GetAllCustomerInvestments).Methods(http.MethodGet).Name("GetAllCustomerInvestments")
	customerById.HandleFunc("", ch.getCustomer).Methods(http.MethodGet).Name("GetCustomer")
	customerById.HandleFunc("/account", ah.NewAccount).Methods(http.MethodPost).Name("NewAccount")
	customerById.HandleFunc("/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost).Name("NewTransaction")

	investments.HandleFunc("/create", ihgorm.InvestmentsCreate).Methods(http.MethodPost).Name("InvestmentCreate")
	investments.HandleFunc("/risk-level/create", ihgorm.InvestmentRiskLevelCreate).Methods(http.MethodPost).Name("InvestmentRiskLevelCreate")
	investments.HandleFunc("/company/create", ihgorm.InvestmentCompanyCreate).Methods(http.MethodPost).Name("InvestmentCompanyCreate")
	investments.HandleFunc("/category/create", ihgorm.InvestmentCategoryCreate).Methods(http.MethodPost).Name("InvestmentCategoryCreate")

	career.HandleFunc("/career-at-seb", jh.GetAllJobs).Methods(http.MethodGet).Name("GetAllJobs")
	career.HandleFunc("/career-at-seb/{job_id:[0-9]+}", jh.GetById).Methods(http.MethodGet).Name("JobById")
	career.HandleFunc("/career-at-seb/new", jhgorm.NewJob).Methods(http.MethodPost).Name("NewJob")
	career.HandleFunc("/career-at-seb/update", jhgorm.UpdateJob).Methods(http.MethodPost).Name("UpdateJob")
	career.HandleFunc("/career-at-seb/delete", jhgorm.DeleteJob).Methods(http.MethodPost).Name("DeleteJob")

	am := AuthMiddleware{domain.NewAuthRepository(getDbClient())}
	router.Use(am.authorizationHandler())
	router.HandleFunc("/hello", Welcome).Methods(http.MethodGet).Name("Welcome")
	router.HandleFunc("/auth/login", ach.Login).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", ServerAddress, ServerPort), router))
	//if err != nil {
	//	log.Println("something went wrong", err)
	//	return
	//}

}

//func migrateAll(db *gorm.DB) {
//	err := db.AutoMigrate(&domain.JobGorm{})
//	if err != nil {
//		fmt.Println("error while migrating")
//	}
//}
//func executeSeeder(*sqlx.DB) {
//	seeds.Execute(getDbClient())
//}
