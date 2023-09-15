package apiserver

import (
	"context"
	"log"

	_ "github.com/lib/pq"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt" // pw hash해서 저장

	"example/ent"
	"example/ent/user"
)

/* 다음 error 해결 과제 postgres 연결하기
2023/09/13 10:20:40 failed opening connection to postgres: sql: unknown driver "postgres" (forgotten import?)
exit status 1
2023/09/13 11:20:~~ 해결
echo로 framework 교체 결정됨.
*/

func ConnectDB() *ent.Client {
	// db pw 환경변수로 뺄 예정
	client, err := ent.Open("postgres", "host='root' port=5432 user=root dbname=postgres12 password=secret sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	// Run the auto migration tool.
	// if err := client.Schema.Create(context.Background()); err != nil {
	//     log.Fatalf("failed creating schema resources: %v", err)
	// }
	return client
}

func signupHandler(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		username := c.PostForm("username")
		password := c.PostForm("password")

		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		_, err = client.User.
			Create().
			SetUsername(username).
			SetPassword(string(hashedPassword)).
			Save(context.Background())

		if err != nil {
			c.JSON(500, gin.H{"status": "Internal server error"})
			return
		}

		// 확인용 라인
		log.Println("userinfo", username, hashedPassword)
		c.JSON(200, gin.H{"status": "Account created successfully"})
	}
}

func SigninHandler(client *ent.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		username := c.PostForm("username")
		password := c.PostForm("password")

		u, err := client.User.
			Query().
			Where(user.Username(username)).
			Only(context.Background())

		if err != nil || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
			c.JSON(401, gin.H{"status": "Unauthorized"})
			return
		}
	}
}

func apiserver() {
	client := ConnectDB()
	defer client.Close()

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true // 모든 도메인에서 오는 요청을 허용
	// config.AllowOrigins = []string{"http://localhost:3000"}

	// 라우터에서 cors 설정 사용하도록 함.
	router.Use(cors.New(config))

	router.POST("/signup", signupHandler(client))
	router.POST("/signin", SigninHandler(client))

	router.Run(":8080")
}
