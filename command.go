package teacher

import (
	"flag"
	"fmt"
	"os"
)

func Execute(course *Course) {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage:  %s COMMAND\n\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Commands:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  course:  Manage a course\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  task:    Manage course tasks\n")
	}
	flag.Parse()

	taskCmd := flag.NewFlagSet("task", flag.ExitOnError)
	listFlag := taskCmd.Bool("list", false, "Show list of available tasks")
	taskNameFlag := taskCmd.String("name", "", "Specify task name")
	taskDescriptionFlag := taskCmd.Bool("description", false, "Show description of task which name is specified by '--name' flag")

	courseCmd := flag.NewFlagSet("course", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Fprintf(flag.CommandLine.Output(), "Please specify a subcommand!\n")
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "course":
		courseCmd.Parse(os.Args[2:])

		args := courseCmd.Args()

		if len(args) < 1 {
			fmt.Println("ERROR")
			os.Exit(1)
		}

		if args[0] == "description" {
			fmt.Println(course.Description)
			return
		} else if args[0] == "checkout" {
			course.Checkout()
			return 
		} else {
			fmt.Println("ERROR")
			os.Exit(1)
		}
	case "task":
		taskCmd.Parse(os.Args[2:])
	default:
		fmt.Println("ERROR: expected 'task' subcommand")
		os.Exit(1)
	}

	if *taskDescriptionFlag {
		if *taskNameFlag == "" {
			fmt.Println("ERROR: task name must be defined!")
			os.Exit(1)
		}

		for _, task := range course.Tasks {
			if task.Name() == *taskNameFlag {
				fmt.Println("Description: " + task.Name() + "\n" + task.Description())
				return
			}
		}

		fmt.Println("ERROR: Task has not been found!")
		os.Exit(1)
	}

	if *listFlag {
		for _, task := range course.Tasks {
			fmt.Println(task.Name())
		}
		return
	}
}
