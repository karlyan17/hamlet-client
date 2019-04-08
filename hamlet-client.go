// hamlet-client.go

package main

import (
    "hamlet/sessions"
    "hamlet-client/graphx"
    "net"
    "encoding/json"
    "bufio"
    "time"
    "os"
    //"strings"
    //"os"
    "log"
    "github.com/nsf/termbox-go"
)

//global variables (yeah those are bad, I know)
var session sessions.Session
var gfx_options graphx.Options

//termbox layout
const backgroundColor = 1
const textColor = 2


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
        time.Sleep(500 * time.Millisecond)
    }

}

func InputHandler(opts *graphx.Options) {
    for {
        event := <-opts.Events
        opts.Monitor = event
        if event.Key == 27 {
            os.Exit(0)
        } else if event.Ch != 0 {
            opts.Input += string(event.Ch)
        }
    }
}

func main() {
    //start termbox display thingy
    gfx_options = graphx.Options{
        BG: backgroundColor,
        FG: textColor,
    }
    gfx_options.Events = make(chan termbox.Event)
    graphx.Init(&gfx_options)

    //establish connection
    conn,_ := net.Dial("tcp", "localhost:6666")
    rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
    session.Conn_rw = rw
    go getState()
    go InputHandler(&gfx_options)

    for {
        graphx.Render(session, gfx_options)
    }
}
