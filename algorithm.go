package main

import "math"
import "math/rand"
import "log"

func getShareSegments(n int, k int, secretSegment uint64) (shareSegments []uint64){
    println("In getShareSegments...")
    shareSegments = make([]uint64, n)
    coefficients := getCoeffecients(k)
    for index := 0; index < len(shareSegments); index++ {
    //for index, shareSegment := range shareSegments {
        shareSegments[index] = calculateShare(secretSegment, coefficients, index+1)
        log.Println("Calculated share: ", shareSegments[index])
        println("Calculated share: ", shareSegments[index])
    }
    return
}

func getCoeffecients(k int) (coefficients []int){
    coefficients = make([]int, k)
    for index := 0; index < len(coefficients); index++ {
    //for index, coefficient := range coefficients{
        coefficients[index] = rand.Int()
        log.Println("Coeffecients ", index, coefficients[index])
        println("Coeffecients: ", index, coefficients[index])
    }
    return
}

func calculateShare(secretSegment uint64, coefficients []int, shareNumber int) (shareSegment uint64){
    shareSegment = secretSegment
    for index, coefficient := range coefficients{
        a := uint64(coefficient)
        x := float64(shareNumber)
        degree := float64(index+1)
        shareSegment = shareSegment + a * uint64(math.Pow(x, degree))
        println("a: ", a)
        println("x: ", x)
        println("degree: ", degree)
        println("shareSegment: ", secretSegment)
    }
    return
}