# Disclog

A simple util that prints discord messages as JSON to stdout.

## Example

```sh
disclog --token="YOUR TOKEN" --channels="600331914041360394,600332006676758548"
```

Errors will be printed to stderr. You can simply pipe the content into a
different application or a file.

## Compiling

In order to compile you need Golang (Go) version 1.13 or later.
Check your installed version via `go version`.
Or [download golang here](https://golang.org/dl/).

You also either need git or download the source-code manually.

```sh
git clone https://github.com/Bios-Marcel/disclog.git
cd disclog
go build .
```

This produces an executable called `disclog` (`disclog.exe` on windows).
