package main

import (
	"context"
	"database/sql"
	"go-uniswap/controllers"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/fx"
)

var ctx = context.Background()

func NewDB() (*sql.DB, error) {
	connStr := "postgres://root:root@go-uniswap-db:5432/go-uniswap-db"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// 測試連接
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func NewRedis() (*redis.Client, error) {
	// 建立 Redis 客戶端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 伺服器地址，這裡假設是 localhost 和 6379 埠
		Password: "",               // 如果有密碼，請填寫
		DB:       0,                // 選擇使用的 DB，默認是 0
	})

	// 驗證 Redis 連線是否正常
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

func NewGinServer() *gin.Engine {
	r := gin.Default()
	return r
}

func SetupServer(r *gin.Engine) {

	r.GET("/", controllers.Home)

	r.Run(":8080")
}

func main() {
	app := fx.New(
		fx.Provide(
			NewGinServer,
			NewDB,
		),
		fx.Invoke(SetupServer),
	)

	app.Start(ctx)
}
