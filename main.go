// Shamir's Secret sharing
// f(x) = a0 + a1x + a2x^2 + ... + a[k-1]x^k-1

package main

import "fmt"
import "log"
import "os"
import "strconv"

func main() {
    setupLog()
    k, err := strconv.Atoi(os.Args[1])
    if err != nil {
        println("Error converting " + os.Args[1] + " to int.")
        return
    }
    
    n, err := strconv.Atoi(os.Args[2])
    if err != nil {
        println("Error converting " + os.Args[1] + " to int.")
        return
    }
    secret := os.Args[3]
    
    if n <= 2 {
        println("n must be greater than 2")
        return
    }
    if k <= 1 {
        println("k must be greater than 1")
        return
    }
    if n <= k {
        println("k must be greater than n")
        return
    }
    createShares(k, n, secret)
}

func setupLog(){
    // open a file
    f, err := os.OpenFile("cs588.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }

    // don't forget to close it
    defer f.Close()

    // assign it to the standard logger
    log.SetOutput(f)
}

