package main

type TypeOFRequest int

const (
	GetSingle TypeOFRequest = iota
	GetAll
	Post
	Delete
	Put
)
