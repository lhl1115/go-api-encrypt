package main

import "github.com/lhl1115/go-api-encrypt/gin"

func main() {

	ginServer := gin.NewRouter()
	err := ginServer.Run(":8082")
	if err != nil {
		panic(err)
		return
	}

}
