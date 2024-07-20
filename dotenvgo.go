package dotenvgo

import (
	"io"
	"os"
	"strings"
)

type Vault map[string]string

func Load(params ...string) (Vault, error) {
	vault := make(Vault)
	filePath := parseFileLocation(params...)
	content, err := fetchFile(filePath)
	if err != nil {
		return vault, err
	}

	vault.parseFileContentToMap(content)

	return vault, nil
}

func parseFileLocation(params ...string) string {
	filePath := ".env"
	if len(params) > 0 {
		filePath = params[0] + filePath
	}
	return filePath
}

func fetchFile(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return []byte(""), err
	}

	return content, nil
}

func (v Vault) parseFileContentToMap(content []byte) {
	normalizedContent := normalizeLineEndings(content)
	lines := strings.Split(normalizedContent, "\n")
	for _, line := range lines {
		lineNoComments := strings.SplitN(line, "#", 2)
		parts := strings.SplitN(lineNoComments[0], "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			v[key] = value
		}

	}
}

func normalizeLineEndings(content []byte) string {
	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	normalizedContent = strings.ReplaceAll(normalizedContent, "\r", "\n")
	return normalizedContent
}
