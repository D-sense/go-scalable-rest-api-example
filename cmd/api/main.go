package main

import (
	"context"
	"flag"
	"github.com/d-sense/go-scalable-rest-api-example/config"
	"github.com/d-sense/go-scalable-rest-api-example/database/postgres"
	"github.com/d-sense/go-scalable-rest-api-example/handlers/author"
	"github.com/d-sense/go-scalable-rest-api-example/handlers/book"
	"github.com/d-sense/go-scalable-rest-api-example/handlers/customer"
	"github.com/d-sense/go-scalable-rest-api-example/server/json"
	"github.com/gin-contrib/cors"
	_ "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("Starting the Scalable-Rest-API Service!")

	var envFile string
	flag.StringVar(&envFile, "env", "", "Environment Variable File")
	flag.Parse()

	if envFile == "" {
		log.Fatal("no ENV file provided")
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error: ", err)
	}

	initContext := &gin.Context{}

	configuration, err := config.LoadEnvFile()
	if err != nil {
		log.Println("error: ", err)
	}

	log.Println("Configuration Variables:", configuration)

	log.Println("Connecting to postgres")
	database := postgres.New(initContext, *configuration)
	defer database.Close()

	// Services
	authorService := postgres.NewAuthorService(initContext, database)
	customerService := postgres.NewCustomerService(initContext, database)
	bookService := postgres.NewBookService(initContext, database)

	// Handlers & Servers
	authorHandler := author.NewHandler(authorService, bookService)
	customerHandler := customer.NewHandler(customerService, bookService)
	bookHandler := book.NewHandler(bookService)

	// Server
	authorJsonServer := json.NewAuthorServer(authorHandler)
	customerJsonServer := json.NewCustomerServer(customerHandler)
	bookJsonServer := json.NewBookServer(bookHandler)

	// Configurations for CORS
	confg := cors.DefaultConfig()
	confg.AllowAllOrigins = true
	confg.AllowHeaders = []string{"Authorization", "Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host"}
	confg.ExposeHeaders = []string{"Content-Length"}
	confg.AllowMethods = []string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}
	confg.AllowCredentials = true

	router := gin.New()
	router.Use(cors.New(confg))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// setting up routes
	setUpRoutes(router, authorJsonServer, customerJsonServer, bookJsonServer)
	//SeedDefaultProducts(*database)

	srv := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	//goroutine this part to avoid blocking the code below it
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panicf("listen: %s\n", err)
		}
		log.Println("Scalable-Rest-API service is listening on address", srv.Addr)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Panicf("server shutdown: err=%v", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 3 seconds.")
	}
	log.Println("Server exiting")

}

func setUpRoutes(
	router *gin.Engine,
	authorJsonServer *json.AuthorJsonServer,
	customerJsonServer *json.CustomerJsonServer,
	bookJsonServer *json.BookJsonServer,
) {
	authorized := router.Group("/scalable-rest/api/v1")

	authorized.Use()
	{
		//Author routes
		authorized.POST("/author/register", authorJsonServer.SignUpAuthor())
		authorized.PUT("/author/login", authorJsonServer.LoginAuthor())
		authorized.GET("/author", authorJsonServer.GetAuthor())
		authorized.GET("/authors", authorJsonServer.GetAuthors())

		//Customer routes
		authorized.POST("/customer/register", customerJsonServer.SignUp())
		authorized.PUT("/customer/login", customerJsonServer.Login())
		authorized.GET("/customer", customerJsonServer.GetCustomer())
		authorized.GET("/customers", customerJsonServer.GetCustomers())

		//Book routes
		authorized.POST("/book/create", bookJsonServer.CreateBook())
		authorized.GET("/book", bookJsonServer.GetBook())
		authorized.GET("/books", bookJsonServer.GetBooks())
	}
}
