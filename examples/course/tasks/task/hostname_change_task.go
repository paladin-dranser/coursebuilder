package task

import (
	"github.com/paladin-dranser/course_builder"
)

type hostnameTaskDefinition struct {
	desiredHostname string
	testCases       []course_builder.TestCase
}

func (t *hostnameTaskDefinition) Name() string {
	return "Hostname Change"
}

func (t *hostnameTaskDefinition) Description() string {
	return "Change hostname of your virtual machine to '" + t.desiredHostname + "'\n"
}

func (t *hostnameTaskDefinition) Afterword() string {
	return "You did it!"
}

func (t *hostnameTaskDefinition) TestCases() []course_builder.TestCase {
	return t.testCases
}

func NewHostnameTaskDefinition(hostname string) *hostnameTaskDefinition {
	t := hostnameTaskDefinition{
		desiredHostname: hostname,
	}

	hostnameTestCase := &hostnameTestCaseDefinition{
		desiredHostname: t.desiredHostname,
	}
	t.testCases = append(t.testCases, hostnameTestCase)

	return &t
}
