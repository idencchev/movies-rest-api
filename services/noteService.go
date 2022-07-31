package services

import (
	"movies-rest-api/middlewares"
)

func init() {
	middlewares.ConnectionWithMongoDB("notes")
}