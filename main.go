package utils4ruts

import (
	"fmt"
	"os"
	"strings"
	"time"

	random "github.com/ruts48code/random4ruts"
	"gopkg.in/yaml.v3"
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

func CheckEpassportType(namex string) string {

	name := strings.ToLower(strings.TrimSpace(namex))

	if len(name) < 2 {
		return "error"
	}

	switch name[0] {
	case 's':
		switch name[1] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return "student"
		}
	}
	return "staff"
}

func ProcessConfig(filename string, conf interface{}) {
	confdata := ReadFile(filename)
	xconf := conf
	yaml.Unmarshal(confdata, xconf)
	conf = xconf
}
