package main

import (
	"fmt"
	"os"
	"os/exec"
)

func getTerminalSize() (int, int, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	var rows, cols int
	fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	return rows, cols, nil
}

func main() {

}
