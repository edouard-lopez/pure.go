package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"
	"time"

	"github.com/carlmjohnson/be"
	"github.com/edouard-lopez/pure.go/internal/constant"
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

	symbol := regexp.MustCompile(constant.PromptSymbol)
	result, _, err := child.Expect(symbol, time.Second)
	if err != nil {
		log.Fatalf("\nExpected: %q\nGot: %q\nError: %v", symbol, result, err)
	}

	expect := fmt.Sprintf("%s%v\n", constant.PromptSymbol, constant.ExitCodeSuccess)
	be.Equal(t, expect, result)
}
