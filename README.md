# Course Builder

**Course Builder** helps create a standard practice course checker.

# How to use

See example of creating a course in `examples` directory.

Available commands and flags of application may be found using `--help` flag, e.g.:

```console

$ go build -o course-checker

$ ./course-checker --help

Usage:  ./course-checker COMMAND

Commands:
  course:  Manage a course
  task:    Manage course tasks

$ ./course-checker course --help

Usage: ./course-checker course COMMAND [OPTION] [ARG]

Commands:
  description:  Show course description
  check:        Run test cases for all tasks

Options:

$ ./course-checker course check --help

Usage: ./course-checker course check COMMAND [OPTION] [ARG]


Options:
  --afterword Shows afterwords if a test or task passed
  --tips Show tips if a test failed

$ ./course-checker task --help

Usage: ./course-checker task COMMAND

Commands:
  description:  Show task description
  list:         Show list of available task names
