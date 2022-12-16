package coursebuilder

import (
	"fmt"
	"strconv"
)

func outputResult(taskName string, testCaseName string, testCaseResult bool, taskSeqNum int, testCaseSeqNum int) {
	const colorReset = "\033[0m"
	const colorRed = "\033[31m"
	const colorGreen = "\033[32m"

	seqNum := strconv.Itoa(taskSeqNum) + "." + strconv.Itoa(testCaseSeqNum)
	if testCaseResult {
		fmt.Printf("%s✓ %s %s: %s%s\n", colorGreen, seqNum, taskName, testCaseName, colorReset)
	} else {
		fmt.Printf("%s✗ %s %s: %s%s\n", colorRed, seqNum, taskName, testCaseName, colorReset)
	}
}

type Task interface {
	Name() string
	Description() string
	TestCases() []TestCase
}

type TestCase interface {
	Name() string
	Description() string
	// TODO Return err to use it in debug mode
	Check() bool
}

type Afterword interface {
	Afterword() string
}

type Tip interface {
	Tip() string
}
