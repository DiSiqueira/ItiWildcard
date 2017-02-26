# ItiWildcard ![Language Badge](https://img.shields.io/badge/Language-Go-blue.svg) ![Go Report](https://goreportcard.com/badge/github.com/DiSiqueira/ItiWildcard) ![License Badge](https://img.shields.io/badge/License-MIT-blue.svg) ![Status Badge](https://img.shields.io/badge/Status-Stable-brightgreen.svg)

ItiWildcard is a matcher to be used in the [Itinerary][itinerary].

The ItiWildcard's goal is to be a simple way to handle routes in Itinerary, using the `*` wildcard. Using the wild card will allow you to create dynamic routes.

[itinerary]: https://github.com/DiSiqueira/Itinerary

## Project Status

ItiWildcard is stable. Pull Requests [are welcome](https://github.com/DiSiqueira/ItiWildcard#social-coding)

## Installation

### Go Get

```bash
$ go get github.com/DiSiqueira/itiwildcard
```

## Basic usage

```go
r := itinerary.NewRouter()

r.NewPath(RateHandler).AddMatcher(ItiWildcard.New("/movies/*/rate"))
```

## Full Example

```go
package main

import (
	"github.com/disiqueira/itinerary"
	"net/http"
	"github.com/disiqueira/itiwildcard"
)

func RateHandler(w http.ResponseWriter, r *http.Request) {
	movie := r.URL.Query().Get("itinerary1")
	msg := "Rate "+movie
	w.Write([]byte(msg))
}

func main() {
	r := itinerary.NewRouter()

	r.NewPath(RateHandler).AddMatcher(itiwildcard.New("/movies/*/rate"))

	http.ListenAndServe(":8000", r)
}
```

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/DiSiqueira/ItiWildcard/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. To begin developing, do this:

```bash
$ git clone https://github.com/DiSiqueira/ItiWildcard.git itiwildcard
$ cd itiwildcard/
```

## Social Coding

1. Create an issue to discuss about your idea
2. [Fork it] (https://github.com/DiSiqueira/ItiWildcard/fork)
3. Create your feature branch (`git checkout -b my-new-feature`)
4. Commit your changes (`git commit -am 'Add some feature'`)
5. Push to the branch (`git push origin my-new-feature`)
6. Create a new Pull Request
7. Profit! :white_check_mark:

## License

The MIT License (MIT)

Copyright (c) 2013-2017 Diego Siqueira

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
