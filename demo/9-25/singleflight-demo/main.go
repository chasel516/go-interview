package main

import (
	"golang.org/x/sync/singleflight"
	"log"
	"sync"
	"time"
)

var singleFlightGetArticle singleflight.Group
var list []int

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	key := "list_id"
	var wg sync.WaitGroup
	//模拟并发请求
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()
			data, err := getArticleData(requestID, key)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(data, requestID)
		}(i)
	}
	wg.Wait()
}

func getArticleData(requestID int, key string) (interface{}, error) {
	data, _ := getArticleDataFromCache(requestID, key)
	if data == nil {
		v, err, _ := singleFlightGetArticle.Do(key, func() (interface{}, error) {
			return getArticleDataFromDB(requestID, key)
		})
		return v, err
	}
	return data, nil
}

func getArticleDataFromCache(requestID int, key string) ([]int, error) {
	//只是为了演示，这里可以从缓存服务中获取数据
	log.Println("get from cache", requestID)
	return list, nil
}

func getArticleDataFromDB(requestID int, key string) ([]int, error) {
	//只是为了演示，这里可以从数据存储服务中获取数据
	log.Println("get from db requestID", requestID)
	time.Sleep(time.Millisecond * 200)
	list = []int{1, 2, 3}
	return list, nil
}
