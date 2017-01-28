package main

import "github.com/kataras/iris"

func main() {
    // declare the routes
    iris.Get("/", testGet)
    //iris.Post("/login", testPost)
    //iris.Put("/add", testPut)
    //iris.Delete("/remove", testDelete)
    //iris.Head("/testHead", testHead)
    //iris.Patch("/testPatch", testPatch)
    //iris.Options("/testOptions", testOptions)
    //iris.Connect("/testConnect", testConnect)
    //iris.Trace("/testTrace", testTrace)

    // start the server
    iris.Listen(":8080")
}

func testGet(ctx *iris.Context) {
    ctx.Writef("Hello")
}
