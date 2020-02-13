# default algoritmh
ALGORITMH?=[ SUM [ DIF ./testdata/a.txt ./testdata/b.txt ./testdata/c.txt ] [ INT ./testdata/b.txt ./testdata/c.txt ] ]

# build binary file
build:
	go build -o scalc cmd/main.go

# run the solution once
# if you want to add custom algoritmh you should run like this example:
# make run-once ALGORITMH='[ SUM ./testdata/a.txt ./testdata/b.txt ./testdata/c.txt ]'
run-once:
	go run cmd/main.go $(ALGORITMH)

# run the solution as the interactive mode
run-interactive:
	go run cmd/main.go