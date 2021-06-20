package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "log"
    "time"
    "math/rand"
    "flag"
)

func main() {
    // 单个value的大小，用string+random生成，单位(B)
    valueByte := 10
    flag.IntVar(&valueByte, "value-bytes", valueByte, "the bytes of redis value")
    flag.Parse()
    // 总共的value大小，单位(B)
    totalBytes := 5 * 1024 * 1024
    // 连接redis-server
    client, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        log.Fatal(err)
    }
    defer client.Close()
    // 密码
    rec, err := client.Do("Auth", "123456")
    if err != nil {
	log.Println(err)
	return
    }
    // info memory
    rec, err = client.Do("info", "memory")
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string(rec.([]byte)))

    // set 5MB value with TTL 60s
    ttl := 60
    valueCount := totalBytes / valueByte
    fmt.Printf("the count of values is: %d\n", valueCount)
    for i := 0; i < valueCount; i++ {
        rec, err = client.Do("Set", i, GetRandomString(valueByte), "EX", ttl)
	if err != nil {
            log.Println(err)
	    return
        }
    }
    rec, err = client.Do("info", "memory")
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(string(rec.([]byte)))
    // waiting for ttl of all value to finish
    fmt.Printf("waiting %d seconds to finish...\n", ttl)
    time.Sleep(time.Duration(ttl) * time.Second)
}

const chars = "0123456789abcdefghijklmnopqrstuvwxyz"
func GetRandomString(l int) string {
    s := make([]byte, l)
    rand.Seed(time.Now().UnixNano())
    for i := range s {
        s[i] = chars[rand.Intn(len(chars))]
    }
    return string(s)
}
