package data

import (
	"encoding/json"
	"strconv"
	"strings"
)

type SetParam struct {
	httpResponse
	ManagementConfig `json:"mgmt_cfg"`
	ServerTime       int `json:"server_time_in_utc"`
}

func (msg *SetParam) unmarshalMap(data map[string]string) (err error) {
	msg.ManagementConfig = make(ManagementConfig)
	if err = msg.ManagementConfig.unmarshalStr(data["mgmt_cfg"]); err != nil {
		return err
	}
	msg.ServerTime, err = strconv.Atoi(data["server_time_in_utc"])
	if err != nil {
		return err
	}
	return nil
}

func (msg *SetParam) Marshal() []byte {
	res, _ := json.Marshal(msg)
	return res
}

func (msg *SetParam) String() string {
	return string(msg.Marshal())
}

type ManagementConfig map[string]string

func (m ManagementConfig) unmarshalStr(str string) error {
	for _, line := range strings.Split(str, "\n") {
		i := strings.IndexByte(line, '=')
		if i != -1 {
			m[line[0:i]] = line[i:]
		}
	}
	return nil
}

func (m ManagementConfig) MarshalJSON() ([]byte, error) {
	str := ""
	for k, v := range m {
		str = str + k + "=" + v + "\n"
	}
	return json.Marshal(str)
}
