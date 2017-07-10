package main

import (
  "log"
  "net"
  "strings"
  "github.com/get-ion/ion"
  "github.com/get-ion/ion/context"
)

// User bind struct
type IP struct {
  ip string `json:"ip"`
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
    conn, err := net.Dial("udp", "8.8.8.8:80")
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    localAddr := conn.LocalAddr().String()
    idx := strings.LastIndex(localAddr, ":")

    return localAddr[0:idx]
  }

func main() {
  app := ion.New()

  // Method:    GET
  // Resource:  http://localhost:8080
  app.Get("/", func(ctx context.Context) {
    addr := GetOutboundIP()
    ctx.StatusCode(ion.StatusOK)
    ctx.JSON(map[string]string{"ip": addr})
  })

  // Start the server using a network address and block.
  app.Run(ion.Addr(":8080"))
}
