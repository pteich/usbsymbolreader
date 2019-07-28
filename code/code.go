package code

import (
	"bytes"
	"errors"
	"strconv"
)

// SymbolType describes the type of the scanned symbol
type SymbolType int

const (
	Code128 SymbolType = iota
	Code39
	Code2of5
	Codabar
	DataMatrix
	EAN8
	EAN13
	UPCA
	UPCE
	PDF417
	PostNet37
	Undefined
)

// Code defines a scanned code
type Code struct {
	input        []byte
	data         bytes.Buffer
	bufferLength byte
	IsNumeric    bool
	Type         SymbolType
}

// New returns a new Code struct from a scanned byte array
func New(input []byte) (*Code, error) {

	bufferLength := input[0]
	if bufferLength == 0 {
		return nil, errors.New("zero length code")
	}

	if int(bufferLength) > len(input) {
		return nil, errors.New("input buffer length exceeded")
	}

	code := &Code{
		input: input[1:bufferLength],
	}

	code.IsNumeric = input[bufferLength-1] != 11
	var offset byte = 2
	if code.IsNumeric {
		offset = 1
	}

	symbolType := input[bufferLength-offset]
	switch symbolType {
	case 24:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = Code128
		}
		break
	case 10:
		if code.IsNumeric {
			code.Type = UPCE
		} else {
			code.Type = Code39
		}
		break
	case 13:
		if code.IsNumeric {
			code.Type = UPCA
		} else {
			code.Type = Code2of5
		}
		break
	case 14:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = Codabar
		}
		break
	case 22:
		if code.IsNumeric {
			code.Type = EAN13
		} else {
			code.Type = Undefined
		}
		break
	case 37:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = PostNet37
		}
		break
	case 46:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = PDF417
		}
		break
	case 50:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = DataMatrix
		}
		break
	case 12:
		if code.IsNumeric {
			code.Type = Undefined
		} else {
			code.Type = EAN8
		}
		break
	default:
		code.Type = Undefined
		break
	}

	codeBufferEnd := bufferLength - 4
	if code.IsNumeric {
		codeBufferEnd = bufferLength - 2
	}

	for i := 3; i < int(codeBufferEnd); i++ {
		if code.IsNumeric && code.Type != UPCE {
			if code.input[i] < 32 {
				code.data.WriteString(strconv.Itoa(int(code.input[i])))
			} else {
				code.data.WriteString(string(code.input[i]))
			}
		} else {
			code.data.WriteByte(code.input[i])
		}
	}

	return code, nil
}

// String is the string representation of the code
func (code *Code) String() string {
	return code.data.String()
}

// Bytes is the byte array representation of the code
func (code *Code) Bytes() []byte {
	return code.data.Bytes()
}
