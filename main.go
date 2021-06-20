package main

import (
    "bufio"
    "fmt"
    "log"
    "net"
    "os"
)

func hostLookup(host string) {
    ip, err := net.LookupIP(host)
    if err != nil {
        fmt.Println(err)
    }

    for _, hostip := range ip {
        fmt.Printf("%s\t\t\tIN\tA\t%s\n", host, hostip)
    }
}

func main() {
    if len(os.Args[1:]) == 0 {
        fmt.Println("missing argument: file\n file containing fqdn's is required\n")
        os.Exit(1)
    }
    fname := os.Args[1]
    file, err := os.Open(fname)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        host := scanner.Text()
        hostLookup(string(host))
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}
