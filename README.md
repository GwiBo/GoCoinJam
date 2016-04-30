# GoCoinJam
A random brute force solution to [Google Code Jam 2016 :CoinJam](https://code.google.com/codejam/contest/6254486/dashboard#s=p2) problem in golang

The provided solution solves the large problem input in `<10 seconds`!!.

Remember the style is not necessarily `idiomatic Go`, however it uses several ( around 200 ) [goroutines](https://tour.golang.org/concurrency/1) ( I know Go can easily handle way more... ) for each coin condition checking as well as a single goroutine to generate a random coin of required size.


#####Getting Program
 - You need [go installed](https://golang.org/doc/install)
 - Then run command on terminal/[bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell)): `$ go get github.com/GwiBo/GoCoinJam`  
    The command wil download source files, compile and install binary in [`$GOPATH/bin`](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable) directory
 - There is a sample input file available in [$GOPATH/src/github.com/GwiBo/GoCoinJam/C-Large.in](https://golang.org/doc/code.html#Workspaces)


#####Usage instruction
 - The program takes input from [stdin](https://en.wikipedia.org/wiki/Standard_streams#Standard_input_.28stdin.29) (standard input) and outputs to [stdout](https://en.wikipedia.org/wiki/Standard_streams#Standard_output_.28stdout.29) (standard output)
 - the first line takes number of test cases
 - the second line takes 2 (space separated) arguments : `size of coins` `number of coins to generate`

#####Hint
 - When running on bash terminal, you can use pipes to get result  
i.e. `$ GoCoinJam < input.file > output.file`
 - With go properly installed / environment variables setup the full command to run the sample is as follows  
    `$ $GOPATH/bin/GoCoinJam < $GOPATH/src/github.com/GwiBo/GoCoinJam/C-Large.in > sample.out` this will place `sample.out` output file in your current working directory
 - If the program runs for more than (around) 30 seconds, then probably: either the input cannot be satisfied OR badluck ( run again )



###Test Hardware:

| Hardware  | Details                                          |
|-----------|--------------------------------------------------|
| CPU       |Intel `Core i7` first generation (mobile edition) |
| Ram       | 8 GB                                             |

|Software              | Details                                                            |
|----------------------|--------------------------------------------------------------------|
| OS                   | [Ubuntu](http://www.ubuntu.com/) 16.04                             |
| Programming Language | [go 1.6.2](https://golang.org/) ( currently, originally used go 1.6|


####And the BIG question : Why such elaborate README.md for such trivial program?
`Ans:` Just wanted to scratch an itch to write a Markdown text document. :) :) :)
