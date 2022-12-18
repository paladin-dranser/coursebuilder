package coursebuilder

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
			fmt.Fprintf(flag.CommandLine.Output(), "  check:        Run test cases for all tasks\n")

			fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
			courseCmd.VisitAll(func(f *flag.Flag) {
				fmt.Fprintf(flag.CommandLine.Output(), "  --"+f.Name+" "+f.Usage+"\n")
			})
		}

		courseCmd.Parse(os.Args[2:])
		checkoutCmd := flag.NewFlagSet("check", flag.ExitOnError)

		switch os.Args[2] {
		case "description":
			fmt.Println(course.Description)
			return
		case "check":
			checkoutCmd.Usage = func() {
				fmt.Fprintf(flag.CommandLine.Output(), "\nUsage: %s %s %s COMMAND [OPTION] [ARG]\n\n", os.Args[0], os.Args[1], os.Args[2])
				fmt.Fprintf(flag.CommandLine.Output(), "\nOptions:\n")
				checkoutCmd.VisitAll(func(f *flag.Flag) {
					fmt.Fprintf(flag.CommandLine.Output(), "  --"+f.Name+" "+f.Usage+"\n")
				})
			}
			afterwordFlag := checkoutCmd.Bool("afterword", false, "Show afterwords if a test or task passed")
			tipsFlag := checkoutCmd.Bool("tips", false, "Show tips if a test failed")
			taskFlag := checkoutCmd.String("task", "", "Check a specific task")
			checkoutCmd.Parse(os.Args[3:])

			course.Check(*afterwordFlag, *tipsFlag, *taskFlag)
		default:
			fmt.Fprintf(flag.CommandLine.Output(), "ERROR: Incorrect command!\n")
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
