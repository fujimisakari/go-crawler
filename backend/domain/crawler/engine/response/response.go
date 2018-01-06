package response

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Response struct {
	Content []byte
}

func New(content []byte) *Response {
	return &Response{
		Content: content,
	}
}

func (r *Response) ToString() string {
	return string(r.Content)
}

func (r *Response) ToJson() ([]map[string]interface{}, error) {
	jsonBody := []map[string]interface{}{}
	err := json.Unmarshal(r.Content, &jsonBody)
	if err != nil {
		return nil, errors.Wrap(err, "response json decode error")
	}
	return jsonBody, nil
}
