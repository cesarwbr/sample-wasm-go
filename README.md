# sample-wasm-go

Implementation of a WebAssembly with Go. The Go function is in charge to handle the image file received from the JavaScript. We will use a package to find the dominant color within an image.

The package called prominentcolor uses the Kmeans++ algorithm to work this out. By default, it will return us the three most popular colors. We can install this package using the follow command:

```sh
$ go get github.com/EdlinOrg/prominentcolor
```

## Generate the Wasm JavaScript Loader

We have to generate the wasm_exec.js file. It's a JavaScript file provided by Go to load your .wasm file into a Web page.

The loader file is available in your standard to Go installation. Just copy it into the project folder with the following command:

```sh
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
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
