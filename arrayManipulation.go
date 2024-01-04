// https://www.hackerrank.com/challenges/crush/problem

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
    "sync"
)

var wg sync.WaitGroup

func incrementArr(n *int, k int32) {
    defer wg.Done()
    
    *n += int(k)
}

func arrayManipulation(n int32, queries [][]int32) (max int) {
    arr := make([]int, n)
    fmt.Printf("%#v\n", arr)
    
    for _, query := range queries {
        a,b,k := query[0], query[1], query[2]
        
        
        for i:=a-1;i<b;i++{
            // wg.Add(1)
            // go incrementArår(&arr[i], k)
            arr[i] += int(k)
        }
        
        wg.Wait()
    }
    
    fmt.Printf("%#v\n", arr)
    
    // find max value element
    max = arr[0]
    for i:=int32(1);i<n;i++ {
        if arr[i] > max {
            max = arr[i]
        }
    }
    
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

    var queries [][]int32
    for i := 0; i < int(m); i++ {
        queriesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != 3 {
            panic("Bad input")
        }

        queries = append(queries, queriesRow)
    }

    result := arrayManipulation(n, queries)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
