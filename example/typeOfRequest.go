package main

type TypeOFRequest int

const (
	GetSingle TypeOFRequest = iota
	GetAll
	GetAllWithQueries
	Post
	Delete
	Put
	Patch
)
