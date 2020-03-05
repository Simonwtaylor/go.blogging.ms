package main

import (
	"authentication/subscriber"

	"authentication/handler"

	micro "github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"

	authentication "authentication/proto/authentication"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.authentication"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// err := godotenv.Load()

	// var r repository.Repository

	// if err != nil {
	// 	log.Fatalf("Error loading .env file, %v", err)
	// }

	// db, err := gorm.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD")))
	// if err != nil {
	// 	log.Fatalf("unable to open postgres db %v", err)
	// }

	// defer db.Close()

	// r := repository.ProvidePostRepository(db)

	// Register Handler
	authentication.RegisterAuthenticationHandler(service.Server(), new(handler.Authentication))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.authentication", service.Server(), new(subscriber.Authentication))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
