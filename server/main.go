package main

import (
	emitter "github.com/emitter-io/go/v2"
)

func main() {
	client, err := emitter.Connect("tcp://127.0.0.1:8089", func(c *emitter.Client, msg emitter.Message) {
		println("recieve message: ", msg.Topic(), string(msg.Payload()))
	})

	if err != nil {
		panic(err)
	}

	client.Subscribe("ROPjRa04WJlvTDlP9njzsF45rccWiaOa", "demo/", nil)

	for {

	}

}
