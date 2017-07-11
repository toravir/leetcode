/*
   Have you had two TCP servers that you wish you connected it back to back ?
   Here's a TCP Bridge process - that takes in two server socket info
   becomes a client to both of them and reads from one and writes that data to 
   another and vice-versa. It will re-attach to the servers if the server socket
   closes - this will continue the bridge relay until this process is terminated.

This will setup a bridge-relay between two server sockets ep1 and ep2
This will create three threads

 Thread #1: Monitor the socket connection and re-establish ep1 or ep2 if down
 Thread #2: Read from EP1 and write to EP2
 Thread #3: Read from EP2 and write to EP1
*/

package main

import "fmt"
import "net"
import "bufio"
import "io"
import "time"

func connect (ep1, ep2 string) {
   var readep1, readep2 *bufio.Reader
   var writeep1, writeep2 *bufio.Writer

   //connection monitor 
   // Sleep for 30 seconds and
   // peek in the socket if still up, if not try to open that socket.
   go func () {
      e1_is_up := false
      e2_is_up := false
      var rc1, rc2 error
      var conn1, conn2 net.Conn

      for {
          if (!e1_is_up) {
	      conn1, rc1 = net.Dial("tcp", ep1)
              if (rc1 == nil) {
                  fmt.Println("Connected to EP1")
		  readep1  = bufio.NewReader(conn1)
		  writeep1 = bufio.NewWriter(conn1)
		  e1_is_up = true
              } else {
                  fmt.Println(rc1)
              }
          }
          if (!e2_is_up) {
	      conn2, rc2 = net.Dial("tcp", ep2)
              if (rc2 == nil) {
                  fmt.Println("Connected to EP2")
		  readep2  = bufio.NewReader(conn2)
		  writeep2 = bufio.NewWriter(conn2)
		  e2_is_up = true
              } else {
                  fmt.Println(rc2)
              }
          }

          time.Sleep(30*time.Second)
          fmt.Println("Checking if connections are still up...")

          if (e1_is_up) {
              e1_is_up = false
              if ((readep1 != nil) && (writeep1 != nil)) {
                  if _, rc1 := readep1.Peek(1) ; rc1 == io.EOF {
                      conn1.Close()
                  } else {
                      fmt.Println("EP1 is still up..")
                      e1_is_up = true
                  }
              }
          } // if (e1_is_up)
          if (e2_is_up) {
              e2_is_up = false
              if ((readep2 != nil) && (writeep2 != nil)) {
                  if _, rc2 := readep2.Peek(1) ; rc2 == io.EOF {
                      conn2.Close()
                  } else {
                      e2_is_up = true
                      fmt.Println("EP2 is still up..")
                  }
              }
          } // if (e2_is_up)
      } // end of for 
   } ()


   //Read from EP1 and write to EP2
   go func () {
    for {
      for {
          if (readep1 != nil) {
              break
          }
          time.Sleep(5*time.Second)
      }
      buf := make([]byte, 128)
      n1, err := readep1.Read(buf)
      if (n1 == 0 && err == io.EOF) {
         fmt.Println("\n Conn1 Terminated..")
         readep1 = nil
         continue
      }
      fmt.Println("Read ", n1, " bytes: ", string(buf))
      for {
          if (writeep2 != nil) {
              break
          }
          time.Sleep(5*time.Second)
      }
      n1, err = writeep2.Write(buf)
      writeep2.Flush()
      if (n1 != len(buf) || err == io.EOF) {
          fmt.Println("\n Conn2 Terminated..")
          writeep2 = nil
      }
    }
   }()

   //Read from EP2 and write to EP1
   go func () {
    for {
      for {
          if (readep2 != nil) {
              break
          }
          time.Sleep(5*time.Second)
      }
      buf := make([]byte, 128)
      n1, err := readep2.Read(buf)
      if (n1 == 0 && err == io.EOF) {
         fmt.Println("\n Conn2 Terminated..")
         readep2 = nil
         continue
      }
      fmt.Println("Read ", n1, " bytes: ", string(buf))
      for {
          if (writeep1 != nil) {
              break
          }
          time.Sleep(5*time.Second)
      }
      n1, err = writeep1.Write(buf)
      writeep1.Flush()
      if (n1 != len(buf) || err == io.EOF) {
          fmt.Println("\n Conn1 Terminated..")
          writeep1 = nil
      }
    }
   }()

}

func main () {
    connect("localhost:5000", "localhost:6000")
    //let this main thread sleep for a looong time
    for {
        time.Sleep(100000*time.Second)
    }
}
