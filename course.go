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

type UnknownTaskError string

func (e UnknownTaskError) Error() string {
	return "task: unknown task " + string(e)
}

// LookupTask looks up a task by name.
// if the task cannot be found, the returned error is of type UnknownTaskError
func (c *Course) LookupTask(name string) (*Task, error) {
	for _, task := range c.Tasks {
		if name == task.Name() {
			return &task, nil
		}
	}

	return nil, UnknownTaskError(name)
}

func (c *Course) Check(afterwordFlag bool, tipsFlag bool, taskFlag string) {
	tasks := c.Tasks

	if taskFlag != "" {
		task, err := c.LookupTask(taskFlag)

		if err != nil {
			fmt.Println(err)
			return
		}

		tasks = []Task{*task}
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
