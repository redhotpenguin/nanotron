package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "github.com/redhotpenguin/nanotron/span"
	"io/ioutil"
)

func main() {
	ct, err := ioutil.ReadFile("./span.json")
	if err != nil {
		panic(err)
	}
	var traceJson map[string]interface{}
	err = json.Unmarshal(ct, &traceJson)
	if err != nil {
		panic(err)
	}
	fmt.Println("json is ", traceJson)

	var span pb.Span
	span.SpanId = 1
	proto.Marshal(span)
	fmt.Println("span is ", span)
}
