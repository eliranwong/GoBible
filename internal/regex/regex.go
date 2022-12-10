package regex

import (
	"fmt"
	"regexp"
)

// re2 syntax: https://github.com/google/re2/wiki/Syntax
// re2 does not support back reference, read my workaround at https://stackoverflow.com/a/74715496/20712208

// search and replace for general cases
func ReplaceAllString(text, flags string, searchReplace [][2]string) string {
	for _, v := range searchReplace {
		search, replace := v[0], v[1]
		compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, flags, search))
		text = compiledPattern.ReplaceAllString(text, replace)
	}
	return text
}

// search and replace with loop to make sure no match is found
func ReplaceAllStringLoop(text, flags, loopPattern string, searchReplace [][2]string) string {
	compiledPattern := regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, flags, loopPattern))
	for compiledPattern.MatchString(text) {
		text = ReplaceAllString(text, flags, searchReplace)
	}
	return text
}

// supports regular expression search in sqlite files
// remarks: official go-sqlite example https://github.com/mattn/go-sqlite3/tree/master/_example/mod_regexp does NOT work on macOS M1
// we use custom function instead
func Re(pattern, text string) bool {
	return regexp.MustCompile(fmt.Sprintf(`(?%v)%v`, "i", pattern)).MatchString(text)
}
