package utils4ruts

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	hcl "github.com/hashicorp/hcl/v2/hclsimple"
	random "github.com/ruts48code/random4ruts"
	"gopkg.in/yaml.v3"
)

func ReadFile(f string) []byte {
	content, err := os.ReadFile(f)
	if err != nil {
		log.Printf("Error: utils4ruts-ReadFile - %v\n", err)
		return []byte("")
	}
	return content
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

func ProcessConfigHCL(filename string, conf interface{}) {
	err := hcl.DecodeFile(filename, nil, conf)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
		panic(err)
	}
}

func HTTPGet(url string) (output []byte) {
	output = make([]byte, 0)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr}
	res, err := client.Get(url)
	if err != nil {
		log.Printf("Error: utils4ruts-HTTPGet 1 - Get error - %v\n", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("Error: utils4ruts-HTTPGet 2 - http return code is not ok - %d\n", res.StatusCode)
		return
	}
	output, err = io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error: utils4ruts-HTTPGet 3 - cannot read body - %v\n", err)
		return
	}
	return
}
