package compute

import "fmt"

const (
	SetCommand = "SET"
	GetCommand = "GET"
	DelCommand = "DEL"
)

const (
	SetCommandArgumetnsCnt = 2
	GetCommandArgumantsCnt = 1
	DelCommandArgumentsCnt = 1
)

func AnalyzeQuery(tokens []string) (Query, error) {
	command := tokens[0]
	arguments := tokens[1:]

	switch command {
	case SetCommand:
		return createSetQuery(arguments)
	case GetCommand:
		return createGetQuery(arguments)
	case DelCommand:
		return createDelQuery(arguments)
	default:
		return Query{}, fmt.Errorf("invalid command name")
	}
}

func createSetQuery(arguments []string) (Query, error) {
	if len(arguments) != SetCommandArgumetnsCnt {
		return Query{}, fmt.Errorf("invalid number command argument")
	}

	return NewQuery(SetCommand, arguments), nil
}

func createGetQuery(arguments []string) (Query, error) {
	if len(arguments) != GetCommandArgumantsCnt {
		return Query{}, fmt.Errorf("invalid number command argument")
	}

	return NewQuery(GetCommand, arguments), nil
}

func createDelQuery(arguments []string) (Query, error) {
	if len(arguments) != DelCommandArgumentsCnt {
		return Query{}, fmt.Errorf("invalid number command argument")
	}

	return NewQuery(DelCommand, arguments), nil
}
