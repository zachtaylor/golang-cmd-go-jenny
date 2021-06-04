package main

// Usage is the usage information
const Usage = `usage: go-jenny [v|version] [h|help] [f=<filepath>] [i=<imports>] p=<package> [t=<typename>] k=<keytype> v=<valtype>

available options
v,version	optional				print the module version number

h,help		optional; default=false			print this help message

f		optional; default="jenny.go"		file name to generate

i		optional				comma-separated list of import paths to add (not including "sync")

t		optional; default="Map"			generated type name

p		required				package name

k		required				key type name

v 		required				value type name
`
