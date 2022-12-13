package course_builder

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
	courseCmd := flag.NewFlagSet("course", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Fprintf(flag.CommandLine.Output(), "Please specify a command!\n")
		flag.Usage()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "course":
		courseCmd.Usage = func() {
			fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s %s COMMAND [OPTION] [ARG]\n\n", os.Args[0], os.Args[1])
			fmt.Fprintf(flag.CommandLine.Output(), "Commands:\n")
			fmt.Fprintf(flag.CommandLine.Output(), "  description:  Show course description\n")
			fmt.Fprintf(flag.CommandLine.Output(), "  checkout:     Run test cases for all tasks\n")

			fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
			courseCmd.VisitAll(func(f *flag.Flag) {
				fmt.Fprintf(flag.CommandLine.Output(), "  --" + f.Name + " " + f.Usage + "\n")
			})
		}
		// TODO Add description to FlagSet Usage()
		// TODO Implement 'course' as a FlagSet and use this flag there
		noAfterwordFlag := courseCmd.Bool("no-afterword", false, "Do not show afterwords")
		courseCmd.Parse(os.Args[2:])

		args := courseCmd.Args()

		if len(args) < 1 {
			fmt.Fprintf(flag.CommandLine.Output(), "Please specify a command!\n")
			courseCmd.Usage()
			os.Exit(1)
		}

		if args[0] == "description" {
			fmt.Println(course.Description)
			return
		} else if args[0] == "checkout" {
			course.Checkout(*noAfterwordFlag)
			return
		} else {
			fmt.Fprintf(flag.CommandLine.Output(), "ERROR: Incorrect command. Please specify an available command!\n")
			courseCmd.Usage()
			os.Exit(1)
		}
	case "task":
		taskCmd.Usage = func() {
			fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s %s COMMAND\n\n", os.Args[0], os.Args[1])
			fmt.Fprintf(flag.CommandLine.Output(), "Commands:\n")
			fmt.Fprintf(flag.CommandLine.Output(), "  description:  Show task description\n")
			fmt.Fprintf(flag.CommandLine.Output(), "  list:         Show list of available task names\n")
		}
		taskCmd.Parse(os.Args[2:])

		args := taskCmd.Args()

		if len(args) < 1 {
			fmt.Fprintf(flag.CommandLine.Output(), "Please specify a command!\n")
			taskCmd.Usage()
			os.Exit(1)
		}

		if args[0] == "list" {
			for _, task := range course.Tasks {
				fmt.Println(task.Name())
			}
			return
			// TODO 'description' as 'flagSet' or 'task name' as a flag
		} else if args[0] == "description" {
			if len(args) == 1 {
				fmt.Println("ERROR: Task name is not provided!")
				os.Exit(1)
			}

			for _, task := range course.Tasks {
				if task.Name() == args[1] {
					fmt.Println("Task: " + task.Name() + "\n\n" + task.Description())
					return
				}
			}

			fmt.Println("ERROR: '" + args[1] + "' task has not been found!")

		}
	default:
		fmt.Fprintf(flag.CommandLine.Output(), "ERROR: Incorrect command!\n")
		flag.Usage()
		os.Exit(1)
	}
}
