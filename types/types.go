package types

type Action string

type Request struct {
	Patterns     []string
	Exclusions   []string
	Files        []string
	InDirectory  string
	OutDirectory string
}

type Config struct {
	InDirectory  string
	OutDirectory string
	Requests     []Request
	Action       Action
}

type Cancel struct{}

func (e *Cancel) Error() string {
	return "Nothing happen"
}

type Execute struct{}

func (e *Execute) Error() string {
	return ""
}
