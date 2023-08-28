package current_working_dir

import (
	"log"
	"os"
	"strings"

	"github.com/edouard-lopez/pure.go/internal/colorize"
)

func ReplaceByTilde(path string) string {
	return strings.Replace(path, "/home/"+os.Getenv("USER"), "~", 1)
}

/**
 * Get the current working directory
 * @return {string} The current working directory
 */
func Get() string {
	current_working_dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cwd := ReplaceByTilde(current_working_dir)

	return colorize.Mute(cwd)
}
