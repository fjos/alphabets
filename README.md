# alphabets

GET /api/v1/pangram/{input}

Checks if the provided {input} is a pangram based on the latin alphabet.
To run: ```go run main.go```

To test:  ```curl http://localhost:8080/api/v1/pangram/latin/this sentence is not a pangram```


GET /api/v1/pangram/{alphabet}/{input}
Checks the provided sentence is a pangram for the given alphabet. Currently available alphabet is "latin".

GET /api/v1/pangram/latin/the quick brown fox jumps over the lazy dog
Response:
```json
{
	"status" : 200,
	"error" : "",
	"data" :
	{
		"alphabet" : "latin",
		"pangram" : true
	}
}
```

#Alternative Alphabets
GET /api/v1/pangram/egyptian/this sentence contains no characters that are egyptian

Command Line: curl "http://localhost:8080/api/v1/pangram/egyptian/this sentence contains no characters that are egyptian"
Response:
```json
{
	"status" : 400,
	"error" : "Selected alphabet 'egyptian' does not exist.",
	"data" :
	{
		"alphabet" : "egyptian",
		"pangram" : false
	}
}
```


POST /api/v1/pangram
Example:
```shell
curl -X POST -d "{\"alphabet\": \"latin\", \"input\":\"the quick brown fox jumps over the lazy dog\"}" http://localhost:8080/api/v1/pangram
```

Request:
```json
{
	"alphabet" : "latin",
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
		"alphabet" : "latin",
		"pangram" : true
	}
}
```

#Future improvements
Alphabet type "custom" where you specify the characters that must exist e.g. to test for vowels
```json
{
	"alphabet" : "custom",
	"customAlphabet" : "aeiou",
	"input" : "the quick brown fox jumps over the lazy dog"
}
```

Implement testing for Response Generation, and move that code into it's own "model" package. Currently it presumes proper creation, but there is no test coverage.
Test coverage for alternative alphabets.
Test coverage for alphabets containing emojis, etc.
With a "custom" alphabet, this could be more easily implemented.

Create an alphabet struct that generates its own runemap; extract that functionality from the pangram check.

POST with a file handle; to find if a pangram exists in an extremely large string. Will cut off file upload and return Response.

SSL: secure pangram testing

More robust logging.
