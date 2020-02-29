package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

func handler(w http.ResponseWriter, req *http.Request) {
	const timeout = 5
	ctx := req.Context()
	ctx, cancel := context.WithTimeout(ctx, timeout*time.Second)
	ctx, cancel = context.WithCancel(ctx)

	defer cancel()
	// Uncomment this to fake a timeout
	// 	time.Sleep((timeout + 1) * time.Second)

	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintln(os.Stdout, "Success")
		fmt.Fprintln(w, "Success")
	case <-ctx.Done():
		fmt.Fprintln(os.Stderr, "Request cancelled")
		fmt.Fprintln(os.Stderr, ctx.Err())
		fmt.Fprintln(w, ctx.Err())
	}
}

// Init : Initialize server
func Init() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
