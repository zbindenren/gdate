# gdate

Set system time from HTTP server headers (useful for systems without access to ntp servers)

```console
$ sudo gdate -url https://example.com
Time Difference: 158.01579ms
Setting system date to: 30 May 2022 11:13:24
```

```console
$ gdate -h
Usage of gdate:
  -header string
        The header name that contains the timestamp (also via GDATE_HEADER env var) (default "Date")
  -layout string
        The layout to use to parse timestamp (also via GDATE_LAYOUT env var) (default "Mon, 2 Jan 2006 15:04:05 MST")
  -url string
        The url which gdate uses to parse timestamp (also via GDATE_URL env var)
```
