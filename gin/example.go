package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"` // `json:"artist"` 는 struct 내용이 JSON 으로 serialized 될  때 field 명을 지정한다.
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) { // gin.Context 는 Gin 에서 가장 주요한 부분이다. 요청 세부사항을 전달하고, 유효성을 검사하며 JSON serialize 와 그 외의 작업도 수행한다.
	c.IndentedJSON(http.StatusOK, albums) // struct 를 JSON 으로 serialize 하고 response 에 보내기 위해 Context.IndentedJSON 를 호출했다.
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil { // Context.BindJSON 을 사용해서 request body 를 newAlbum 변수에 맵핑하였다.

		return
	}

	albums = append(albums, newAlbum)            // albums slice 에 album 을 추가했다.
	c.IndentedJSON(http.StatusCreated, newAlbum) // album 이 성공적으로 등록되었다는 의미로 상태코드 201(http.StatusCreated) 를 응답에 추가했다.
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found."})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.Run("localhost:8080")
}

// func main() {
// 	router := gin.Default()

// 	// This handler will match /user/john but will not match /user/ or /user
// 	router.GET("/user/:name", func(c *gin.Context) {
// 		name := c.Param("name")
// 		c.String(http.StatusOK, "Hello %s", name)
// 	})

// 	// However, this one will match /user/john/ and also /user/john/send
// 	// If no other routers match /user/john, it will redirect to /user/john/
// 	router.GET("/user/:name/*action", func(c *gin.Context) {
// 		name := c.Param("name")
// 		action := c.Param("action")
// 		message := name + " is " + action
// 		c.String(http.StatusOK, message)
// 	})

// 	// For each matched request Context will hold the route definition
// 	// router.POST("/user/:name/*action", func(c *gin.Context) {
// 	// 	c.FullPath() == "/user/:name/*action"
// 	// })

// 	router.Run(":8080")
// }
