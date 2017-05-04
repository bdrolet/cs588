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
    
    sliceShareLength := int(math.Ceil(float64(len(secretBytes))/8))
    log.Println("slice share length: ", sliceShareLength)
    
    shareSlices := make([][]byte, sliceShareLength)
    
    for index := 0; index < sliceShareLength; index++{
	    shareSlices[index] = make([]byte, 8)
	    shareSlices[index][0] = secretBytes[index*8]
	    shareSlices[index][1] = secretBytes[index*8+1]
	    shareSlices[index][2] = secretBytes[index*8+2]
	    shareSlices[index][3] = secretBytes[index*8+3]
	    shareSlices[index][4] = secretBytes[index*8+4]
	    shareSlices[index][5] = secretBytes[index*8+5]
	    shareSlices[index][6] = secretBytes[index*8+6]
	    shareSlices[index][7] = secretBytes[index*8+7]
    }
    
    shareIntSlices := make([]uint64, sliceShareLength)
    
    for index := 0; index < sliceShareLength; index++{
        shareIntSlices[index] = binary.BigEndian.Uint64(shareSlices[index])
        println(shareIntSlices[index])
    }

    for index := 0; index < sliceShareLength; index++{
        getShareSegments(k,n, shareIntSlices[index])
    }

    shares = []string{"1","2"}
    return
}

func combineShares(shares []string) (secret string) {
    secret = "I am a pretty pony."
    return
}
