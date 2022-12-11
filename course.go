package teacher

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

func (c *Course) Checkout() {
	fmt.Println("It's time to check what you have done!")
	fmt.Println("--------------------------------------------------------------------------------")
	for i, task := range c.Tasks {
		testCases := task.TestCases()
		taskSeqNum := i + 1

		for j, testCase := range testCases {
			testCaseSeqNum := j + 1
			outputResult(task.Name(), testCase.Name(), testCase.Check(), taskSeqNum, testCaseSeqNum)
		}
	}
	fmt.Println("--------------------------------------------------------------------------------")
}
