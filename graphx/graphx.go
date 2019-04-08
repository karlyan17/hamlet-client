// hamlet-client/graphx.go

package graphx

import (
    "github.com/nsf/termbox-go"
    "hamlet/sessions"
    "log"
    "fmt"
)

type Options struct {
    BG termbox.Attribute
    FG termbox.Attribute
    Input string
    Events chan termbox.Event
    Monitor termbox.Event
}

func Render(session sessions.Session, opts Options) {
    termbox.Clear(opts.BG,opts.BG)
    x := 1
    for _,char := range(fmt.Sprint(session)) {
        termbox.SetCell(x, 1, char, opts.FG, opts.BG)
        x++
    }
    x = 1
    for _,char := range(fmt.Sprint(opts.Monitor)) {
        termbox.SetCell(x, 2, char, opts.FG, opts.BG)
        x++
    }
    termbox.SetCell(1, 3, '>', opts.FG, opts.BG)
    x = 3
    for _,char := range(fmt.Sprint(opts.Input)) {
        termbox.SetCell(x, 3, char, opts.FG, opts.BG)
        x++
    }
    termbox.Flush()
}

func Init(options *Options) {
    err := termbox.Init()
    if err != nil {
        log.Println("ERROR initializing termbox display: ",err)
    }
    termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
    termbox.SetOutputMode(termbox.Output256)
    go func() {
        for {
            options.Events <- termbox.PollEvent()
        }
    }()
}
