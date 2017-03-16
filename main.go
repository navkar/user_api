package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

// Users (Id, Firstname, Lastname)
type Users struct {
	Id        int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Firstname string `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string `gorm:"not null" form:"lastname" json:"lastname"`
}

func InitDb() *gorm.DB {
	// var host string
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DBNAME")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	connString := "host=" + host + " user=" + user + " dbname=" + dbName + " sslmode=disable password=" + password

	fmt.Println("connection string: " + connString)
	// Openning file
	db, err := gorm.Open("postgres", connString)
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}
	// Creating the table
	if !db.HasTable(&Users{}) {
		db.CreateTable(&Users{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Users{})
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {

	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/users", PostUser)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
		v1.PUT("/users/:id", UpdateUser)
		v1.DELETE("/users/:id", DeleteUser)
	}

	r.Run(":8080")
}

func PostUser(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var user Users
	c.Bind(&user)

	if user.Firstname != "" && user.Lastname != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display error
		c.JSON(201, gin.H{"success": user})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/users
}

func GetUsers(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var users []Users
	// SELECT * FROM users
	db.Find(&users)

	// Display JSON result
	c.JSON(200, users)

	// curl -i http://localhost:8080/api/v1/users
}

func GetUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Id != 0 {
		// Display JSON result
		c.JSON(200, user)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i http://localhost:8080/api/v1/users/1
}

func UpdateUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Firstname != "" && user.Lastname != "" {

		if user.Id != 0 {
			var newUser Users
			c.Bind(&newUser)

			result := Users{
				Id:        user.Id,
				Firstname: newUser.Firstname,
				Lastname:  newUser.Lastname,
			}

			// UPDATE users SET firstname='newUser.Firstname', lastname='newUser.Lastname' WHERE id = user.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "User not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/users/1
}

func DeleteUser(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id user
	id := c.Params.ByName("id")
	var user Users
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, id)

	if user.Id != 0 {
		// DELETE FROM users WHERE id = user.Id
		db.Delete(&user)
		// Display JSON result
		c.JSON(200, gin.H{"success": "User #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "User not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/users/1
}

func OptionsUser(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST,PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
