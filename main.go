package main



import (
"log"
"fmt"
"io/ioutil"
"encoding/json"
"net/http"
"github.com/gorilla/mux"
)



type Player struct {
Id string `json:"playerid"`
PName string `json:"name"`
PTeam string `json:"team"`
}



type Score struct {
Id string `json:"scoreid,omitempty"`
Match string `json:"match"`
Runs string `json:"runs"`
Wickets string `json:"wickets"`
}




var players []Player
var scores []Score



func homePage(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Welcome to the HomePage!")
fmt.Println("Endpoint Hit:homepage")
}



func getAllPlayers(w http.ResponseWriter, r *http.Request) {
fmt.Println("Endpoint Hit:returnAllPlayers")
json.NewEncoder(w).Encode(players)
}



func getSinglePlayer(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
key := vars["id"]



for _, player := range players {
if player.Id == key {
json.NewEncoder(w).Encode(player)
}
}
}



func createNewPlayer(w http.ResponseWriter, r *http.Request) {



reqBody, _ := ioutil.ReadAll(r.Body)
var player Player
json.Unmarshal(reqBody, &player)



players = append(players, player)



json.NewEncoder(w).Encode(player)
}
func createNewScore(w http.ResponseWriter, r *http.Request) {
reqBody, _ := ioutil.ReadAll(r.Body)
var score Score
json.Unmarshal(reqBody, &score)



scores = append(scores, score)



json.NewEncoder(w).Encode(score)
}



func deletePlayer(w http.ResponseWriter, r *http.Request) {
vars := mux.Vars(r)
id := vars["id"]



for index, player := range players {
if player.Id == id {
players = append(players[:index], players[index+1:]...)
}
}
}




func handleRequests() {
router := mux.NewRouter().StrictSlash(true)
router.HandleFunc("/", homePage)
router.HandleFunc("/player", createNewPlayer).Methods("POST")
router.HandleFunc("/player/{id}/score", createNewScore).Methods("POST")
router.HandleFunc("/player/{id}", deletePlayer).Methods("DELETE")
router.HandleFunc("/players", getAllPlayers).Methods("GET")
router.HandleFunc("/players/scores", getSinglePlayer).Methods("GET")
log.Fatal(http.ListenAndServe(":8080", router))




}



func main(){



players = []Player{
{Id: "1", PName: "Virat Kohli", PTeam: "RCB"},
{Id: "2", PName: "M S Dhoni", PTeam: "CSK"},
}

handleRequests()
}