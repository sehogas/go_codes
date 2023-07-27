package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Session struct {
	Issued   int64  `json:"issued"`
	Expires  int64  `json:"expires"`
	Username string `json:"username"`
	UserId   int64  `json:"user_id"`
}

func main() {
	log.SetPrefix("redis client|")
	log.SetFlags(log.Ltime)

	host := flag.String("host", "127.0.0.1", "Host redis")
	port := flag.String("port", "6379", "Port redis")
	pass := flag.String("pass", "password", "Password redis")
	flag.Parse()

	addr := *host + ":" + *port

	log.Println("Connect to the server at", addr)
	cache := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: *pass,
	})
	ping := cache.Ping(context.Background())

	fmt.Print("PING -> ")
	result, err := ping.Result()
	if err != nil {
		fmt.Println("ERROR")
		log.Fatalf("Server connection error: %s\n", err)
	}
	defer cache.Close()
	fmt.Println(result)

	issued := time.Now()
	expires := time.Now().Add(24 * time.Hour)
	timeout := expires.Sub(issued)

	// data example
	sessionData := &Session{
		Issued:   issued.Unix(),
		Expires:  expires.Unix(),
		Username: "PEPE",
		UserId:   1,
	}

	sessionJson, err := json.Marshal(sessionData)
	if err != nil {
		log.Fatalf("encoding error: %s\n", err)
	}
	sessionKey := "1"

	// Write
	err = cache.Set(context.Background(), sessionKey, sessionJson, timeout).Err()
	if err != nil {
		log.Fatalf("error writing data: %s\n", err)
	}
	log.Printf("Set key [%s]\n", sessionKey)

	// Read
	val, err := cache.Get(context.Background(), sessionKey).Bytes()
	if err != nil {
		log.Fatalf("error reading data: %s\n", err)
	}

	var data Session
	err = json.Unmarshal(val, &data)
	if err != nil {
		log.Fatalf("decoding error: %s\n", err)
	}

	log.Printf("Get key [%s] - Issued: %v, Expires: %v, Username: %s, UserId: %d\n",
		sessionKey, time.Unix(data.Issued, 0), time.Unix(data.Expires, 0), data.Username, data.UserId)

}
