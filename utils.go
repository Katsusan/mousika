package mousika

import "strings"

// getParams split parameters into
func getParams(path string) (params []string) {
	if path == "/" {
		return
	}

	segments := strings.Split(path, "/")
	replacer := strings.NewReplacer(":", "", "?", "")
	for _, seg := range segments {
		if seg == "" {
			continue
		} else if seg[0] == ':' {
			params = append(params, replacer.Replace(seg))
		}
	}
	return
}
