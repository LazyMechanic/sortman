package types

type Action string

type Request struct {
	Patterns     []string
	Exclude      []string
	Files        []string
	InDirectory  string
	OutDirectory string
}

type Config struct {
	WorkingDirectory string
	Requests         []Request
	Action           Action
}

type Cancel struct{}

func (e *Cancel) Error() string {
	return "Nothing happen"
}

type Execute struct{}

func (e *Execute) Error() string {
	return ""
}
