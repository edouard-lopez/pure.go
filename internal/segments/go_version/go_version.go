package go_version

import (
	"fmt"
	"os/exec"
	"regexp"
)

// var GetGoVersion = runtime.Version // so we can mock it
var GetGoVersion = exec.Command("go", "version").Output // so we can mock it

func Parse(output string) string {
	regex := regexp.MustCompile(`go([^\s]+)`)

	match := regex.FindStringSubmatch(output)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func Get() string {
	// check command exists before running it
	_, err := exec.LookPath("go")
	if err != nil {
		return ""
	}

	output, err := GetGoVersion()
	if err != nil {
		fmt.Println("go missing")
	}
	return Parse(string(output))
}
