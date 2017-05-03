// Shamir's Secret sharing
// f(x) = a0 + a1x + a2x^2 + ... + a[k-1]x^k-1

package main

import "fmt"
import "os"

func main() {
    k := os.Args[1]
    n := os.Args[2]
    secret := os.Args[3]
    
    fmt.Println("k: " + k)
    fmt.Println("n: " + n)
    fmt.Println("secret: " + secret)
}