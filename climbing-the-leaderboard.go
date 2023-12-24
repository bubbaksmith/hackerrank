package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

type ranking struct{
    score int32
    rank int32
}

func populateLeaderBoard(scores []int32) []ranking {
    var leaderboard []ranking
    for index, score := range scores {
        
        r := ranking{}
        // First score will always be first place
        if index == 0 {
            r.score = score
            r.rank = 1
        } else {
            // if score is same, copy rank
            if leaderboard[index-1].score == score {
                r.rank = leaderboard[index-1].rank    
                r.score = score
            } else {
                // score is not the same, new rank
                r.score = score
                r.rank = leaderboard[index-1].rank + 1
            }   
        }
        leaderboard = append(leaderboard, r)
    }
    return leaderboard
}

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    leaderboard := populateLeaderBoard(ranked)
    // fmt.Printf("%#v\n", leaderboard)
   
    var ranks []int32 
    for _, score := range player {
        // fmt.Printf("score = %d\n", score)
        var found_rank int32
        
        for _, ranking := range leaderboard {
            if score == ranking.score || score > ranking.score {
                // fmt.Printf("Rank Found: %d, comparing player score '%d' to '%d'\n", ranking.rank, score, ranking.score)
                found_rank = ranking.rank
                break
            } 
        }
        
        if found_rank != 0 {
            ranks = append(ranks, found_rank)
        } else {
            ranks = append(ranks, leaderboard[len(leaderboard)-1].rank + 1)
        }
        
        //fmt.Printf("%#v\n", ranks)
    }

    return ranks
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ranked []int32

    for i := 0; i < int(rankedCount); i++ {
        rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
        checkError(err)
        rankedItem := int32(rankedItemTemp)
        ranked = append(ranked, rankedItem)
    }

    playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var player []int32

    for i := 0; i < int(playerCount); i++ {
        playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
        checkError(err)
        playerItem := int32(playerItemTemp)
        player = append(player, playerItem)
    }

    result := climbingLeaderboard(ranked, player)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

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
