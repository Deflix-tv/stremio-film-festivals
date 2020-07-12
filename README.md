stremio-film-festivals
==================

Stremio addon for multiple film festival catalogs:

- Academy Award for Best Picture winners
- Cannes Film Festival Palme d'Or winners
- Venice Film Festival Golden Lion winners
- Berlin International Film Festival Golden Bear winners

Contents
--------

1. [Install](#install)
2. [Run locally](#run-locally)
   1. [Configuration](#configuration)

Install
-------

This addon is a remote addon, so it's an HTTP web service and Stremio just sends HTTP requests to it. You dont't need to run any untrusted code on your machine.

You only have to enter the addon URL in the search box of the addons section of Stremio, like this:  
`https://stremio-film-festivals.deflix.tv/manifest.json`

That's it!

Run locally
-----------

Alternatively you can also run the addon locally and use that in Stremio. The addon is written in Go and compiles to a single executable file without dependencies, so it's really easy to run on your machine.

You can use one of the precompiled binaries from GitHub:

1. Download the binary for your OS from <https://github.com/Deflix-tv/stremio-film-festivals/releases>
2. Simply run the executable binary
3. To stop the program press `Ctrl-C` (or `⌃-C` on macOS)

Or use Docker:

1. `docker pull doingodswork/stremio-film-festivals`
2. `docker run --name stremio-film-festivals -v /path/to/data:/data -p 8080:8080 doingodswork/stremio-film-festivals`
3. To stop the container: `docker stop stremio-film-festivals`

Then similar to installing the publicly hosted addon you just enter the following URL in the search box of the addon section of Stremio:  
`http://localhost:8080/manifest.json`

### Configuration

The following options can be configured via command line argument:

```text
Usage of stremio-film-festivals:
  -bindAddr string
        Local interface address to bind to. "localhost" only allows access from the local host. "0.0.0.0" binds to all network interfaces. (default "localhost")
  -cacheAge string
        Max age for a client or proxy cache. The format must be acceptable by Go's 'time.ParseDuration()', for example "24h". (default "24h")
  -dataDir string
        Location of the data directory. It contains CSV files with IMDb IDs and a "metas" subdirectory with meta JSON files (default ".")
  -logLevel string
        Log level to show only logs with the given and more severe levels. Can be "debug", "info", "warn", "error" (default "info")
  -port int
        Port to listen on (default 8080)
```
