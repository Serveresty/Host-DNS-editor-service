package internal

import (
	"os"
	"strings"
)

func RemoveDNCServer(server string) error {
	const path = "/etc/resolv.conf"

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var newContent strings.Builder
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		if !strings.Contains(line, server) {
			newContent.WriteString(line + "\n")
		}
	}

	err = os.WriteFile(path, []byte(newContent.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}
