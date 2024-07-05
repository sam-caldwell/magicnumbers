package magicnumbers

import (
	"os"
)

// Extract - reads the first few bytes from a file to extract its magic number
func Extract(filePath string, n int) (out []byte, err error) {

	var file *os.File

	defer func() {
		if cerr := file.Close(); err != nil && cerr != nil {
			err = cerr
		}
	}()

	// Open the file
	if file, err = os.Open(filePath); err != nil {
		return nil, err
	}

	// Create a byte slice to hold the magic number
	magicNumber := make([]byte, n)

	// Read the first n bytes
	if _, err = file.Read(magicNumber); err != nil {
		return nil, err
	}

	return magicNumber, nil
}
