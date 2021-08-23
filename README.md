# http-echo

HTTP Echo writes back specified HTTP request to the response. It can also optionally add a delay in writing the response.

Currently, the following HTTP methods are supported:

1. **GET**

   HTTP Echo echoes back the `message` query parameter value.

   ```bash
   $ curl 'http://localhost:8080/?message=foo'
   foo
   ```

2. **POST**

   HTTP Echo echoes back the HTTP request body data.

   ```bash
   $ curl -X POST -d '{"foo": "bar"}' 'http://localhost:8080'
   {"foo": "bar"}
   ```

## Running the server

By default, HTTP Echo uses port 8080. To use another port, set `PORT` environment variable.

```sh
PORT=9090 go run main.go
```
