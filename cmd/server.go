package main


import (
	"os"
	"fmt"
	"net/http"
	"net"
	"log"
	"context"
	"time"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	if err := godotenv.Load();err != nil {
		log.Panic("Error loading .env file",err)
	}
	port := os.Getenv("SERVER_PORT")

	listner,err := net.Listen("tcp",fmt.Sprintf(":%s",port))
	if err != nil {
		log.Panic("Error listening on port",err)
	}


}