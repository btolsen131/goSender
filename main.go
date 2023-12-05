package main

import (
  "fmt"
  "net"
  "os"
  "bufio"
)

func main() {

  var server string
  var message string

  if len(os.Args) == 3 {
    server = os.Args[1]
    message = os.Args[2]
  } else {
    //default is localhost:1234 "Hello this is a test"
    server = "localhost:1234"
    message = "Hello, this is a test"
  }

  //make connection
  conn, err := net.Dial("tcp", server)
  if err != nil {
    fmt.Println("Error connecting to server: ", err)
    os.Exit(1)
  }

  defer conn.Close()

  scanner := bufio.NewScanner(os.Stdin)
  
  for { 
    fmt.Println("Enter message (or 'exit' to quit): ")
    scanner.Scan()
    message = scanner.Text()

    if message == "exit" {
      break
    }

    _ , err = conn.Write([]byte(message))
    if err != nil {
      fmt.Println("Error sending message: ", err)
      os.Exit(1)
    }

    fmt.Println("Message sent to server!")
  }

  fmt.Println("Sender exiting.")

}
