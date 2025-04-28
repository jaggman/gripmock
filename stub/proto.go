package stub

import (
	"context"
	"encoding/json"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func FindStub(ctx context.Context, service, method string, headers metadata.MD, in, out proto.Message) error {
	pyl := struct {
		Service string            `json:"service"`
		Method  string            `json:"method"`
		Data    interface{}       `json:"data"`
		Headers map[string]string `json:"headers"`
	}{
		Service: service,
		Method:  method,
		Data:    in,
	}
	if headers != nil {
		pyl.Headers = make(map[string]string)
		for header, values := range headers {
			pyl.Headers[header] = values[0]
		}
	}

	byt, err := json.Marshal(pyl)
	if err != nil {
		return err
	}

	stubPyl := findStubPayload{}
	if err := json.Unmarshal(byt, &stubPyl); err != nil {
		return err
	}

	respRPC, err := findStub(&stubPyl)
	if err != nil {
		return err
	}

	if respRPC.Error != "" || respRPC.Code != nil {
		if respRPC.Code == nil {
			abortedCode := codes.Aborted
			respRPC.Code = &abortedCode
		}
		if *respRPC.Code != codes.OK {
			return status.Error(*respRPC.Code, respRPC.Error)
		}
	}

	if respRPC.Headers != nil {
		md := metadata.New(respRPC.Headers)
		grpc.SetHeader(ctx, md)
	}

	if respRPC.Latency != nil {
		time.Sleep(*respRPC.Latency * time.Millisecond)
	}

	data, _ := json.Marshal(respRPC.Data)
	return protojson.Unmarshal(data, out)
}
