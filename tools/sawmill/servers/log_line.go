package servers

type LogLine struct {
	Time    string
	Level   string
	Message string
	Source  string
	Object  map[string]any
}
