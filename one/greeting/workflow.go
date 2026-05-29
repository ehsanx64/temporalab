package greeting

import (
	"time"
	"log"

	"go.temporal.io/sdk/workflow"
)

func SayHelloWorkflow(ctx workflow.Context, name string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 10,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, Greet, name).Get(ctx, &result)
	if err != nil {
		log.Printf("Failed to execute Greet activity with name: %s", name)
		return "", nil
	}

	log.Println("Running Log activity ...")
	err = workflow.ExecuteActivity(ctx, Log).Get(ctx, nil)
	if err != nil {
		log.Println("Failed to execute Log activity")
		return "", nil
	}

	_ = workflow.ExecuteActivity(ctx, Dummy).Get(ctx, nil)

	return result, nil
}
