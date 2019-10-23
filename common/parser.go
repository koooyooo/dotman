package common

import "strings"

func ParseHeader(headers string) map[string][]string {
	if strings.TrimSpace(headers) == "" {
		return map[string][]string{}
	}
	headerMap := map[string][]string{}
	elements := strings.Split(headers, ",")
	for _, e := range elements {
		kv := strings.Split(e, ":")
		k := strings.TrimSpace(kv[0])
		v := strings.TrimSpace(kv[1])
		headerMap[k] = []string{v}
	}
	return headerMap
}
