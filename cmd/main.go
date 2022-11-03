package main

import (
	"github.com/Vzhrkv/avito_internship/internal"
	"github.com/Vzhrkv/avito_internship/internal/handler"
	"github.com/Vzhrkv/avito_internship/internal/repository"
	"github.com/Vzhrkv/avito_internship/internal/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatal(err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	db, err := repository.NewPostgresDb(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbName"),
		SSLMode:  viper.GetString("db.sslMode"),
	})

	if err != nil {
		logrus.Fatal(err)
	}

	repo := repository.NewRepository(db)
	service_main := service.NewService(repo)
	handler_main := handler.NewHandler(service_main)

	srv := new(server.Server)
	port := viper.GetString("port")
	if err := srv.Run(port, handler_main.InitRoutes()); err != nil {
		logrus.Fatal(err)
	}
}
func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
