package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	fmt.Println("Connect OK.")
	defer c.Close()

	res, err := c.Do("SET", "user", "chanki")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}

	fmt.Printf("Do return: %v \n", res)

	// get the key
	username, err := redis.String(c.Do("GET", "user"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	/*
	console output:
		Connect OK.
		Do return: OK
		Get mykey: chanki
	*/
}