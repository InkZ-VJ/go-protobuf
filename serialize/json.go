package serialize

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) string {
	marshaller := protojson.MarshalOptions{
		UseEnumNumbers:    false,
		EmitDefaultValues: true,
		Indent:            "  ",
		UseProtoNames:     true,
	}
	return marshaller.Format(message)
}
