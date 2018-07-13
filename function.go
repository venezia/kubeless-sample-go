package kubeless

import (
	"encoding/json"
	"fmt"
	"github.com/kubeless/kubeless/pkg/functions"
	"time"
)

type Message struct {
	Name string
}

// Foo sample function
func DoStuff(event functions.Event, context functions.Context) (string, error) {
	var input Message

	start := time.Now()

	if event.Data == "" {
		return "You didn't provide a name!", nil
	}
	err := json.Unmarshal([]byte(event.Data), &input)
	if err != nil {
		return "You did not provide valid input.  Please provide JSON that has \"Name\" field only", nil
	}

	result := fmt.Sprintf("Hello %v! This sample function took %v nanoseconds!\n", input.Name, time.Since(start).Nanoseconds())

	return result, nil
}
