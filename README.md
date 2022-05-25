# [CipherAPI](https://cipher-e7737.web.app) &middot; [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/facebook/react/blob/main/LICENSE)   [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://reactjs.org/docs/how-to-contribute.html#your-first-pull-request)

CipherAPI is an open API for encoding and decoding text using post requests.

## Methods

* [Morse](#morse)
* [Usage](#usage)
* [Installation](#installation)
* [API](#API)
  
## Usage

Cipher API can be used three ways:

### Using Post Requests

The simplest way to try it out and the way it's meant to be used. As a public API you can do POST requests:

```bash
curl -X POST https://ciphertapi.herokuapp.com/encode/morse -H Content-Type:application/json -d '{"text":"Hello world"}'
```

```js
const response = await fetch("https://ciphertapi.herokuapp.com/encode/morse", {
    method: "POST",
    headers: { "content-type": "application/json"},
    body: JSON.stringify({
        text: "Hello world"
    }),
})
const data = await response.text();

if (!response.ok) {
    alert(data);
} else {
    console.log(data);
}
```

### Website

Provides a full deployed [website](https://cipher-e7737.web.app) so you can try all methods for encoding! It saves your historical encodings with user sessions :)

### Locally

If you wish try it on your own and modify it freely, please refer to [installation](#installation).

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

Make sure you have [Go](https://go.dev/doc/install) installed.

If you do wish to run the server locally, run this command to clone the project

```bash
git clone https://github.com/cheveuxdelin/cipher.git
```

And then run this command on the cloned project to start the server in your local workspace

```bash
go run main.go
```

And thats it! you can now make requests on port :8080, or if you can change the port if you need to

## API

All methods support both actions encode and decode.

:exclamation: marked as required:

| Method | Parameters |
| --- | --- |
| Morse | text :exclamation: |
| Caesar | text :exclamation:, n :exclamation:, onlyLetters  |
| Atbash | text :exclamation: |

