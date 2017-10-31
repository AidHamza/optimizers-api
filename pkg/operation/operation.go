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
		operation.Command = []*Operation_Command{
			{
				Command: "jpegoptim",
				Flags: []string{"-s", "--max=80", "--dest=", "DEST", "SOURCE"},
			},
		};
	} else {
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