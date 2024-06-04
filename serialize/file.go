package serialize

import (
	"fmt"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data := ProtobufToJSON(message)
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file: %w", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := protojson.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtobufToBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data to file: %w", err)
	}

	err = protojson.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal proto message to binary: %w", err)
	}

	return nil
}
