package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/phper95/pkg/cache"
	"github.com/phper95/pkg/db"
	"github.com/phper95/pkg/es"
	"github.com/phper95/pkg/httpclient"
	"github.com/phper95/pkg/logger"
	"github.com/phper95/pkg/mq"
	"github.com/phper95/pkg/nosql"
	"github.com/phper95/pkg/prome"
	"github.com/phper95/pkg/shutdown"
	"github.com/phper95/pkg/trace"
	"log"
	"net/http"
	"safe-shutdown/metric"
	"time"
)

func init() {
	InitLog()
	initMysqlClient()
	initRedisClient()
	initMongoClient()
	initESClient()
	initProme()

	err := mq.InitSyncKafkaProducer(mq.DefaultKafkaSyncProducer,
		[]string{"127.0.0.1:9092"}, nil)
	if err != nil {
		log.Fatal("InitSyncKafkaProducer err", err, "client", mq.DefaultKafkaSyncProducer)
	}
}
func InitLog() {
	//日志显示行号和文件名
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
func initMysqlClient() {
	err := db.InitMysqlClient(db.DefaultClient, "user", "pwd", "127.0.0.1:3306", "test")
	if err != nil {
		log.Fatal("mysql init error", err)
	}
}
func initRedisClient() {
	opt := redis.Options{
		Addr:         "127.0.0.1:6379",
		DB:           0,
		MaxRetries:   3,
		PoolSize:     20,
		MinIdleConns: 100,
	}
	redisTrace := trace.Cache{
		Name:                  "redis",
		SlowLoggerMillisecond: 500,
		Logger:                logger.GetLogger(),
		AlwaysTrace:           true,
	}
	err := cache.InitRedis(cache.DefaultRedisClient, &opt, &redisTrace)
	if err != nil {
		log.Fatal("redis init error", err)
	}
}

func initESClient() {
	err := es.InitClientWithOptions(es.DefaultClient, []string{"127.0.0.1:9200"},
		"User",
		"Password",
		es.WithScheme("https"))
	if err != nil {
		log.Fatal("InitClientWithOptions error", err, es.DefaultClient)
	}
}

func initMongoClient() {
	err := nosql.InitMongoClient(nosql.DefaultMongoClient, "user",
		"pwd", []string{"[127.0.0.1:27017"}, 200)
	if err != nil {
		log.Fatal("InitMongoClient error", err, nosql.DefaultMongoClient)
	}

}

func initProme() {
	prome.InitPromethues("172.0.0.1:9091", time.Second*60, metric.AppName, httpclient.DefaultClient, metric.TestCostTime)
}
func main() {
	router := gin.Default()
	listenAddr := fmt.Sprintf(":%d", 8888)
	server := &http.Server{
		Addr:           listenAddr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("http server start error", err)
		}
	}()

	//优雅关闭（封装版）
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Println("http server shutdown err", err)
			}
		},
		//关闭kafka producer
		func() {
			if err := mq.GetKafkaSyncProducer(mq.DefaultKafkaSyncProducer).Close(); err != nil {
				log.Println("kafka close error", err, "client", mq.DefaultKafkaSyncProducer)
			}
		},
		func() {
			es.CloseAll()
		},
		func() {
			//关闭mysql
			if err := db.CloseMysqlClient(db.DefaultClient); err != nil {
				log.Println("mysql shutdown err", err, db.DefaultClient)
			}
		},

		func() {
			err := cache.GetRedisClient(cache.DefaultRedisClient).Close()
			if err != nil {
				log.Println("redis close error", err, cache.DefaultRedisClient)
			}
		},
		func() {
			if nosql.GetMongoClient(nosql.DefaultMongoClient) != nil {
				nosql.GetMongoClient(nosql.DefaultMongoClient).Close()
			}
		},
	)
}
