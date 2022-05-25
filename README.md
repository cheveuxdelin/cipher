# [CipherAPI](https://cipher-e7737.web.app) &middot; [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/facebook/react/blob/main/LICENSE)   [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://reactjs.org/docs/how-to-contribute.html#your-first-pull-request)

CipherAPI is an open API for encoding and decoding text using post requests.

## Methods

* [Morse](#morse)
* [Usage](#usage)
* [Installation](#installation)
  
## Usage



```go
package main

import (
 "github.com/cheveuxdelin/cipher/atbash"
 "github.com/cheveuxdelin/cipher/caesar"
 "github.com/cheveuxdelin/cipher/morse"
)

func main() {
        //Using different methods for encoding/decoding
        encoded, err := morse.Encode("Hello world")

        if err != nil {
            fmt.Println(err)
        }

        decoded, err := morse.Decode("Hello world")

        if err != nil {
            fmt.Println(err)
        }

        fmt.Println(encoded, decoded)
}
```

## Installation

If you do wish to run the server locally, run this command to clone the project

```bash
git clone https://github.com/cheveuxdelin/cipher.git
```

And then run this command on the cloned project to start the server in your local workspace

```bash
go run main.go
```
