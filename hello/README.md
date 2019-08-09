# Hello Example
This example demonstrates how you implement a very simple micro-service using github.com/go-msvc/domain.
The service has two operations:
* *greet* responds with "Hi"
* *sing*  responds with "La-la-la"

The example is served as HTTP REST on localhost:12345 so you get:
curl -XGET 'http://localhost:12345/hello/greet"
Hi
curl -XGET 'http://localhost:12345/hello/sing"
La-la-la
