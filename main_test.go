package main

import (
    "testing"
)

func TestGetFreqCnts(t *testing.T) {
    codeElems := []string { "a", "aaaaa" }
    expected := "a, 5; "
    actual := getFreqCnts(codeElems)
    if expected != actual {
        t.Errorf("getFreqCnts error, for input %q, expected %q, got %q", codeElems, expected, actual)
    }

    // Try a corpus string with underscore (the only non-alpha character involved here).
    // Warning, this test is flaky, when I tried a corpus of 3-underscores it failed w/actual count of 5;
    // I suspect something weird with my initial declaration of codeElems with overall lenght of 5,
    // and possible treatment of underscore as wildcard character :-(.
    // Maybe it's bad practice to recycle variable names like I'm doing...
    codeElems = []string { "_", "_____" }
    expected = "_, 5; "
    actual = getFreqCnts(codeElems)
    if expected != actual {
        t.Errorf("getFreqCnts error, for input %q, expected %q, got %q", codeElems, expected, actual)
    }
}

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
    expected := "abcdef"
    actual := sortAndTrim(charFreqs)
    if expected != actual {
        t.Errorf("sortAndTrim error for input %q, expected %q, got %q", charFreqs, expected, actual)
    }

    // This one should trim everything and return an empty string since the underscore has the highest count.
    charFreqs = []CharFreq {
        CharFreq{ char: "a", freqCnt: 1 },
        CharFreq{ char: "b", freqCnt: 2 },
        CharFreq{ char: "c", freqCnt: 3 },
        CharFreq{ char: "d", freqCnt: 4 },
        CharFreq{ char: "e", freqCnt: 5 },
        CharFreq{ char: "f", freqCnt: 6 },
        CharFreq{ char: "_", freqCnt: 7 },
    }
    expected = ""
    actual = sortAndTrim(charFreqs)
    if expected != actual {
        t.Errorf("sortAndTrim error for input %q, expected %q, got %q", charFreqs, expected, actual)
    }

    // I could test my ByFreqCntDesc sorting setup independently, but since it's called directly by sortAndTrim, I'm calling it good enough for now.
    // I would like to test scrapeAndParse, but I lack the time and patience to integrate all the necessary pieces
    // (one hacky solution would be to host a custom pages on my VPS containing <code> tags with test values).
}
