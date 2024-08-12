package defang

import (
	"regexp"
)

// DefangURL defangs a given URL by replacing . with [.] and : with [:]
func DefangURL(url string) string {
	regex := regexp.MustCompile(`(\.|\:)`)
	defanged := regex.ReplaceAllStringFunc(url, func(s string) string {
        if s == "." {
            return "[.]"
        } else if s == ":" {
            return "[:]"
        }
        return s
    })
    return defanged
}