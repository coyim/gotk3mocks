# Mocks for gotk3

This project provides mocks to be used for `gotk3adapter`, so that you can test the usage of `gotk3` without having to
use real types. The `gotk3adapter` project provides empty implementations of all methods, but you have to implement them
yourselves. This project instead provides mock implementations for every method and type using the `testify` mock
framework, making it easy to test and hook up to any project.

This project does make use of Generics, and is thus only compatible with Go 1.18 and up.

This code is released under GPL 3.0.
