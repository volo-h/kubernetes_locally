package main

import (
	"github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
	// "github.com/jinzhu/gorm"
  // _ "github.com/jinzhu/gorm/dialects/mysql"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"

  "fmt"
  "log"
  "os"
	"net/http"
)

func main() {
  // var err error
  _, err := os.Stat(".env")
  if err == nil {
    err := godotenv.Load()

    if err != nil {
      log.Fatalf("Error loading .env file")
    }
  }
  // if err != nil {
  //   log.Fatalf("Error getting env, %v", err)
  // } else {
  //   fmt.Println("We are getting values")
  // }

	//Migrate schema
	db := Database()
	db.AutoMigrate(&Todo{})

	router := gin.Default()
       router.Use(CORSMiddleware())

	v1 := router.Group("/api/v1/todos")
	 {
	 	v1.POST("/", CreateTodo)
	 	v1.GET("/", FetchAllTodo)
	 	v1.GET("/:id", FetchSingleTodo)
	 	v1.PUT("/:id", UpdateTodo)
	 	v1.DELETE("/:id", DeleteTodo)
	 }

  apiPort := fmt.Sprintf(":%s", os.Getenv("API_SERVER_PORT"))

  router.Run(apiPort)
}

type Todo struct {
	gorm.Model
	Title		string 	`json:"title"`
	Completed 	int 	`json:"completed"`
}

type TransformedTodo struct {
	ID 			uint	`json:"id"`
	Title		string 	`json:"title"`
	Completed 	bool 	`json:"completed"`
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
}

func Database() *gorm.DB {
  // "root:password@tcp(127.0.0.1:3306)/demo"
  DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

  db, err := gorm.Open(mysql.Open(DBURL), &gorm.Config{}) // local ony


	if err != nil {
		panic("failed to connect database")
	}
  	return db
}

func CreateTodo(c *gin.Context) {
       var todo Todo
       c.BindJSON(&todo)
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// todo := Todo{Title: c.PostForm("title"), Completed: completed};
	db := Database()
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"status" : http.StatusCreated, "message" : "Todo item created successfully!", "resourceId" : todo.ID})
}

func FetchAllTodo(c *gin.Context) {
       var todos []Todo
       var _todos []TransformedTodo

       db := Database()
       db.Find(&todos)

       if (len(todos) <= 0) {
              c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
              return
       }

       //transforms the todos for building a good response
       for _, item := range todos {
              completed := false
              if (item.Completed == 1) {
                     completed = true
              } else {
                     completed = false
              }
              _todos = append(_todos, TransformedTodo{ID: item.ID, Title:item.Title, Completed: completed})
       }
       c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _todos})
}

func FetchSingleTodo(c *gin.Context) {
       var todo Todo
       todoId := c.Param("id")

       db := Database()
       db.First(&todo, todoId)

       if (todo.ID == 0) {
              c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
              return
       }

       completed := false
       if (todo.Completed == 1) {
              completed = true
       } else {
              completed = false
       }

       _todo := TransformedTodo{ID: todo.ID, Title:todo.Title, Completed: completed}
       c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "data" : _todo})
}

func UpdateTodo(c *gin.Context) {
       var todo Todo
       var updatedTodo Todo

       todoId := c.Param("id")
       db := Database()
       db.First(&todo, todoId)

       if (todo.ID == 0) {
              c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
              return
       }

       c.BindJSON(&updatedTodo)
       db.Model(&todo).Update("title", updatedTodo.Title)
       db.Model(&todo).Update("completed", updatedTodo.Completed)
       c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo updated successfully!"})
}

func DeleteTodo(c *gin.Context) {
       var todo Todo
       todoId := c.Param("id")
       db := Database()
       db.First(&todo, todoId)

       if (todo.ID == 0) {
              c.JSON(http.StatusNotFound, gin.H{"status" : http.StatusNotFound, "message" : "No todo found!"})
              return
       }

       db.Delete(&todo)
       c.JSON(http.StatusOK, gin.H{"status" : http.StatusOK, "message" : "Todo deleted successfully!"})
}
