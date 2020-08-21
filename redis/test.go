package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	// c-connection
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	} else {
		fmt.Println("连接redis成功！")
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

	/*
		批量处理命令即redis对应的命令：

		mget(适用于string类型)
		mset(适用于string类型)
		hmget(适用于hash类型)
		hmset(适用于hash类型)
	*/

	key := make([]string, 0)
	value := make([]string, 0)

	vec := make([]int, 3)
	for i :=0; i < 3; i++ {
		vec = append(vec, i)
	}
	fmt.Println(vec)
	// [0 0 0 0 1 2]

	for i := 0; i < 5; i++ {
		var tmp string = "a"
		key = append(key, tmp)
		tmp += "b"
		value = append(value, "toutiao")
	}

	res, merr := c.Do("MSET", key, value, "EX", "10")
	if merr != nil {
		fmt.Println("redis set failed:", merr)
	} else {
		// OK
		fmt.Printf("redis mset: %s", res)
		fmt.Println()
	}

	val, merr := c.Do("MGET", key)
	if merr != nil {
		fmt.Println("redis get failed:", merr)
	} else {
		fmt.Printf("redis mget: %s", val)
		fmt.Println()
	}

}
