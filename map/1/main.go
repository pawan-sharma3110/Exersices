package main

import "fmt"

func main() {
    // Creating an empty map to store player scores
    scores := make(map[string]int)

    // Adding player scores
    scores["Alice"] = 10
    scores["Bob"] = 15
    scores["Charlie"] = 20

    // Printing all player scores
    fmt.Println("Player Scores:")
    for player, score := range scores {
        fmt.Printf("%s: %d\n", player, score)
    }

    // Updating a player's score
    scores["Bob"] = 18

    // Printing Bob's updated score
    fmt.Printf("\nBob's Updated Score: %d\n", scores["Bob"])

    // Deleting a player's score
    delete(scores, "Charlie")

    // Printing all player scores after deletion
    fmt.Println("\nPlayer Scores After Deletion:")
    for player, score := range scores {
        fmt.Printf("%s: %d\n", player, score)
    }
}
