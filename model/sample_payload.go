package model

import (
	"encoding/json"

	"github.com/google/jsonapi"
)

const PayloadType = "payload"

type SamplePayload struct {
	Payload string `json:"payload"`
}

func (s *SamplePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(jsonapi.Node{
		Type: PayloadType,
		Attributes: map[string]interface{}{
			"payload": s.Payload,
		},
	})
}

func (s *SamplePayload) UnmarshalJSON(body []byte) error {
	var payload jsonapi.OnePayload
	if err := json.Unmarshal(body, &payload); err != nil {
		return err
	}

	if payload, ok := payload.Data.Attributes["payload"].(string); ok {
		s.Payload = payload
	}

	return nil
}
