package compute

type Query struct {
	command   string
	arguments []string
}

func NewQuery(command string, arguments []string) Query {
	return Query{
		command:   command,
		arguments: arguments,
	}
}

func (q *Query) Command() string {
	return q.command
}

func (q *Query) Arguments() []string {
	return q.arguments
}
