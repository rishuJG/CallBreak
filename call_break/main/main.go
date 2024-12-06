package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
   "bitbucket.org/junglee_games/call_break/game"
)

func main() {
    fmt.Println("Welcome to Call Break!")
    fmt.Println("----------------------")
    
    players := getPlayerNames()
    engine := game.NewGameEngine(players)
    engine.StartGame()
}

func getPlayerNames() []string {
    reader := bufio.NewReader(os.Stdin)
    players := make([]string, 4)
    
    for i := 0; i < 4; i++ {
        fmt.Printf("Enter name for Player %d: ", i+1)
        name, _ := reader.ReadString('\n')
        players[i] = strings.TrimSpace(name)
    }
    
    return players
} 