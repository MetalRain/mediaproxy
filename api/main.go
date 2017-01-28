package main

import (
        "os"
        "io"
	"fmt"
	"crypto/rand"
	"github.com/kataras/iris"
	"gopkg.in/redis.v5"
)

func main() {

    addr := os.Getenv("REDIS_ADDR")
    passwd := os.Getenv("REDIS_PASSWORD")
    db := 0

    client := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: passwd,
        DB:       db,
    })

    _, err := client.Ping().Result()
    if err != nil {
	fmt.Fprintf(os.Stderr, "Could not connect to Redis at %s", addr)
        return
    }

    iris.Set(iris.OptionGzip(true), iris.OptionCharset("UTF-8"))

    iris.Get("/", func(ctx *iris.Context) {
	ctx.Writef("Mediaproxy")
    })

    iris.Post("/subscribe/:service", func(ctx *iris.Context) {
	id, err := newUUID()
	if err != nil {
	  ctx.Writef("500, cannot generate UUID")
	  return
        }
        service := ctx.Param("service")
	key := fmt.Sprintf("sub:%s:%s", service, id)
	client.Set(key, "{}", 0)
	ctx.Writef("OK " + key) 
    })

    // start the server
    iris.Listen(":8080")
}

func newUUID() (string, error) {
    uuid := make([]byte, 16)
    n, err := io.ReadFull(rand.Reader, uuid)
    if n != len(uuid) || err != nil {
        return "", err
    }
    // variant bits; see section 4.1.1
    uuid[8] = uuid[8]&^0xc0 | 0x80
    // version 4 (pseudo-random); see section 4.1.3
    uuid[6] = uuid[6]&^0xf0 | 0x40
    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
