# Autocompleteth

Go web service that provides autocompletion of terms, drawn from
the works of Shakespeare.

## Usage

Clone the repo, build, and run:
```
  $ git clone git@github.com:rwtnorton/autocompleteth.git
  $ cd autocompleteth
  $ go build && ./autocompleteth
```

Once the server is running, use `curl` to exercise the service:
```
  $ curl 'http://localhost:9000/autocomplete?term=th' 2>/dev/null |jq '.'
```

Optionally, there are some tests, which are exercised in the usual way:
```
  $ cd autocompleteth
  $ go test -count=1 ./...
```

## Author

Richard W. Norton, <rwtnorton@gmail.com>
