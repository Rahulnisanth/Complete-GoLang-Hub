// SENTENCE SIMILARITY III :
package dailystreaks

import (
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
    words1 := strings.Split(sentence1, " ")
    words2 := strings.Split(sentence2, " ")
    if len(words1) > len(words2) {
        words1, words2 = words2, words1
    }
    i := 0
    j := 0
    N1 := len(words1)
    N2 := len(words2)
    for (i < N1) && (i < N2) && (words1[i] == words2[i]) {
        i += 1
    }
    for (j < N1 - i) && (j < N2 - i) && (words1[N1 - j - 1] == words2[N2 - j - 1]) {
        j += 1
    }
    return (i + j) == N1 || (i + j) == N2
}