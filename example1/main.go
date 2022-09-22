package example1

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

var data = `
{
	"data": [
		{"id": 1, "name": "one"},
		{"id": 2, "name": "two"},
		{"id": 3, "name": "three"},
		{"id": 4, "name": "four"},
		{"id": 5, "name": "five"}
	]
}
`

// A interface
type AData interface {
	Unmarshal(io.Reader) error
}

// B generic interface
type BData[P any] interface {
	Unmarshal(io.Reader) error
	*P
}

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Data struct {
	Items []Item `json:"data"`
}

// Unmarshal is to make Data satisfy the interface and constraint but also do some work.
func (d *Data) Unmarshal(r io.Reader) error {
	return json.NewDecoder(r).Decode(d)
}

func decodeWithAny[T any]() (*T, error) {
	jsonData := strings.NewReader(data)

	var t T
	err := json.NewDecoder(jsonData).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func decodeWithAny[T any]() (*T, error) {
	jsonData := strings.NewReader(data)

	var t T
	err := json.NewDecoder(jsonData).Decode(&t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func decodeWithConstraint[T AData]() (*T, error) {
	jsonData := strings.NewReader(data)

	var t T
	err := json.NewDecoder(jsonData).Decode(&t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func decodeWithConstraintAndMethod[T any, PT BData[T]]() (*T, error) {
	jsonData := strings.NewReader(data)

	var t T
	p := PT(&t)
	err := p.Unmarshal(jsonData)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func plainDecode() (*Data, error) {
	jsonData := strings.NewReader(data)

	var foo Data
	err := json.NewDecoder(jsonData).Decode(&foo)
	if err != nil {
		return nil, err
	}

	return &foo, nil
}
