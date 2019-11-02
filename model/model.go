package model

type Request struct {
	Headers map[string][]string
	Params  map[string][]string
	Method  string
	Url     string
}
