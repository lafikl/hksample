# HKSample
An algorithm to sample a complete session by hashing its key.


## Installation

     go get github.com/lafikl/hksample


## Example
```go
package main

import (
    "log"
    "github.com/lafikl/hksample"
)
func main() {
    hk := NewHKSample(.20)
    key := []byte(`5efe2993-6b3d-48e2-ac94-1afd543d9190`)
    if hk.Sample(key) {
        log.Printf("Sampled: %s", key)
    }
}
```

## Key Features:

- Guarantees that if an event is chosen it'll always pick it in the following messages, unless the sampling percentage was reduced.
- No extra metadata to add in every message you send, nor receive.
- No need for communications between systems to sample/pick the same events/sessions.
- Tested in real world scenarios, against millions of events.


## Why i made this library?

I needed it to build a tracing system that samples a complete session/flow,
to gain visibility on a data pipeline that receives events/messages
with a single key/id identifying a complete session.


## Documentation
[GoDoc](https://godoc.org/github.com/lafikl/hksample)


# License
```
BSD 3-Clause License

Copyright (c) 2017, Khalid Lafi
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

* Neither the name of the copyright holder nor the names of its
  contributors may be used to endorse or promote products derived from
  this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

```
