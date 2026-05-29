package main

import (
	"log"
	"fmt"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	greeting "temporalab/one/greeting"
)

func main() {
	fmt.Println("Running worker ...")

	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client: ", err)
	}
	defer c.Close()

	w := worker.New(c, greeting.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(greeting.SayHelloWorkflow)
	w.RegisterActivity(greeting.Greet)
	w.RegisterActivity(greeting.Log)
	w.RegisterActivity(greeting.Dummy)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
