package main

import (
    "net/http"

    "github.com/labstack/echo"
	"github.com/keploy/go-sdk/keploy"
    "github.com/keploy/go-sdk/integrations/01project"
	"github.com/keploy/go-sdk/integrations/kecho/v4"
)

type HelloWorld struct {
    Message string `json:"message"`
}

func main() {
    e := echo.New()
    e.GET("/hello", Greetings)
    e.Logger.Fatal(e.Start(":3000"))

}

func Greetings(c echo.Context) error {
    return c.JSON(http.StatusOK, HelloWorld{
        Message: "Hello World",
    })

	port := "6789"
	k := keploy.New(keploy.Config{
		App: keploy.AppConfig{
			Name: "my-app",
			Port: port,
		},
		Server: keploy.ServerConfig{
			URL: "http://localhost:6789/api",
		},
	})

	e := echo.New()
	port := "6789"
	k := keploy.New(keploy.Config{
	  App: keploy.AppConfig{
		  Name: "my-app",
		  Port: port,
	  },
	  Server: keploy.ServerConfig{
		  URL: "http://localhost:6789/api",
	  },
	})
	kecho.EchoV4(k, e)
	e.Start(":" + port)
}
