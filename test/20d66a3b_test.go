package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test20d66a3b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./20d66a3b.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./20d66a3b.rb:::6:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
