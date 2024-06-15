package internal

import (
	"fmt"
	"os"
	"strings"
)

func SetHostname(hostname string) error {
	if err := os.WriteFile("/etc/hostname", []byte(hostname), 0644); err != nil {
		return err
	}

	hosts, err := os.ReadFile("/etc/hosts")
	if err != nil {
		return err
	}

	var newHosts strings.Builder
	lines := strings.Split(string(hosts), "\n")
	for _, line := range lines {
		if strings.Contains(line, "127.0.0.1") {
			newHosts.WriteString(fmt.Sprintf("127.0.0.1\t%s\n", hostname))
		} else {
			newHosts.WriteString(line + "\n")
		}
	}

	err = os.WriteFile("/etc/hosts", []byte(newHosts.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}
