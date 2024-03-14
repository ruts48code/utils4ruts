package utils4ruts

import (
	"fmt"
	"os"
	"strings"
	"time"

	random "github.com/ruts48code/random4ruts"
)

func ReadFile(f string) []byte {
	content, err := os.ReadFile(f)
	if err != nil {
		return []byte("")
	}
	return content
}

func MakeString(s string) string {
	return fmt.Sprintf("%s", s)
}

func RandomArrayString(data []string) []string {
	output := make([]string, 0)
	input := make([]string, len(data))
	copy(input, data)
	for {
		if len(input) == 0 {
			return output
		}
		choose := int(random.GetRandomInt(int64(len(input))))
		output = append(output, input[choose])
		input = DeleteArrayString(input, choose)
	}
}

func DeleteArrayString(data []string, n int) []string {
	return append(data[:n], data[n+1:]...)
}

func GetTimeStamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05")
}

func GetDateStamp(t time.Time) string {
	return t.Format("2006-01-02")
}

func NormalizeUsername(username string) string {
	return strings.TrimSpace(strings.ToLower(username))
}

func NormalizedEloginToken(token string) string {
	return strings.Replace(token, "_", ".", 1)
}
