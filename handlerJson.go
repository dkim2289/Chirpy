package main

{
	"body": "",
}
func handlerJson(w http.ResponseWriter, r *http.Request) {
	type parameters struct{
		Body string `json: "body"`

	}
}
