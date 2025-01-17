//go:build ignore
//+build ignore
package main

import (
	"context"
	"fmt"

	"github.com/sjqzhang/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(_ context.Context, e *fsm.Event) { d.enterState(e) },
			//"open": func(ctx context.Context, event *fsm.Event) {
			//	fmt.Println(event.Src,event.Dst)
			//	fmt.Println("xxxxxxxxxxx")
			//},
			//"closed": func(ctx context.Context, event *fsm.Event) {
			//	fmt.Println(event.Src,event.Dst)
			//	fmt.Println(event.Event)
			//	fmt.Println("xxxxxx close")
			//},
		},

	)


	return d
}

func (d *Door) enterState(e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}

func main() {
	door := NewDoor("heaven")


	err := door.FSM.Event(context.Background(), "open")
	if err != nil {
		fmt.Println(err)
	}

	err = door.FSM.Event(context.Background(), "close")
	if err != nil {
		fmt.Println(err)
	}
}
