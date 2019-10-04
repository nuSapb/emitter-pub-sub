package main

import (
	"fmt"

	emitter "github.com/emitter-io/go/v2"
)

func main() {
	client, err := emitter.Connect("tcp://127.0.0.1:8089", func(c *emitter.Client, msg emitter.Message) {
		println("recieve message: ", msg.Payload())
	})
	if err != nil {
		panic(err)
	}

	// create a private link
	fmt.Println("creating a private link")
	if _, err := client.CreatePrivateLink("yqheMTz7m61mHQbhDyMD35R0wATLaW5T", "demo/", "1", func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("received from private link: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	}); err != nil {
		panic(err)
	}
	fmt.Println("received link")

	// Publish to the private link
	fmt.Println("publishing to private link")
	client.PublishWithLink("1", "hi from private link")

}
