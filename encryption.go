// https://www.hackerrank.com/challenges/encryption

package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "math"
)

type table struct {
    rows []string
    numColumns int
    numRows int
}

func (t* table) printEncrypted() string {
    s := []byte{}
    
    for i:=0;i<=t.numColumns-1;i++ {
        for j:=0;j<=t.numRows-1;j++ {
            
           // handle last column differently
           if j == t.numRows-1  {
               
               // elements may not exist
               if i < len(t.rows[j]) {
                   s = append(s, t.rows[j][i])
               }
               // and add a space
               s = append(s, byte(' '))
           } else {
               s = append(s, t.rows[j][i])
           }
           
        }
    }
    return string(s[:])
}


func encryption(s string) string {
    withoutSpaces := strings.ReplaceAll(s, " ", "")
    fmt.Printf("%s\n", withoutSpaces)
    t := table{}
    
    sqrtOfS := math.Sqrt(float64(len(s)))
    t.numRows = int(math.Floor(sqrtOfS))
    t.numColumns = int(math.Ceil(sqrtOfS))
    
    // Increment numRows based on problem constraints
    if t.numColumns * t.numRows < len(withoutSpaces) {
        t.numRows++
    }
    
    for i:=0;i<t.numRows;i++ {
        startIndex := i * t.numColumns
        endIndex := (i+1) * t.numColumns
        
        // Check if partial row
        if (i+1)*t.numColumns > len(withoutSpaces) {
            endIndex = len(withoutSpaces)
        }
        
        newRow := withoutSpaces[startIndex:endIndex]
        fmt.Printf("newRow: %s\n", newRow)
        t.rows = append(t.rows, newRow)
    }
    
    fmt.Printf("Table: %#v\n", t)
    
    return t.printEncrypted()
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := encryption(s)

    fmt.Fprintf(writer, "%s\n", result)

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
