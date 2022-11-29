package fastcgi

import "bytes"

type recType uint8

const (
	typeBeginRequest recType = iota + 1
	typeAbortRequest
	typeEndRequest
	typeParams
	typeStdin
	typeStdout
	typeStderr
	typeData
	typeGetValues
	typeGetValuesResult
	typeUnknownType
)

// Role for fastcgi application in spec
type Role uint16

const (
	roleResponder = iota + 1 // only Responders are implemented.
	roleAuthorizer
	roleFilter
)

func nameValuePair11(nameData, valueData string) []byte {
	return bytes.Join(
		[][]byte{
			{byte(len(nameData)), byte(len(valueData))},
			[]byte(nameData),
			[]byte(valueData),
		},
		nil,
	)
}

func makeRecord(
	recordType recType,
	requestID uint16,
	contentData []byte,
) []byte {
	requestIDB1 := byte(requestID >> 8)
	requestIDB0 := byte(requestID)
	contentLength := len(contentData)
	contentLengthB1 := byte(contentLength >> 8)
	contentLengthB0 := byte(contentLength)
	return bytes.Join([][]byte{
		{1, byte(recordType), requestIDB1, requestIDB0, contentLengthB1,
			contentLengthB0, 0, 0},
		contentData,
	}, nil)
}

// NewFastCGIRecord -
func NewFastCGIRecord(env map[string]string, data []byte, reqID ...int) []byte {
	var _reqID uint16 = 1
	if len(reqID) > 0 && reqID[0] > 0 {
		_reqID = uint16(reqID[0])
	}
	params := make([][]byte, 0)
	for k, v := range env {
		params = append(params, makeRecord(typeParams, _reqID, nameValuePair11(k, v)))
	}
	return bytes.Join([][]byte{
		// set up request 1
		makeRecord(typeBeginRequest, _reqID, []byte{0, byte(roleResponder), 0, 0, 0, 0, 0, 0}),
		// add required parameters
		bytes.Join(params, nil),
		makeRecord(typeParams, _reqID, nil),
		// begin sending body
		makeRecord(typeStdin, _reqID, data),
		makeRecord(typeEndRequest, _reqID, nil),
	}, nil)
}
