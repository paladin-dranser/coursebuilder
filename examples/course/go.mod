module example

go 1.19

replace github.com/paladin-dranser/course_builder => ../../

replace course/task => ./tasks/task

require (
	course/task v0.0.0-00010101000000-000000000000
	github.com/paladin-dranser/course_builder v0.0.0-00010101000000-000000000000
)
