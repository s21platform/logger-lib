package logger_lib

type Entry struct {
	Streams []StreamEntry `json:"streams"`
}

type StreamEntry struct {
	Stream Stream     `json:"stream"`
	Values [][]string `json:"values"`
}

type Stream struct {
	Service      string `json:"service"`
	Level        string `json:"level"`
	Environment  string `json:"environment"`
	FunctionName string `json:"function_name,omitempty"`
}
