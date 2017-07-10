package main

import (
  "fmt"
  "log"
  "net"
  "strings"
)

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
  fmt.Println(GetOutboundIP())
}
