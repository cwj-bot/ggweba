package main
/**
 * <p> @Description TODO
 * @author weijian
 * @date 2020-02-22 18:18
 */

import (
    "fmt"
    "project2/router"
)

func main() {
    fmt.Println("start server ...")
    r := router.InitRouter()
    _ = r.Run()
    //address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
    //server := endless.NewServer(address, r)
    //server.BeforeBegin = func(add string) {
    //    log.Printf("Actual pid is %d", syscall.Getpid())
    //}
    //
    //err := server.ListenAndServe()
    //if err != nil {
    //    log.Printf("Server err: %v", err)
    //}
}


