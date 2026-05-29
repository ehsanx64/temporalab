package greeting

import (
	"context"
	"fmt"
	"log"
)

func Greet(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello %s", name), nil
}

func Log(ctx context.Context) (string, error) {
	log.Printf("Log called")
	return "", nil
}

func Dummy(ctx context.Context) (error) {
	log.Printf("Dummy called")
	return nil
}
