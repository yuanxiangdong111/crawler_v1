package word2

import (
    "unicode"
)

func IsPalindrome(s string) bool {
    var letters []rune
    for _, v := range s {
        if unicode.IsLetter(v) {
            letters = append(letters, unicode.ToLower(v))
        }
    }

    for i := range letters {
        if letters[i] != letters[len(letters)-i-1] {
            return false
        }
    }
    return true
}
