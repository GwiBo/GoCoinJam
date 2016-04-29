# GoCoinJam
A random brute force solution to [Google Code Jam 2016 :CoinJam](https://code.google.com/codejam/contest/6254486/dashboard#s=p2) problem in golang

The provided solution solves the large problem input in `<10 seconds`!!.

Remember the style is not necessarily `idiomatic Go`, however it uses several ( around 200 ) [goroutines](https://tour.golang.org/concurrency/1) ( I know Go can easily handle way more... ) for each coin condition checking as well as a single goroutine to generate a random coin of required size.

###Program Usage
 - The program takes input from `stdin`(standard input) and outputs to `stdout`(standard output)
 - the first line takes number of test cases
 - the second line takes `size of coins` `number of coins to generate`
 
When running on bash terminal, you can use pipes to get result  
i.e. `$ GoCoinJam < input.file > output.file`

###Test Hardware:

| Hardware  | Details                                          |
|-----------|--------------------------------------------------|
| CPU       |Intel `Core i7` first generation (mobile edition) |
| Ram       | 8 GB                                             |

|Software              | Details                                                            |
|----------------------|--------------------------------------------------------------------|
| OS                   | [Ubuntu](http://www.ubuntu.com/) 16.04                             |
| Programming Language | [go 1.6.2](https://golang.org/) ( currently, originally used go 1.6|
