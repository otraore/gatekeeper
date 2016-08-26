# gatekeeper
[![GoDoc](https://godoc.org/github.com/otraore/gatekeeper?status.svg)](https://godoc.org/github.com/otraore/gatekeeper)
[![Go Report Card](https://goreportcard.com/badge/github.com/otraore/gatekeeper)](https://goreportcard.com/report/github.com/otraore/gatekeeper)
![License](https://img.shields.io/badge/License-MIT-blue.svg)

Hydra middleware for golang web frameworks. Inspired by [gin-hydra](https://github.com/janekolszak/gin-hydra), it current has support for [gin](https://github.com/gin-gonic/gin), [echo](https://github.com/labstack/echo), and [goa](https://github.com/goadesign/goa). More framework support is planned.

## Install
``` bash
go get -u github.com/otraore/gatekeeper
```

## Gin
``` go
import (
    "github.com/gin-gonic/gin"
    "github.com/ory-am/hydra/firewall"
    hydra "github.com/ory-am/hydra/sdk"
    "github.com/otraore/gatekeeper"
)

func handler(c *gin.Context) {
	ctx := c.Get("hydra").(*firewall.Context)
	// Now you can access ctx.Subject etc.
}

func main(){
	// Initialize Hydra
	hc, err := hydra.Connect(
		hydra.ClientID("..."),
		hydra.ClientSecret("..."),
		hydra.ClusterURL("..."),
	)

	if err != nil {
		panic(err)
	}
	
        // Create a gatekeeper instance for Gin
	gk := gatekeeper.NewGin(hc)

 	r := gin.Default()
	r.GET("/protected", gk.ScopesRequired("scope1", "scope2"), handler)
	r.Run()
}
```
## Echo
``` go
import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
    "github.com/ory-am/hydra/firewall"
    hydra "github.com/ory-am/hydra/sdk"
    "github.com/otraore/gatekeeper"
)

func handler(c *gin.Context) {
	ctx := c.Get("hydra").(*firewall.Context)
	// Now you can access ctx.Subject etc.
}

func main(){
	// Initialize Hydra
	hc, err := hydra.Connect(
		hydra.ClientID("..."),
		hydra.ClientSecret("..."),
		hydra.ClusterURL("..."),
	)

	if err != nil {
		panic(err)
	}
	
        // Create a gatekeeper instance for Echo
	gk := gatekeeper.NewGin(hc)

 	e := echo.Default()
	e.GET("/protected", handler, gk.ScopesRequired("scope1", "scope2"),)

	e.Run(standard.New(":8080"))
}
```
## Goa

Example coming soon

## License

MIT