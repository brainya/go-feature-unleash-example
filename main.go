package main

import (
	"fmt"
	"net/http"

	"github.com/Unleash/unleash-client-go/v3"
)

func init() {
	unleash.Initialize(
		unleash.WithListener(&unleash.DebugListener{}),
		unleash.WithAppName("default"),
		unleash.WithUrl("http://localhost:4242/api/"),
		unleash.WithCustomHeaders(http.Header{"Authorization": {"default:development.unleash-insecure-api-token"}}),
	)
}

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Feature is enabled %t\n", unleash.IsEnabled("my-awesome-feature-toggle"))
}
