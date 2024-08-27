package word2

import (
    "testing"
)

func TestPalindrome(t *testing.T) {
    var tests = []struct {
        input string
        want  bool
    }{
        {"", true},
        {"a", true},
        {"aa", true},
        {"ab", false},
        {"kayak", true},
        {"detartrated", true},
        {"A man, a plan, a canal: Panama", true},
        {"Evil I did dwell; lewd did I live.", true},
        {"Able was I ere I saw Elba", true},
        {"été", true},
        {"Et se resservir, ivresse reste.", true},
        {"palindrome", false}, // non-palindrome
        {"desserts", false},   // semi-palindrome
    }

    for _, v := range tests {
        if got := IsPalindrome(v.input); got != v.want {
            t.Errorf(`IsPalindrome(%q) = %v`, v.input, got)
        }
    }
    // 快速排序
}
