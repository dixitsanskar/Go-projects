package main

import ("github.com/gin-gonic/gin"
	"net/http"
	"errors"
)
	


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

func addTodos(context *gin.Context){
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.JSON(http.StatusCreated, todos)
}

func goToId(id string) (*todo, error) {
	
	for i, todoId:= range todos{
		if todoId.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}
func getTodo(context *gin.Context){
	id := context.Param("id")
	if todo, err := goToId(id); err == nil {
		context.JSON(http.StatusOK, todo)
		return
	
	}
	 context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
}

func changeTodo(context *gin.Context){
	id := context.Param("id")
	if tod, err := goToId(id); err == nil {
		tod.Completed = !tod.Completed
		context.JSON(http.StatusOK, tod)
		return
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})

}

func main(){
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id",getTodo )
	router.PATCH("/todos/:id",changeTodo )
	router.POST("/todos", addTodos)
	router.Run("localhost:9090")
}