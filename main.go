package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisConn redis.Conn
)

func main() {
	ep := os.Args[1]
	if ep == "" {
		return
	}
	netConn, err := net.Dial("tcp", ep+":6379")
	if err != nil {
		panic(err)
	}

	defer netConn.Close()

	var redisConnOpt struct {
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	redisConnOpt.ReadTimeout, _ = time.ParseDuration("1s")
	redisConnOpt.WriteTimeout, _ = time.ParseDuration("1s")

	redisConn = redis.NewConn(netConn, redisConnOpt.ReadTimeout, redisConnOpt.WriteTimeout)
	defer redisConn.Close()

	reply, err := redisConn.Do("ping")
	if err != nil {
		log.Fatal(err)
	}

	log.Println(reply)
}
