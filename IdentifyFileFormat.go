package magicnumbers

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/errors"
)

// IdentifyFileFormat - identifies the file format based on its magic number
func (m *Map) IdentifyFileFormat(filePath string) (result string, err error) {
	var magicNumber []byte
	if magicNumber, err = Extract(filePath, extractionSize); err != nil {
		return "", err
	}
	for format, magic := range *m {
		if len(magicNumber) > len(magic) && bytes.Equal(magicNumber[:len(AviMagicNumber)], magic) {
			return format, nil
		}
	}
	return "", fmt.Errorf(errors.UnknownFileFormat)
}
