package main

import (
	"github.com/paladin-dranser/coursebuilder"
	"course/task"
)

func main() {
	hostnameTask := task.NewHostnameTaskDefinition("example")

	tasks := []coursebuilder.Task{
		hostnameTask,
	}

	courseDescription := "Example Course\n\n" +
		"The primary purpose of this course is to show example of using course-builder package."

	course := coursebuilder.NewCourse("Example Course", courseDescription, tasks)

	coursebuilder.Execute(course)
}
