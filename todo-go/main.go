package main

import ("github.com/gin-gonic/gin"
	"net/http")
	


type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
} 

var todos = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Clean Class", Completed: false},
	{ID: "3", Item: "Clean roof", Completed: false},

}

func getTodos (context *gin.Context){
	context.JSON(http.StatusOK, todos)

}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:9090")
}