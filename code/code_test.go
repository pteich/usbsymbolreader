package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {

	testData := []struct {
		name     string
		input    []byte
		expected string
		codeType SymbolType
	}{
		{name: "testcase 1 code39", input: []byte{9, 16, 3, 0, 81, 49, 0, 10, 11}, expected: "Q1", codeType: Code39},
		{name: "testcase 2 code39", input: []byte{9, 16, 3, 0, 67, 78, 0, 10, 11}, expected: "CN", codeType: Code39},
		{name: "testcase 3 code39", input: []byte{15, 16, 3, 0, 68, 49, 50, 77, 65, 82, 49, 57, 0, 10, 11}, expected: "D12MAR19", codeType: Code39},
		{name: "testcase 4 ean13 ascii", input: []byte{18, 16, 3, 0, 51, 53, 55, 52, 54, 54, 49, 52, 48, 52, 50, 54, 52, 22}, expected: "3574661404264", codeType: EAN13},
		{name: "testcase 5 ean13 ascii", input: []byte{18, 16, 3, 0, 52, 48, 50, 54, 54, 48, 48, 56, 57, 49, 54, 48, 56, 22}, expected: "4026600891608", codeType: EAN13},
		{name: "testcase 6 ean13 numbers", input: []byte{18, 16, 3, 0, 3, 5, 7, 4, 6, 6, 1, 4, 0, 4, 2, 6, 4, 22}, expected: "3574661404264", codeType: EAN13},
		{name: "testcase 7 datamatrix", input: []byte{44, 16, 3, 0, 66, 67, 83, 84, 85, 81, 66, 77, 84, 90, 70, 70, 66, 66, 71, 73, 68, 83, 52, 73, 75, 77, 82, 70, 90, 77, 124, 65, 49, 56, 48, 55, 49, 56, 48, 49, 124, 0, 50, 11}, expected: "BCSTUQBMTZFFBBGIDS4IKMRFZM|A18071801|", codeType: DataMatrix},
		{name: "testcase 8 datamatrix", input: []byte{20, 16, 3, 0, 65, 66, 67, 32, 49, 50, 51, 52, 53, 54, 55, 56, 57, 0, 50, 11}, expected: "ABC 123456789", codeType: DataMatrix},
	}

	for _, test := range testData {

		t.Run(test.name, func(t *testing.T) {
			scannedCode, err := New(test.input)
			assert.NoError(t, err)
			assert.Equal(t, test.expected, scannedCode.String())
			assert.Equal(t, test.codeType, scannedCode.Type)
			t.Log(scannedCode.IsNumeric)
		})

	}

}
