package response

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Response struct {
	Content []byte
	Err     error
}

func NewChanel() chan *Response {
	return make(chan *Response)
}

func New(content []byte, err error) *Response {
	return &Response{
		Content: content,
		Err:     err,
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
