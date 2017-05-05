package main

import "fmt"
import "log"
import "strconv"
import "encoding/binary"
import "math"

func createShares(k int, n int, secret string) (shares []string) {
    fmt.Println("k: " + strconv.Itoa(k))
    fmt.Println("n: " + strconv.Itoa(n))
    fmt.Println("secret: " + secret)

    secretBytes := []byte(secret)
    if len(secretBytes) % 8 != 0{
        addZeros := 8 - (len(secretBytes) % 8)
        for i := 0; i < addZeros; i++ {
            secretBytes = append(secretBytes, 0)
        }
    }
    
    sliceSecretLength := int(math.Ceil(float64(len(secretBytes))/8))
    log.Println("slice share length: ", sliceSecretLength)
    
    secretSlices := make([][]byte, sliceSecretLength)
    
    for index := 0; index < len(secretSlices); index++{
	    secretSlices[index] = make([]byte, 8)
	    secretSlices[index][0] = secretBytes[index*8]
	    secretSlices[index][1] = secretBytes[index*8+1]
	    secretSlices[index][2] = secretBytes[index*8+2]
	    secretSlices[index][3] = secretBytes[index*8+3]
	    secretSlices[index][4] = secretBytes[index*8+4]
	    secretSlices[index][5] = secretBytes[index*8+5]
	    secretSlices[index][6] = secretBytes[index*8+6]
	    secretSlices[index][7] = secretBytes[index*8+7]
    }
    
    secretIntSlices := make([]uint64, sliceSecretLength)
    for index := 0; index < len(secretIntSlices); index++ {
        secretIntSlices[index] = binary.BigEndian.Uint64(secretSlices[index])
        log.Println("secret int slice [", index, "]: ", secretIntSlices[index])
    }

    shareIntSlices := make([][]uint64, sliceSecretLength)
    for index := 0; index < len(shareIntSlices); index++ {
        shareIntSlices[index] = getShareSegments(k, n, secretIntSlices[index])
        log.Println("shareIntSlices[", index, "]: ", shareIntSlices[index])
    }
    
    shares = make([]string, n)
    for indexI := 0; indexI < len(shareIntSlices); indexI++ {
        for indexJ := 0; indexJ < len(shareIntSlices[indexI]); indexJ++ {
            shares[indexJ] = toString(shareIntSlices[indexI][indexJ])
            log.Println("shareIntSlices[", indexI, "][", indexJ, "]: ", shareIntSlices[indexI][indexJ])
        }
    }
    
    return
}

func combineShares(shares []string) (secret string) {
    secret = "I am a pretty pony."
    return
}
