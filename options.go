package main

// Options is the type consumed by Template
type Options struct {
	// Package name
	Package string
	// Type name
	Type string
	// Key type name
	Key string
	// Val type name
	Val string
	// Off default value of val type
	Off string
	// Stdlib imports from std lib
	Stdlib []string
	// Remote imports
	Remote []string
}
