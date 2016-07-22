package main

import (
     "os"
     "log"
     "fmt"
     "net/http"
     "golang.org/x/net/html"
     "github.com/yhat/scrape"
     "strings"
     "strconv"
     "sort"
)

type CharFreq struct {
    char string
    freqCnt int
}

func (c CharFreq) String() string {
    return fmt.Sprintf("%s: %d", c.char, c.freqCnt)
}

type ByFreqCntDesc []CharFreq

func (a ByFreqCntDesc) Len() int           { return len(a) }
func (a ByFreqCntDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFreqCntDesc) Less(i, j int) bool { return a[i].freqCnt > a[j].freqCnt }

func handler (w http.ResponseWriter, r *http.Request) {
    log.Println("handler func received request: ", r.URL.Path)

    // I could hard-code the strings and avoid getting this page every time...maybe later.
    resp, err := http.Get("http://www.fogcreek.com/jobs/supportengineer")
    if err != nil {
        panic(err)
    }
    root, err := html.Parse(resp.Body)
    if err != nil {
        panic(err)
    }
    // Define var with this scope to be able to get to them later.
    var codeElems []string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "code" {
            // I tried very hard to avoid taking on `scrape` as a dependency.
            // However, reading out the text contents from within an html.Node data type was NOT straightforward for me.
            // The `scrape` package made it easy, and the code w/in the package is only about 150 lines.
            // At that length, it could be undertood with a little study or even pasted into this project if licensing allows.
            fmt.Fprintf(w, scrape.Text(n) + "\n\n")
            codeElems = append(codeElems, scrape.Text(n))
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
            // Assume we only need the 2 <code> tags on the job posting page, don't parse any further.
            if len(codeElems) >= 2 { break }
        }
    }
    f(root)

    // Make some dicey assumptions about scraped elements here:
    // - 1st <code> tag will be the characters to get counts for
    // - 2nd <code> tag will be corpus to run freq count on
    // Define a slice of characters to get counts for, the corpus string, and another slice for counts.
    // I tried a 2-dimensional array to hold both chars to get counts for and the counts,
    //  but kept getting unexpected results when filling the 2nd dimension, so I ditched it.
    // The better way forward here is probably to define a struct to hold letters and counts.
    charsToCnt := strings.Split(codeElems[0], "")
    corpus := codeElems[1]
    var charFreqs []CharFreq

    // Loop over slice w/ chars to get counts for; get a freq count of letter in corpus string;
    // Print results to web page for a sanity check.
    // I was hoping to start a goroutine for each iteration of this loop to get concurrency
    // (and because I've never played with goroutines before), but I quickly got lost in the semantics of channels, semaphores, etc.
    for i, char := range charsToCnt {
        log.Println("corpus cnt: " + string(len(corpus)))
        charFreqs = append(charFreqs, CharFreq{char, strings.Count(corpus, char)})
        fmt.Fprintf(w, char + ", " + strconv.Itoa(charFreqs[i].freqCnt) + "; ")
    }

    fmt.Fprintf(w, "\n\n")

    // Order by freq cnt, descending.
    // drop all chars after (and including) the _ to get the secret word; print to page and stdout
    sort.Sort(ByFreqCntDesc(charFreqs))
    var secretWord string
    for _, charFreq := range charFreqs {
        if charFreq.char == "_" { break }
        fmt.Fprintf(w, charFreq.String())
        secretWord = secretWord + charFreq.char
    }

    fmt.Fprintf(w, "\n\n" + secretWord)
    fmt.Println(secretWord)

    /*
        TODO:
        - write some tests
        - figure out how to adapt setup to run as web handler for debugging and to print to stdout
            - this might get ugly/hacky
            - or not, might be able to call handler() directly if no port set
        - rebuild docker image as jkaplon/fog-creek-supp-eng; push to hub
    */
}

func main () {
    http.HandleFunc("/", handler)
   
    port := os.Getenv("PORT")
    if port == "" {
        log.Fatal("PORT environment variable was not set")
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatal("Could not listen: ", err)
    }
}
