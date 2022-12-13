package coursebuilder

import (
	"fmt"
)

type Course struct {
	Name        string
	Description string
	Tasks       []Task
}

func NewCourse(name string, description string, tasks []Task) *Course {
	return &Course{
		Name:        name,
		Description: description,
		Tasks:       tasks,
	}
}

func (c *Course) Checkout(noAfterwordFlag bool, tipsFlag bool) {
	fmt.Println("It's time to check what you have done!")
	fmt.Println("--------------------------------------------------------------------------------")
	for i, task := range c.Tasks {
		taskResult := true
		testCases := task.TestCases()
		taskSeqNum := i + 1

		for j, testCase := range testCases {
			testCaseResult := testCase.Check()
			testCaseSeqNum := j + 1
			outputResult(task.Name(), testCase.Name(), testCaseResult, taskSeqNum, testCaseSeqNum)

			if noAfterwordFlag == false && testCaseResult && testCase.Afterword() != "" {
				fmt.Println("Test Case Afterword:\n" + testCase.Afterword())
			}

			if tipsFlag && testCaseResult == false && testCase.Tip() != "" {
				fmt.Println("Test Case Tips:\n" + testCase.Tip())
			}

			if testCaseResult == false {
				taskResult = false
			}
		}

		if noAfterwordFlag == false && taskResult && task.Afterword() != "" {
			fmt.Println("Task Afterword:\n" + task.Afterword())
		}
	}
	fmt.Println("--------------------------------------------------------------------------------")
}
