package main

// import "golang.org/x/tour/reader"

type MyReader struct{}

func (reader MyReader) Read(s []byte) (int, error) {
	s = s[:cap(s)]
	for i := range s {
		s[i] = 'A'
	}
	return cap(s), nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main_exe_reader() {
	// reader.Validate(MyReader{})
}
