package main

import (
	"REDIS_IN_ACTION/initrouter"
	"REDIS_IN_ACTION/redisclient"
	"net/http"
	"time"
)

func main() {
	Router := initrouter.InitRouter()

	_ = redisclient.InitRedis() 

	s := &http.Server{
		Addr:           ":8888",
		Handler:        Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	time.Sleep(10 * time.Microsecond)
	
	_ = s.ListenAndServe()
}
