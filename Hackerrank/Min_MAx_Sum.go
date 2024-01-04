package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sort"
)

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func miniMaxSum(arr []int32) {
    // Write your code here
    
    
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] < arr[j]
    })
    //fmt.Println(arr)
    LenArr := len(arr) - 1
    var MinSum int32 
    var MaxSum int32 
    MinSum = 0
    MaxSum =0
    for i :=0; i < LenArr ; i++{
        MinSum += arr[i]
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i] > arr[j]
    })
   // fmt.Println(arr)
     for i :=0; i < LenArr ; i++{
        MaxSum += arr[i]
    }


    fmt.Println(MinSum,MaxSum)
        
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var arr []int32

    for i := 0; i < 5; i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    miniMaxSum(arr)
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
