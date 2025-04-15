package axo

import (
	"math/rand"
	"regexp"
	"strings"
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

func GenerateMemCode(blockCount int) string {
	// Genişletilmiş kelime listesi
	words := []string{
		"blue", "cake", "door", "fish", "gold", "home", "jump", "king",
		"lamp", "moon", "note", "park", "queen", "road", "star", "tree",
		"wind", "book", "duck", "frog", "goat", "hill", "kite", "lion",
		"milk", "nest", "owl", "pear", "rain", "song", "time", "wave",
		"bird", "card", "cool", "dark", "echo", "fire", "glow", "hope",
		"iris", "jade", "lava", "mint", "nova", "opal", "pink", "quiz",
		"rice", "snow", "tone", "undo", "view", "wood", "xray", "yarn",
		"zinc", "acme", "brio", "clay", "dune", "envy", "flux", "grin",
	}

	// Blok sayısı kelime sayısından fazla olamaz
	if blockCount > len(words) {
		blockCount = len(words)
	}

	used := make(map[string]bool)
	var parts []string

	for len(parts) < blockCount {
		index := rand.Intn(len(words))
		word := words[index]
		if used[word] {
			continue
		}
		used[word] = true
		if len(word) > 4 {
			word = word[:4]
		}
		parts = append(parts, word)
	}

	return strings.Join(parts, "-")
}
