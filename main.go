package main
/**
 * <p> @Description TODO
 * @author weijian
 * @date 2020-02-22 18:18
 */

import (
    "fmt"
    _ "github.com/gin-gonic/gin"
    "project2/router"
    "rsc.io/quote"
)

func main() {
    fmt.Println("This works")
    r := router.InitRouter()
    _ = r.Run()
}

func Hello() string {
    return quote.Hello()
}
