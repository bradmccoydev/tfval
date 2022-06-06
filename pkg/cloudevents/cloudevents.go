package cloudevents

import (
	"encoding/json"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func CreateCloudEvent(payload string) string {
	event := cloudevents.NewEvent()
	event.SetSource("https://github.com/bradmccoydev/tfval")
	event.SetType("example.type")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"pass": payload})
	bytes, err := json.Marshal(event)

	if err != nil {
		fmt.Println(err)
	}

	return string(bytes)
}
