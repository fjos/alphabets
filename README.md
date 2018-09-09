# alphabets

GET /api/v1/pangram/{input}

Checks if the provided {input} is a pangram based on the latin alphabet.
To run: go run main.go

To test: curl http://localhost:8080/api/v1/pangram/this sentence is not a pangram

Example:
GET /api/v1/pangram/the quick brown fox jumps over the lazy dog
{
	"status" : 200,
	"error" : "",
	"data" : 
	{
		"alphabet" : "latin",
		"pangram" : true
	}
}

GET /api/v1/pangram/custom/{alphabet}/{input}
Checks the provided sentence is a pangram for the given alphabet. Currently available alphabet is latin.

GET /api/v1/pangram/custom/latin/the quick brown fox jumps over the lazy dog
Response:
{
	"status" : 200,
	"error" : "",
	"data" : 
	{
		"alphabet" : "latin",
		"pangram" : true
	}
}

GET /api/v1/pangram/custom/egyptian/this sentence contains no characters that are egyptian
Response:
{
	"status" : 400,
	"error" : "Selected alphabet 'egyptian' does not exist.",
	"data" : 
	{
		"alphabet" : "egyptian",
		"pangram" : false
	}

}


POST /api/v1/pangram
Example: 
curl -X POST -d "{\"alphabet\": \"latin\", \"input\":\"the quick brown fox jumps over the lazy dog\"}" http://localhost:8080/api/v1/pangram


Request:
{
	"alphabet" : "latin",
	"input" : "the quick brown fox jumps over the lazy dog"
}


Response:
{
	"status" : 200,
	"error" : "",
	"data" : 
	{
		"alphabet" : "latin",
		"pangram" : true
	}
}
