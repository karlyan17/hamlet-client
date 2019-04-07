// hamlet-client.go

package main

import (
    "fmt"
    "net"
    "hamlet/sessions"
    "encoding/json"
    "bufio"
    "time"
    //"strings"
    //"os"
    "log"
)

var session sessions.Session
func getState() {
    for {
        var remote_session sessions.Session
        srv_message,err := session.Conn_rw.ReadBytes('\n')
        if err != nil {
            log.Println("ERROR reading message: ",err)
        }
        err = json.Unmarshal(srv_message, &remote_session)
        session.Account = remote_session.Account
        session.ID = remote_session.ID
        log.Print("\r", session)
        time.Sleep(500 * time.Millisecond)
    }

}

func main() {
    conn,_ := net.Dial("tcp", "localhost:6666")
    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    session.Conn_rw = rw
    //reader := bufio.NewReader(os.Stdin)
    getState()

    for {
        srv_message,err := rw.ReadBytes('\n')
        if err != nil {
            log.Println("ERROR reading message: ",err)
        }
        fmt.Println(string(srv_message))
        fmt.Print("> ")
        //name,_ := reader.ReadString('\n')
        //name = strings.Replace(name, "\n", "", -1)
        //message,_ := json.Marshal(name)
        //rw.Write(message)
        //rw.WriteByte('\n')
        //rw.Flush()
    }
}
