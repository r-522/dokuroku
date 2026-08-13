package urlutil

import (
	"net/url"
	"regexp"
	"strings"
)

var urlPattern = regexp.MustCompile(`https?://[^\s<>()"']+`)

func Extract(input string) []string {
	matches := urlPattern.FindAllString(input, -1)
	seen := map[string]bool{}
	urls := make([]string, 0, len(matches))
	for _, match := range matches {
		candidate := strings.TrimRight(match, ".,;:、。）」』]}")
		parsed, err := url.ParseRequestURI(candidate)
		if err != nil || parsed.Scheme == "" || parsed.Host == "" {
			continue
		}
		if !seen[candidate] {
			seen[candidate] = true
			urls = append(urls, candidate)
		}
	}
	return urls
}
