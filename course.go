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

func (c *Course) Check(afterwordFlag bool, tipsFlag bool, taskFlag string) {
	tasks := c.Tasks

	if taskFlag != "" {
		found := false
		for _, task := range tasks {
			if taskFlag == task.Name() {
				tasks = []Task{task}
				found = true
				break
			}
		}

		if found == false {
			fmt.Println("'" + taskFlag + "' task has not been found in the course. Exiting...")
			return
		}
	}

	fmt.Println("It's time to check what you have done!")
	fmt.Println("--------------------------------------------------------------------------------")

	for i, task := range tasks {
		taskResult := true
		testCases := task.TestCases()
		taskSeqNum := i + 1

		for j, testCase := range testCases {
			testCaseResult := testCase.Check()
			testCaseSeqNum := j + 1
			outputResult(task.Name(), testCase.Name(), testCaseResult, taskSeqNum, testCaseSeqNum)

			if tca, ok := testCase.(Afterword); ok {
				if afterwordFlag && testCaseResult {
					fmt.Println("Test Case Afterword:\n" + tca.Afterword())
				}
			}

			if tct, ok := testCase.(Tip); ok {
				if tipsFlag && testCaseResult == false {
					fmt.Println("Test Case Tips:\n" + tct.Tip())
				}
			}

			if testCaseResult == false {
				taskResult = false
			}
		}

		if ta, ok := task.(Afterword); ok {
			if afterwordFlag && taskResult {
				fmt.Println("Task Afterword:\n" + ta.Afterword())
			}
		}
	}
	fmt.Println("--------------------------------------------------------------------------------")
}
