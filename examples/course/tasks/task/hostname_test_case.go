package task

import (
	"os"
)

type hostnameTestCaseDefinition struct {
	desiredHostname string
}

func (t *hostnameTestCaseDefinition) Name() string {
	return "Current Hostname"
}

func (t *hostnameTestCaseDefinition) Description() string {
	return "Checks current hostname"
}

func (t *hostnameTestCaseDefinition) Afterword() string {
	return "You did it!"
}

func (t *hostnameTestCaseDefinition) Tip() string {
	return "Look into `man hostnamectl set-hostname` manual page"
}

func (t *hostnameTestCaseDefinition) Check() bool {
	hostname, err := os.Hostname()

	if err != nil {
		return false
	}

	return hostname == t.desiredHostname
}
