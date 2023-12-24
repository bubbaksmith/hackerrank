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

type vector struct {
        length int32
        direction coordinate
        origin coordinate
        current coordinate
}

func (v *vector) nextCoordinate() coordinate {
        next := coordinate{}
        next.x = v.current.x + v.direction.x
        next.y = v.current.y + v.direction.y
        return next
}

type coordinate struct {
        x int32
        y int32
}

func (c *coordinate) inBounds(boardLength int32) bool {
        if c.y == 0 || c.y > boardLength {
                return false
        }
        if c.x == 0 || c.x > boardLength {
                return false
        }
        return true
}

func (c *coordinate) noObstacle(obstacles [][]int32) bool {
        for _, obstacle := range obstacles {
                if obstacle[0] == c.y && obstacle[1] == c.x {
                        return false
                }
        }
        return true
}

var possibleDirections = [3]int32{-1,0,1}

 
// createAllDirections in the form of a one unit vector with 0,0 as its origin
func createAllDirections() []coordinate {
        directions := []coordinate{}
        for _, x := range possibleDirections{
                for _, y := range possibleDirections{
                        // 0,0 is not a direction
                        if x == 0 && y == 0 {
                                continue
                        }
                        d := coordinate{x:x, y:y}
                        directions = append(directions, d)
                }
        }
        return directions
}

// createVector for each direction
func createVectors(directions []coordinate, origin coordinate) []vector {
        vectors := []vector{}
        for _, direction := range directions {
                v := vector{length: 0, direction: direction, origin: origin, current: origin}
                vectors = append(vectors, v)
        }
        return vectors
}

func findPotentialSpaces(v vector, n int32, obstabcles [][]int32, wg *sync.WaitGroup, ch chan int32) {
        fmt.Printf("%#v\n", v)
        
        next := v.nextCoordinate()
        if next.inBounds(n) && next.noObstacle(obstabcles) {
                v.current.x = next.x
                v.current.y = next.y
                v.length += 1
                
                findPotentialSpaces(v, n, obstabcles, wg, ch)
        } else { 
                ch <- v.length
                wg.Done()
                fmt.Println("Done")
        }
}

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
        origin := coordinate{x: c_q, y: r_q}
        directions := createAllDirections()
        // fmt.Printf("%#v\n", directions)
        
        vectors := createVectors(directions, origin)
        
        ch := make(chan int32, len(vectors))
        wg := &sync.WaitGroup{}
        
        for _, vector := range vectors {
                wg.Add(1)
                go findPotentialSpaces(vector, n, obstacles, wg, ch)
        }
        wg.Wait()
        close(ch)
        
        var sum int32
        for spaces := range ch {
                sum += spaces
                fmt.Printf("%d\n", sum)
        }
        
        return sum
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

    kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
    checkError(err)
    k := int32(kTemp)

    secondMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    r_qTemp, err := strconv.ParseInt(secondMultipleInput[0], 10, 64)
    checkError(err)
    r_q := int32(r_qTemp)

    c_qTemp, err := strconv.ParseInt(secondMultipleInput[1], 10, 64)
    checkError(err)
    c_q := int32(c_qTemp)

    var obstacles [][]int32
    for i := 0; i < int(k); i++ {
        obstaclesRowTemp := strings.Split(strings.TrimRight(readLine(reader)," \t\r\n"), " ")

        var obstaclesRow []int32
        for _, obstaclesRowItem := range obstaclesRowTemp {
            obstaclesItemTemp, err := strconv.ParseInt(obstaclesRowItem, 10, 64)
            checkError(err)
            obstaclesItem := int32(obstaclesItemTemp)
            obstaclesRow = append(obstaclesRow, obstaclesItem)
        }

        if len(obstaclesRow) != 2 {
            panic("Bad input")
        }

        obstacles = append(obstacles, obstaclesRow)
    }

    result := queensAttack(n, k, r_q, c_q, obstacles)

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
