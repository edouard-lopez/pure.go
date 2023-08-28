package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/edouard-lopez/pure.go/internal/constants"
	expect "github.com/google/goexpect"
)

func getPureBinPath() string {
	testDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	purePath := filepath.Join(filepath.Dir(testDir), "pure")
	fmt.Println(purePath)

	return purePath
}

func TestCLIE2E(t *testing.T) {
	child, _, err := expect.Spawn(getPureBinPath(), -1)
	if err != nil {
		log.Fatal(err)
	}
	defer child.Close()

	symbol := regexp.MustCompile("(?m)" + constants.PromptSymbol)
	actual, _, err := child.Expect(symbol, time.Second)
	if err != nil {
		log.Fatalf("\nExpected: %q\nGot: %q\nError: %v", symbol, actual, err)
	}

	// expected := fmt.Sprintf("%s%v\n", constants.PromptSymbol, constants.ExitCodeSuccess)
	// assert.Equal(t, expected, actual)
}
