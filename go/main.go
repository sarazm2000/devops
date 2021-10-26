
package main

import( 
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
 	"fmt"
	 "crypto/sha256"
	 "encoding/base64"
	 "github.com/go-redis/redis"
	//  "context"
)
var router *gin.Engine

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next();
		
    }
}



func main(){
	router = gin.Default()
	router.Use(cors.Default())
	router.Use(CORSMiddleware())

	client :=redis.NewClient(&redis.Options{
		Addr : "redis:6379" , 
		Password: "",
		DB : 0,
	
	})

	Pong , err := client.Ping().Result()
	fmt.Println(Pong , err)
	router.POST("/", func(c *gin.Context){
		input := c.PostForm("input")
		// hashed := sha256.Sum([]byte(input))
		if (len([]rune(input)) < 8){
			return
		}
		hasher := sha256.New()
		hasher.Write([]byte (input))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		fmt.Println(sha)
		// fmt.Println(input)
		// fmt.Println("djj")
		// var ctx = context.Background()
		err := client.Set(sha, input , 0).Err()
		if err != nil {
			panic(err)
		}
		
		c.JSON(200, gin.H{
			"message" : sha,
		})
	})
	router.GET("/", func(c *gin.Context) {
		fmt.Println("lllll")
		q := c.Request.URL.Query()
		i := q["input"][0]
		fmt.Println(i)
		
		fmt.Println(client.Exists(i))
		val, err := client.Get(i).Result()
		if err == redis.Nil {
			fmt.Println("nadarim")
			c.JSON(200, gin.H{
				"message" : "nadarim",
			})
			return
		} else if err != nil {
			panic(err)
			fmt.Println("input does not exist")
			
		} else {
			fmt.Println("input", val)
		}
		c.JSON(200, gin.H{
			"message" : val,
		})
	})
	router.Run()
	
}
