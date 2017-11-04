package operation

import (
	proto "github.com/golang/protobuf/proto"
)

func NewOperation(opId uint64, fileName string, fileType string) ([]byte, error) {
	operation := &Operation{
		Id: opId,
		File: fileName,
	}

	if fileType == "image/jpeg" {
		operation.Type = Operation_JPEG
		// TODO: use mozjpeg instead for better results, but slow
		operation.Command = []*Operation_Command{
			{
				Command: "jpegoptim",
				Flags: []string{"-s", "--max=80", "--stdout", "--stdin"},
			},
		};
	} else {
		//TODO: use pngquant and pngout instead, Check github.com/alexanderteves/pngo
		operation.Type = Operation_PNG
		operation.Command = []*Operation_Command{
			{
				Command: "pngquant",
				Flags: []string{"--strip", "--quality=65-80", "-"},
			},
		};
	}

	operationBytes, err := proto.Marshal(operation)
	if err != nil {
		return []byte{0}, err
	}

	return operationBytes, nil
}