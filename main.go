// Shamir's Secret sharing
// f(x) = a0 + a1x + a2x^2 + ... + a[k-1]x^k-1

package cs588

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

