# urlValues
Golang's net/url values edited to allow encoding without sorting. It is highly recommended to use the built-in Encode function in net/url, however some web APIs require it to be in a certain order.

# Installation
```
go get -u github.com/RambIing/urlValues
```
# Usage
```go
package main
import (
  "fmt"
  "github.com/RambIing/urlValues"
)
func Encode() {
  values := urlValues.Values{}
  values.Add("xyz", "abc")
  values.Add("abc", "xyz")
  fmt.Println(values.Encode())
  // Expected output:
  // abc=xyz&xyz=abc
  fmt.Println(values.EncodeWithOrder())
  // Expected output:
  // xyz=abc&abc=xyz
}
```
# Notice

This package is meant to be lightweight, so its missing parts of net/url. This is only intended to add and encode values.

# Copyright
This code uses parts of net/url code, and with BSD license, I must include the copyright below.

Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
