# sample-wasm-go

Implementation of a WebAssembly with Go. The Go function is in charge to handle the image file received from the JavaScript. We will use a package to find the dominant color within an image.

The package called prominentcolor uses the Kmeans++ algorithm to work this out. By default, it will return us the three most popular colors. We can install this package using the follow command:

```sh
$ go get github.com/EdlinOrg/prominentcolor
```

## Compile

To compile our main.go file for the web we have to set some environments variables: GOOS and GOARCH. GOOS which stands for Go Operating System, we will set to js. GOARCH which stands for Go Architecture, we will set to wasm. And then we can compile our file using build:

```sh
$ GOOS=js GOARCH=wasm go build -o main.wasm
```

## Run

We can do that in many ways. We can use the Python [SimpleHTTPServer](https://docs.python.org/2/library/simplehttpserver.html), the node package [http-server](https://github.com/http-party/http-server), or use the [goexec](https://github.com/shurcooL/goexec):

```sh
$ goexec 'http.ListenAndServe(`:8080`,http.FileServer(http.Dir(`.`)))'
```

