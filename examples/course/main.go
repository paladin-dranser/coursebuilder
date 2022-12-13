package main

import (
	"github.com/paladin-dranser/course_builder"
	"course/task"
)

func main() {
	hostnameTask := task.NewHostnameTaskDefinition("example")

	tasks := []course_builder.Task{
		hostnameTask,
	}

	courseDescription := "Example Course\n\n" +
		"The primary purpose of this course is to show example of using course-builder package."

	course := course_builder.NewCourse("Example Course", courseDescription, tasks)

	course_builder.Execute(course)
}
