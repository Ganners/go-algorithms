package lda

import (
	"math/rand"
	"regexp"
	"strings"
)

func VecSum(a []float64) float64 {
	sum := 0.
	for _, v := range a {
		sum += v
	}
	return sum
}

func VecAdd(a []float64, b float64) {
	for i := range a {
		a[i] += b
	}
}

func VecDiv(a []float64, b float64) {
	if b == 0 {
		return
	}
	for i := range a {
		a[i] /= b
	}
}

func VecMul(a []float64, b []float64) []float64 {
	multiplied := make([]float64, len(a))
	for i := range a {
		multiplied[i] = a[i] * b[i]
	}
	return multiplied
}

func ArgMax(a []float64) int {
	maxV := 0.0
	maxI := 0
	for i, v := range a {
		if v > maxV {
			maxI = i
			maxV = v
		}
	}
	return maxI
}

func SampleArg(a []float64) int {
	point := rand.Float64() * VecSum(a)
	cumulative := 0.0
	for i, v := range a {
		cumulative += v
		if point < cumulative {
			return i
		}
	}
	return len(a) - 1
}

var stopwords = map[string]struct{}{
	"a":          struct{}{},
	"about":      struct{}{},
	"above":      struct{}{},
	"after":      struct{}{},
	"again":      struct{}{},
	"against":    struct{}{},
	"all":        struct{}{},
	"am":         struct{}{},
	"an":         struct{}{},
	"and":        struct{}{},
	"any":        struct{}{},
	"are":        struct{}{},
	"aren't":     struct{}{},
	"as":         struct{}{},
	"at":         struct{}{},
	"be":         struct{}{},
	"because":    struct{}{},
	"been":       struct{}{},
	"before":     struct{}{},
	"being":      struct{}{},
	"below":      struct{}{},
	"between":    struct{}{},
	"both":       struct{}{},
	"but":        struct{}{},
	"by":         struct{}{},
	"can't":      struct{}{},
	"cannot":     struct{}{},
	"could":      struct{}{},
	"couldn't":   struct{}{},
	"did":        struct{}{},
	"didn't":     struct{}{},
	"do":         struct{}{},
	"does":       struct{}{},
	"doesn't":    struct{}{},
	"doing":      struct{}{},
	"don't":      struct{}{},
	"down":       struct{}{},
	"during":     struct{}{},
	"each":       struct{}{},
	"few":        struct{}{},
	"for":        struct{}{},
	"from":       struct{}{},
	"further":    struct{}{},
	"had":        struct{}{},
	"hadn't":     struct{}{},
	"has":        struct{}{},
	"hasn't":     struct{}{},
	"have":       struct{}{},
	"haven't":    struct{}{},
	"having":     struct{}{},
	"he":         struct{}{},
	"he'd":       struct{}{},
	"he'll":      struct{}{},
	"he's":       struct{}{},
	"her":        struct{}{},
	"here":       struct{}{},
	"here's":     struct{}{},
	"hers":       struct{}{},
	"herself":    struct{}{},
	"him":        struct{}{},
	"himself":    struct{}{},
	"his":        struct{}{},
	"how":        struct{}{},
	"how's":      struct{}{},
	"i":          struct{}{},
	"i'd":        struct{}{},
	"i'll":       struct{}{},
	"i'm":        struct{}{},
	"i've":       struct{}{},
	"if":         struct{}{},
	"in":         struct{}{},
	"into":       struct{}{},
	"is":         struct{}{},
	"isn't":      struct{}{},
	"it":         struct{}{},
	"it's":       struct{}{},
	"its":        struct{}{},
	"itself":     struct{}{},
	"let's":      struct{}{},
	"me":         struct{}{},
	"more":       struct{}{},
	"most":       struct{}{},
	"mustn't":    struct{}{},
	"my":         struct{}{},
	"myself":     struct{}{},
	"no":         struct{}{},
	"nor":        struct{}{},
	"not":        struct{}{},
	"of":         struct{}{},
	"off":        struct{}{},
	"on":         struct{}{},
	"once":       struct{}{},
	"only":       struct{}{},
	"or":         struct{}{},
	"other":      struct{}{},
	"ought":      struct{}{},
	"our":        struct{}{},
	"ours":       struct{}{},
	"ourselves":  struct{}{},
	"out":        struct{}{},
	"over":       struct{}{},
	"own":        struct{}{},
	"same":       struct{}{},
	"shan't":     struct{}{},
	"she":        struct{}{},
	"she'd":      struct{}{},
	"she'll":     struct{}{},
	"she's":      struct{}{},
	"should":     struct{}{},
	"shouldn't":  struct{}{},
	"so":         struct{}{},
	"some":       struct{}{},
	"such":       struct{}{},
	"than":       struct{}{},
	"that":       struct{}{},
	"that's":     struct{}{},
	"the":        struct{}{},
	"their":      struct{}{},
	"theirs":     struct{}{},
	"them":       struct{}{},
	"themselves": struct{}{},
	"then":       struct{}{},
	"there":      struct{}{},
	"there's":    struct{}{},
	"these":      struct{}{},
	"they":       struct{}{},
	"they'd":     struct{}{},
	"they'll":    struct{}{},
	"they're":    struct{}{},
	"they've":    struct{}{},
	"this":       struct{}{},
	"those":      struct{}{},
	"through":    struct{}{},
	"to":         struct{}{},
	"too":        struct{}{},
	"under":      struct{}{},
	"until":      struct{}{},
	"up":         struct{}{},
	"very":       struct{}{},
	"was":        struct{}{},
	"wasn't":     struct{}{},
	"we":         struct{}{},
	"we'd":       struct{}{},
	"we'll":      struct{}{},
	"we're":      struct{}{},
	"we've":      struct{}{},
	"were":       struct{}{},
	"weren't":    struct{}{},
	"what":       struct{}{},
	"what's":     struct{}{},
	"when":       struct{}{},
	"when's":     struct{}{},
	"where":      struct{}{},
	"where's":    struct{}{},
	"which":      struct{}{},
	"while":      struct{}{},
	"who":        struct{}{},
	"who's":      struct{}{},
	"whom":       struct{}{},
	"why":        struct{}{},
	"why's":      struct{}{},
	"with":       struct{}{},
	"won't":      struct{}{},
	"would":      struct{}{},
	"wouldn't":   struct{}{},
	"you":        struct{}{},
	"you'd":      struct{}{},
	"you'll":     struct{}{},
	"you're":     struct{}{},
	"you've":     struct{}{},
	"your":       struct{}{},
	"yours":      struct{}{},
	"yourself":   struct{}{},
	"yourselves": struct{}{},
}

func wordsFromDocument(document string) []string {
	document = document
	document = strings.ToLower(document)
	document = stripPunctuation(document)
	wordsRaw := strings.Split(document, " ")
	words := make([]string, 0, len(wordsRaw))
	for _, word := range wordsRaw {
		if word == "" || word == " " {
			continue
		}
		if len(word) <= 3 {
			continue
		}
		if _, ok := stopwords[word]; ok {
			continue
		}
		word = strings.TrimSpace(word)
		words = append(words, word)
	}

	return words
}

func stripPunctuation(in string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(in, " ")
}
