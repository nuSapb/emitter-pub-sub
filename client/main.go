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
	// 6o5_uAuAx4lV0g-IJVae_9YV2680UkOVAAAAAAAAAAI
	// w-U-eGJeV3alJahXvzc_A4gDTdcj0Gu

	// Create a private link
	// client.CreatePrivateLink("p0R6hGd_G3QTXuuZhofoq7JmdwBF7_lX", "demo/", "s", "optionalHandler", true)
	// client.CreatePrivateLink("p0R6hGd_G3QTXuuZhofoq7JmdwBF7_lX", "demo/", "s", func(_ *emitter.Client, msg emitter.Message) {
	// 	fmt.Printf("[emitter] -> [B] received from private link")
	// })

	// Ask to create a private link
	// fmt.Println("[emitter] <- [B] creating a private link")
	if _, err := client.CreatePrivateLink("yqheMTz7m61mHQbhDyMD35R0wATLaW5T", "demo/", "1", func(_ *emitter.Client, msg emitter.Message) {
		fmt.Printf("[emitter] -> [B] received from private link: '%s' topic: '%s'\n", msg.Payload(), msg.Topic())
	}); err != nil {
		panic(err)
	}
	// fmt.Println("[emitter] -> [B] received link ")

	// Publish to the private link
	// fmt.Println("[emitter] <- [B] publishing to private link")
	client.PublishWithLink("1", "hi from private link")

	// client.Publish("p0R6hGd_G3QTXuuZhofoq7JmdwBF7_lX", "demo/", "hi")
}
