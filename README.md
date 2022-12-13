# Course Builder

**Course Builder** helps create a standard practice course checker.

# How to use

See example of creating a course in `examples` directory.

Available commands and flags of application may be found using `--help` flag, e.g.:

```console
$ go build -o course-builder

$ ./course-builder --help

Usage:  ./course-builder COMMAND

Commands:
  course:  Manage a course
  task:    Manage course tasks

$ ./course-builder course --help

Usage: ./course-builder course COMMAND [OPTION] [ARG]

Commands:
  description:  Show course description
  checkout:     Run test cases for all tasks

Options:
  --no-afterword Do not show afterwords
  --tips Show tips if a test failed

$ ./course-builder task --help

Usage: ./course-builder task COMMAND

Commands:
  description:  Show task description
  list:         Show list of available task names
```
