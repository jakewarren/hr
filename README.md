# hr
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/hr/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/hr)](https://goreportcard.com/report/github.com/jakewarren/hr)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=shields)](http://makeapullrequest.com)

> generate horizontal rule (hr) for terminals

## Install

```
go get github.com/jakewarren/hr/...
```

## Usage
### As a CLI
![screenshot](images/demo.jpg)
### As a library

```golang
package main

import (
	"fmt"

	"github.com/jakewarren/hr"
)

func main() {
	fmt.Println(hr.HorizontalRule("-"))
}

```

## Acknowledgements

Heavily based on [sashka/hr](https://github.com/sashka/hr).  
Format specification adapted from [LuRsT/hr](https://github.com/LuRsT/hr).  

## License

MIT © 2019 Jake Warren
