package funcs

import (
	"os"
	"testing"
)

// go test -v -failfast -run ^TestWriteReadLocal_WriteLocalFile$
func TestWriteReadLocal_WriteLocalFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()
		err := WriteLocalFile("./test_files/write_test.json", []byte("test write"))
		if err != nil {
			t.Errorf("TestWriteReadLocal_WriteLocalFile() error = %v", err)
			return
		}
		b, err := os.ReadFile("./test_files/write_test.json")
		if err != nil {
			t.Errorf("TestWriteReadLocal_WriteLocalFile() error = %v", err)
			return
		}
		if string(b) != "test write" {
			t.Errorf("TestWriteReadLocal_WriteLocalFile() fail result not match %s", string(b))
			return
		}
	})
}

// go test -v -failfast -run ^TestWriteReadLocal_ReadLocalFile$
func TestWriteReadLocal_ReadLocalFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		err := os.WriteFile("./test_files/read_test.json", []byte("test read"), 0777)
		if err != nil {
			t.Errorf("TestWriteReadLocal_ReadLocalFile() error = %v", err)
			return
		}

		b, err := ReadLocalFile("./test_files/read_test.json")
		if err != nil {
			t.Errorf("TestWriteReadLocal_ReadLocalFile() error = %v", err)
			return
		}

		if string(b) != "test read" {
			t.Errorf("TestWriteReadLocal_ReadLocalFile() fail result not match %s", string(b))
			return
		}
	})
}
