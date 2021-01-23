package tracemate

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	otco "github.com/open-telemetry/opentelemetry-proto/gen/go/common/v1"
	ott "github.com/open-telemetry/opentelemetry-proto/gen/go/trace/v1"
	"time"
)

// trace
type Trace struct {
	Timestamp string `json:"@timestamp"`
	Span      struct {
		Id             string
		Transaction_id string
		Trace_id       string
		Parent_id      string
		Name           string
		Duration       struct {
			US int
		}
		Http struct {
			Url struct {
				Original string
			}
			Response struct {
				Status int64 `json:"status_code"`
			}
		}
	}
}

func JsonToProto(ct []byte) *ott.Span {
	var trace Trace
	err := json.Unmarshal(ct, &trace)
	if err != nil {
		panic(err)
	}

	var oSpan ott.Span

	oSpan.SpanId = []byte(trace.Span.Id)
	oSpan.TraceId = []byte(trace.Span.Trace_id)
	oSpan.ParentSpanId = []byte(trace.Span.Parent_id)
	oSpan.Name = trace.Span.Name

	timeStart, _ := time.Parse(time.RFC3339, trace.Timestamp)
	oSpan.StartTimeUnixNano = uint64(timeStart.UnixNano())

	// parse the microsecond duration, add to the start
	u, _ := time.ParseDuration(fmt.Sprintf("%dµs", trace.Span.Duration.US))
	oSpan.EndTimeUnixNano = uint64(timeStart.Add(u).UnixNano())

	// add attributes
	oSpan.Attributes = []*otco.KeyValue{
		{
			Key:   "status",
			Value: &otco.AnyValue{Value: &otco.AnyValue_IntValue{IntValue: trace.Span.Http.Response.Status}},
		},
		{
			Key:   "original_url",
			Value: &otco.AnyValue{Value: &otco.AnyValue_StringValue{StringValue: trace.Span.Http.Url.Original}}},
	}

	// serialize
	proto.Marshal(&oSpan)

	return &oSpan
}
