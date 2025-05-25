package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gordonklaus/portaudio"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	hook "github.com/robotn/gohook"
)

type App struct {
	ctx             context.Context
	keypressChannel chan string
}

func NewApp() *App {
	return &App{}
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx

	// hook.Register(hook.KeyUp, []string{"w"}, func(e hook.Event) {
	// 	_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
	// 		Type:    runtime.InfoDialog,
	// 		Title:   "Hook works!",
	// 		Message: "You pressed W!",
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// })

	b.keypressChannel = make(chan string)
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		b.keypressChannel <- "w"
	})

	s := hook.Start()
	hook.Process(s)

	go b.eventProcessor()

	if err := portaudio.Initialize(); err != nil {
		panic(err)
	}

	audioDevices, err := portaudio.Devices()
	if err != nil {
		panic(err)
	}

	go func() {
		var deviceInfoString []string
		for _, device := range audioDevices {
			deviceInfoString = append(deviceInfoString, device.Name)
		}

		_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Title:   "Audio devices!",
			Message: fmt.Sprintf("Audio Devices:\n%s", strings.Join(deviceInfoString, "\n")),
		})

		if err != nil {
			panic(err)
		}
	}()
}

func (b *App) eventProcessor() {
	var lastKey string
	// for keypress := range b.keypressChannel {
	// 	if keypress == lastKey {
	// 		continue
	// 	}
	// 	lastKey = keypress

	// 	_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
	// 		Type:    runtime.InfoDialog,
	// 		Title:   "Hook works!",
	// 		Message: fmt.Sprintf("You pressed %s!", keypress),
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	timeOut := make(chan string)

	for {
		select {
		case keypress := <-b.keypressChannel:
			if keypress == lastKey {
				continue
			}
			lastKey = keypress
			go func() {
				<-time.After(1 * time.Second)
				timeOut <- keypress
			}()

			go func() {
				_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
					Type:    runtime.InfoDialog,
					Title:   "Hook works!",
					Message: fmt.Sprintf("You pressed %s!", keypress),
				})

				if err != nil {
					panic(err)
				}
			}()
		case <-timeOut:
			lastKey = ""
		}
	}
}

func (b *App) shutdown(ctx context.Context) {
	// Perform your teardown here
	hook.End()
	close(b.keypressChannel)
	portaudio.Terminate()
}

func (b *App) domReady(ctx context.Context) {
	// b.ShowDialog()
}

func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (b *App) ShowDialog() {
	_, err := runtime.MessageDialog(b.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Title:   "Native Dialog from Go",
		Message: "This is a Native Dialog send from Go.",
	})

	if err != nil {
		panic(err)
	}
}
