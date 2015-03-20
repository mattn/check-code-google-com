# check-code-google-com

Check your package is depend on code.google.com

March 12, 2015, Google decide to close Google Code.

http://google-opensource.blogspot.com/2015/03/farewell-to-google-code.html

I am grateful to Google Code. When there was not GitHub, I put my codes on
Google Code. It was very stable. I had never see the server is down.
Unfortunately, Google Code will close. However we, gopher has many codes that
contains packages which is on code.google.com. We must finish transition to
find  alternative packages until the Google Code will close.

## Usage

```
$ check-code-google-com [package]
```

For example:

![Bad](http://go-gyazo.appspot.com/d7648e8179e45bf8.png)

![Good](http://go-gyazo.appspot.com/7f26fede3724d46a.png)

## Installation

```
$ go get github.com/mattn/check-code-google-com
```

## License

MIT

## Author

Yasuhiro Matsumoto (a.k.a mattn)
