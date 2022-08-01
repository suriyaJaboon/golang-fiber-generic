package main

import (
	"context"
	"fg/handles"
	"fg/stores"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", " _"))
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	opt := stores.Option{
		ServiceName: viper.GetString("service"),
		URI:         viper.GetString("mongo.uri"),
		Addr:        viper.GetString("mongo.host"),
		Username:    viper.GetString("mongo.username"),
		Password:    viper.GetString("mongo.password"),
		Database:    viper.GetString("mongo.database"),
	}

	mgo, err := stores.MongoConnect(opt)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = mgo.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	log.Println("Connection Database-Server...")

	f := fiber.New(handles.Fc)

	f.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	g := f.Group("/api/v1")
	handles.NewProduct(g, mgo)

	if err = f.Listen(":3200"); err != nil {
		panic(err)
	}
}
