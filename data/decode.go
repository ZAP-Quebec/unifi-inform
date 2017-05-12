package data

import (
	"encoding/json"
)

func Unmarshal(data []byte) (Message, error) {

	var result map[string]string
	json.Unmarshal(data, &result)
	msgType := result["_type"]
	switch msgType {
	case "setparam":
		msg := &SetParam{
			httpResponse: ResponseFromHttpCode(200).(httpResponse),
		}
		err := msg.unmarshalMap(result)
		return msg, err
	}

	return nil, nil
}
