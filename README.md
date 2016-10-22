# str-freq-sort
One-off golang project to solve a coding challenge:
- Take the sample string posted on a web page
- Get frequency counts of characters
- Order characters by frequency
- And (hopefully) reveal the secret word

I chose golang for the project with the hopes of learning more about concurrency with goroutines and channels.
That didn't work out.
Instead, I ended up getting experience with golang's [interfaces](https://gobyexample.com/interfaces), and a couple new packages:
- [scrape](github.com/yhat/scrape)
- [net/html](golang.org/x/net/html)

This repo also contains some basic unit tests that can be run with `go test`. 

My results are published on [my Docker Hub account](https://hub.docker.com/r/jkaplon/str-freq-sort/).
Give it a pull with `docker pull jkaplon/str-freq-sort`.
Give it a try with `docker run --rm --name str-freq-sort jkaplon/str-freq-sort`.

In the process, I built a hacky-but-useful development docker image with auto-reload for Go web apps.
The development branch contains the alternate Dockerfile with the auto-reload setup.
