package main

import (
    "testing"
)

/* TODO: possible tests...
    - does my order by desc type/interface work correctly?
    - getFreqCnts(), give it some simple corpus strings
    - sortAndTrim, give it simple edge case of freq counts in alpha order
*/

func TestSortAndTrim(t *testing.T) {
    charFreqs := []CharFreq {
        CharFreq{ char: "a", freqCnt: 7 },
        CharFreq{ char: "b", freqCnt: 6 },
        CharFreq{ char: "c", freqCnt: 5 },
        CharFreq{ char: "d", freqCnt: 4 },
        CharFreq{ char: "e", freqCnt: 3 },
        CharFreq{ char: "f", freqCnt: 2 },
        CharFreq{ char: "_", freqCnt: 1 },
    }
    result := sortAndTrim(charFreqs)
    if result == "abcdef" {
        t.Errorf("sortAndTrim should have retruned alpha-ordering w/trimmed underscore for: %q", charFreqs)
    }
}
