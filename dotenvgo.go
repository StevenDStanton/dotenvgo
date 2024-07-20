package dotenvgo

import (
	"io"
	"os"
	"strings"
)

type Vault map[string]string
type ReturnType int

const (
	Enviroment ReturnType = iota
	Map
	Both
)

func Load(returnType ReturnType, params ...string) (Vault, error) {
	vault := make(Vault)
	filePath := parseFileLocation(params...)
	content, err := fetchFile(filePath)
	if err != nil {
		return vault, err
	}

	lines := normalizeLineEndings(content)
	if returnType == Enviroment || returnType == Both {
		saveFileContentToEnviroment(lines)
	}

	if returnType == Map || returnType == Both {
		vault.parseFileContentToMap(lines)
	}

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

func (v Vault) parseFileContentToMap(lines []string) {
	for _, line := range lines {
		key, value, exists := parseKeyValue(line)
		if exists {
			v[key] = value
		}
	}
}

func saveFileContentToEnviroment(lines []string) {
	for _, line := range lines {
		key, value, exists := parseKeyValue(line)
		if exists {
			os.Setenv(key, value)
		}
	}
}

func parseKeyValue(line string) (key, value string, exists bool) {
	lineNoComments := strings.SplitN(line, "#", 2)
	parts := strings.SplitN(lineNoComments[0], "=", 2)
	if len(parts) == 2 {
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		return key, value, true
	}

	return "", "", false
}

func normalizeLineEndings(content []byte) []string {
	normalizedContent := strings.ReplaceAll(string(content), "\r\n", "\n")
	normalizedContent = strings.ReplaceAll(normalizedContent, "\r", "\n")
	lines := strings.Split(normalizedContent, "\n")
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines
}
