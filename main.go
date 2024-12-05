package main

import (
	"github.com/go-toast/toast"
	"log"
	"time"
)

func main() {
	go func() {
		for {
			notification := toast.Notification{
				AppID:   "Test",
				Title:   "Test",
				Message: "Test",
				Audio:   toast.Reminder,
			}
			err := notification.Push()
			if err != nil {
				log.Fatalln(err)
			}

			time.Sleep(5 * time.Second)
		}
	}()

	select {}
}
