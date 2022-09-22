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

type AData interface {
	Unmarshal(io.Reader) error
}

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

func (d *Data) Unmarshal(r io.Reader) error {
	err := json.NewDecoder(r).Decode(d)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
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
