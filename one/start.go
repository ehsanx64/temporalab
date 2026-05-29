package main

import (
	"context"
	"log"
	"fmt"
	"os"

	"go.temporal.io/sdk/client"
	"github.com/google/uuid"

	"temporalab/one/greeting"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client: ", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID: fmt.Sprintf("%s::%s", greeting.OneStarterID, uuid.New().String()),
		TaskQueue: greeting.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, 
		greeting.SayHelloWorkflow, os.Args[1])
	if err != nil {
		log.Fatalln("Unable to execute the workflow: ", err)
	}

	log.Println("Started workflow", "WorkflowID", we.GetID(), 
		"RunID", we.GetRunID())

	
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result: ", err)
	}

	log.Println("Workflow result: ", result)
}
