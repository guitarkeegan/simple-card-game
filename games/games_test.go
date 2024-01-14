package games

import (
	"io"
	"os"
	"testing"
)

func TestPlayGame(t *testing.T) {
	// Println in the og function won't match without the \n
	expected := "exiting because there is no game of that name...\n"

	// Redirect standard output to a buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the PlayGame function
	wrongMessage := "hey man nice one!"
	PlayGame(&wrongMessage) // Updated the argument to pass a pointer to a string

	// Close the writer and restore the standard output
	w.Close()
	os.Stdout = old

	// Read the output from the buffer
	out, err := io.ReadAll(r)
	if err != nil {
		panic("read all fail")
	}

	// Convert the output to a string
	output := string(out)

	// Check if the output matches the expected string
	if output != expected {
		t.Errorf("Expected output: %s, but got: %s", expected, output)
	}
}
