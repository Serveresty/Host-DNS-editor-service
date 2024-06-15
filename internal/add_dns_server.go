package internal

import (
	"fmt"
	"os"
	"strings"
)

func AddDNSServer(newServer string) error {
	const path = "/etc/resolv.conf"

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if strings.Contains(string(content), newServer) {
		return nil
	}

	newContent := fmt.Sprintf("%s\nnameserver %s\n", string(content), newServer)
	err = os.WriteFile(path, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
