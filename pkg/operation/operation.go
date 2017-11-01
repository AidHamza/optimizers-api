package operation

import (
	"github.com/AidHamza/optimizers-api/pkg/helpers"
	proto "github.com/golang/protobuf/proto"
)

func NewOperation(fileName string, fileType string) ([]byte, uint64, error) {
	operation := &Operation{
		Id: helpers.RandomID(),
		File: fileName,
	}

	if fileType == "image/jpeg" {
		operation.Type = Operation_JPEG
		// TODO: use mozjpeg instead for better results, but slow
		operation.Command = []*Operation_Command{
			{
				Command: "jpegoptim",
				Flags: []string{"-s", "--max=80", "--dest=", "DEST", "SOURCE"},
			},
		};
	} else {
		//TODO: use pngquant and pngout instead, Check github.com/alexanderteves/pngo
		operation.Type = Operation_PNG
		operation.Command = []*Operation_Command{
			{
				Command: "optipng",
				Flags: []string{"-o2", "SOURCE", "-out", "DEST"},
			},
		};
	}

	operationBytes, err := proto.Marshal(operation)
	if err != nil {
		return []byte{0}, 0, err
	}

	return operationBytes, operation.Id, nil
}