package main

import (
	"time"

	"github.com/bryanwsebaraj/httpserver/server"
)

func AboutHandler(req *server.Request) *server.Response {
	println(req.AsString())
	return &server.Response{
		Version:    req.Version,
		StatusCode: server.OK,
		StatusText: "OK",
		Headers:    map[string]string{"Content-Type": "text/html", "Date": time.Now().String()},
		Body:       "<h1>About</h1>",
	}

}

func HomeHandler(req *server.Request) *server.Response {
	retBody := ""
	if req.Method == server.GET {
		retBody = "<h1>Home Get </h1>"
	} else {
		retBody = "<h1>Home Post </h1>"
	}
	return &server.Response{
		Version:    req.Version,
		StatusCode: server.OK,
		StatusText: "OK",
		Headers:    map[string]string{"Content-Type": "text/html", "Date": time.Now().String()},
		Body:       retBody,
	}
}

func PathHandler(req *server.Request) *server.Response {
	pathVars := server.GetPathVars(req, "/test/{path}")
	println(pathVars)
	return &server.Response{
		Version:    req.Version,
		StatusCode: server.OK,
		StatusText: "OK",
		Headers:    map[string]string{"Content-Type": "text/html", "Date": time.Now().String()},
		Body:       "<h1>Path</h1>" + req.URI,
	}

}

func IconHandler(req *server.Request) *server.Response {
	return &server.Response{
		Version:    req.Version,
		StatusCode: server.OK,
		StatusText: "OK",
		Headers:    map[string]string{"Content-Type": "text/html", "Date": time.Now().String()},
		Body:       "<h1>Icon</h1>",
	}

}

func JsonExample(req *server.Request) *server.Response {
	return &server.Response{
		Version:    req.Version,
		StatusCode: server.OK,
		StatusText: "OK",
		Headers:    map[string]string{"Content-Type": "application/json", "Date": time.Now().String()},
		Body:       "{\n\t\"type\": \"json\",\n\t\"message\": \"hello world\"\n}",
	}

}
