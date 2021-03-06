// Shamir's Secret sharing
// f(x) = a0 + a1x + a2x^2 + ... + a[k-1]x^k-1

package main	

import "fmt"
import "log"
import "os"
import "strconv"

func main() {
    // open a file
    f, err := os.OpenFile("cs588.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
    if err != nil {
        fmt.Printf("error opening file: %v", err)
    }

    // don't forget to close it
    defer f.Close()

    // assign it to the standard logger
    log.SetOutput(f)
    
    if os.Args[1] == "create" {        
        k, err := strconv.Atoi(os.Args[2])
        if err != nil {
            println("Error converting " + os.Args[1] + " to int.")
            return
        }
        
        n, err := strconv.Atoi(os.Args[3])
        if err != nil {
            println("Error converting " + os.Args[3] + " to int.")
            return
        }
        secret := os.Args[4]
        
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
        
        shares := createShares(k, n, secret)
        writeShare(shares)
    } else if os.Args[1] == "combine" {
        
    }
}
    
func writeShare(shares []string){
    sharesFile, err := os.OpenFile("shares.txt", os.O_CREATE | os.O_TRUNC, 0666)
    if err != nil {
        log.Printf("error opening file: %v", err)
    }
    
    for index := 0; index < len(shares); index++ {
        share := fmt.Sprintf("%[1]v, %[2]v\n", (index+1), shares[index])
        log.Println(share)
        sharesFile.WriteString(share)
    }
    
    sharesFile.Close()
}