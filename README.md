<h1 align="center">Welcome to go-youtube 👋 *** Work In Progress ***</h1>
<p align="center">
  <a href="https://github.com/tsirysndr/go-youtube/commits/master">
    <img src="https://img.shields.io/github/last-commit/tsirysndr/go-youtube.svg" target="_blank" />
  </a>
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/tsirysndr/go-youtube">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/tsirysndr/go-youtube">
  <img alt="GitHub pull requests" src="https://img.shields.io/github/issues-pr/tsirysndr/go-youtube">
  <img alt="GitHub issues" src="https://img.shields.io/github/issues/tsirysndr/go-youtube">
  <img alt="GitHub contributors" src="https://img.shields.io/github/contributors/tsirysndr/go-youtube">
  <a href="https://github.com/tsirysndr/go-youtube/blob/master/LICENSE">
    <img alt="License: BSD" src="https://img.shields.io/badge/license-BSD-green.svg" target="_blank" />
  </a>
  <a href="https://twitter.com/tsiry_sndr">
    <img alt="Twitter: tsiry_sndr" src="https://img.shields.io/twitter/follow/tsiry_sndr.svg?style=social" target="_blank" />
  </a>
</p>

go-youtube is a Go client library for accessing the [YouTube API](https://developers.google.com/youtube/v3/docs)


## Install

```sh
go get github.com/tsirysndr/go-youtube
```

## Usage

Import the package into your project.

```Go
import "github.com/tsirysndr/go-youtube"
```

Construct a new YouTube client, then use the various services on the client to access different parts of the YouTube API. For example:

```Go
client := youtube.NewClient("<YOUR TOKEN AUTHORIZATION>")
```


## Author

👤 **Tsiry Sandratraina**

* Twitter: [@tsiry_sndr](https://twitter.com/tsiry_sndr)
* Github: [@tsirysndr](https://github.com/tsirysndr)

## Show your support

Give a ⭐️ if this project helped you!

