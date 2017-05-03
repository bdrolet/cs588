// Shamir's Secret sharing
// f(x) = a0 + a1x + a2x^2 + ... + a[k-1]x^k-1

package main

import "fmt"
import "os"
import "strconv"

func main() {
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
    
    shamir(k, n, secret)
}

func shamir(k int, n int, secret string){
    fmt.Println("k: " + strconv.Itoa(k))
    fmt.Println("n: " + strconv.Itoa(n))
    fmt.Println("secret: " + secret)
}