# alphabets
A "RESTful" style utility that determines if a provided string is a pangram (contains all letters of an alphabet).
Can support pangram checking for different alphabets.

To run: `go run main.go`

To test:  `curl "http://localhost:8080/api/v1/pangram/latin/this%20sentence%20is%20not%20a%20pangram`
## Pangram

### GET /api/v1/pangram/{alphabet}/{input}

Checks the provided sentence is a pangram for the given alphabet. Currently available alphabet is "latin".

#### Examples
GET /api/v1/pangram/latin/the quick brown fox jumps over the lazy dog

`curl "http://localhost:8080/api/v1/pangram/latin/the%20quick%20brown%20fox%20jumps%20over%20the%20lazy%20dog"`

Response:
```json
{
	"status" : 200,
	"error" : "",
	"data" :
	{
		"alphabet" : {
			"name" : "latin",
			"contents" : "abcdefghikljmnopqrstuvwxyz"
	},
		"pangram" : true
	}
}
```

#Alternative Alphabets
GET /api/v1/pangram/egyptian/this sentence contains no characters that are egyptian

Command Line: `curl "http://localhost:8080/api/v1/pangram/egyptian/this%20sentence%20contains%20no%20characters%20that%20are%20egyptian"`
Response:
```json
{
	"status" : 400,
	"error" : "Selected alphabet 'egyptian' does not exist.",
	"data" :
	{
		"alphabet" : {
			"name" : "egyptian",
			"contents" : ""
	},
		"pangram" : false
	}
}
```



### POST /api/v1/pangram
Allows you to provide json formatted request to check if it is a pangram, also allowing you to check inputs without manually replacing spaces
#### Examples

Latin alphabet post:
```shell
curl -X POST -d "{\"alphabet\":{\"name\":\"latin\"}, \"input\":\"the quick brown fox jumps over the lazy dog\"}" http://localhost:8080/api/v1/pangram
```

Request:
```json
{
	"alphabet" : {
		"name" : "latin"
	},
	"input" : "the quick brown fox jumps over the lazy dog"
}
```

Response:
```json
{
	"status" : 200,
	"error" : "",
	"data" :
	{
		"alphabet" : {
			"name" : "latin",
			"contents" : "abcdefghikljmnopqrstuvwxyz"
	},
		"pangram" : true
	}
}
```

Custom Alphabet
POST /api/v1/pangram

```shell
curl -X POST -d "{\"alphabet\":{\"name\":\"custom\", \"contents\" : \"aeiou\"}, \"input\":\"wait I know this you're famous\"}" http://localhost:8080/api/v1/pangram
```


```json
{
	"alphabet" : {
		"name" : "custom",
		"contents" : "aeiou"
	},
	"input" : "wait I know this you're famous"
}

```

### GET /debug/vars

Uses the expvar package to export useful statistics about server responses, commonly requested alphabets, etc.

## Future improvements

Implement testing for Response Generation, and move that code into it's own "model" package. Currently it presumes proper creation, but there is no test coverage.


POST with a file handle; to find if a pangram exists in an extremely large string. Will cut off file upload and return Response.

SSL: secure pangram testing
