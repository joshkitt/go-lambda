package main

import (
    "context"
//    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/aws/aws-lambda-go/lambda"
)

// Request struct to define the incoming request format
type Request struct {
    URL string `json:"url"`
}

// Response struct to define the response format
type Response struct {
    StatusCode int    `json:"statusCode"`
    Body       string `json:"body"`
}

func handler(ctx context.Context, req Request) (Response, error) {
    resp, err := http.Get(req.URL)
    if err != nil {
        return Response{StatusCode: 500, Body: fmt.Sprintf("Error: %s", err)}, nil
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return Response{StatusCode: 500, Body: fmt.Sprintf("Error reading response body: %s", err)}, nil
    }

    return Response{StatusCode: 200, Body: string(body)}, nil
}

func main() {
    lambda.Start(handler)
}

