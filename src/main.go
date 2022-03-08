package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NaokiYazawa/clean-architecture-go/domain/model"
	"github.com/NaokiYazawa/clean-architecture-go/infrastructure/persistence"
	"github.com/NaokiYazawa/clean-architecture-go/interface/controller"
	"github.com/NaokiYazawa/clean-architecture-go/usecase"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load env file: %v", err)
	}

	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	// db.AutoMigrate(&model.User{})

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return
	}

	// 依存関係の注入
	userRepository := persistence.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)

	// Echoのインスタンス作る
	e := echo.New()

	// 全てのリクエストで差し込みたいミドルウェア
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// ルーティング
	controller.InitRouting(e, userController)

	// サーバー起動
	e.Logger.Fatal(e.Start(":8080"))
}
