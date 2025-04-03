package axo

import (
	"regexp"
)

// RegexTest checks if the given text matches the provided regex pattern.
// It returns a boolean indicating if there's a match and any error that occurred.
func RegexTest(text, pattern string) (bool, error) {
	// Compile the regular expression pattern
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}

	// Check if the text matches the compiled regex pattern
	return re.MatchString(text), nil
}

// MultiReplace
func MultiReplace(target string, replacements map[string]string) string {
	for old, new := range replacements {
		target = regexp.MustCompile(old).ReplaceAllString(target, new)
	}
	return target
}
