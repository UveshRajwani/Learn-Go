package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Time Study")
	rnTime := time.Now()
	fmt.Println(rnTime)
	fmt.Println(rnTime.Format("01-02-2006 15:04:05 Monday"))
	myDate := time.Date(2023, time.February, 10, 10, 10, 10, 10, time.UTC)
	fmt.Println(myDate.Format(time.DateTime))

}

//package main
//
//import (
//	"fmt"
//	"github.com/gofiber/fiber/v2/middleware/monitor"
//	"github.com/gofiber/swagger"
//	"log"
//	"time"
//)
//
//func main() {
//	app := fiber.New()
//	app.Get("/docs/*", swagger.HandlerDefault)
//	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))
//	app.Get("/", func(c *fiber.Ctx) error {
//		fmt.Println("hello")
//		return c.SendString(time.Now().String())
//	})
//
//	log.Fatal(app.Listen(":3000"))
//
//}
