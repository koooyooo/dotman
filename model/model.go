package model

type Request struct {
	Headers map[string][]string
	Params  map[string][]string
	Method  string
	Url     string
}

type Config struct {
	VerboseResponse bool
	VerboseTime     bool
}
