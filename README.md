## HTTP Mock Server

Simple, fast, runtime configurable HTTP mock server.

Unlike most similar mock servers, you can change what a specific URL returns at runtime. So you can point your application to the same URL, and try different responses in time.

### Usage

Just GET any URL and you'll get a default response with status code 200 and empty body:

```GET /foo/bar```

To change what a specific URL returns, GET the URL with special query params:

* Change `/foo/bar` to return `hello world` in the body

```GET /foo/bar?__body=hello%20world```

* Change `/foo/bar/` status code to `500`

```GET /foo/bar?__status=500```

* Change `/foo/bar/` status code to return `200` 60% percent of the time, then return `500` for the remaining 40%

```GET /foo/bar?__status=60:200:500```



### TODOs

* Support for all HTTP verbs beside GET
* More overridable values like Content-Type
* Preconfigure routes at startup? (easy workaround is a startup script that initializes routes with curl)