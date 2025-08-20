package handler

import (
    "encoding/json"
    "net/http"
    "github.com/Jisin0/filmigo/imdb"
    "github.com/Jisin0/filmigo/omdb" 
    "github.com/Jisin0/filmigo/justwatch"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    // Your existing filmigo code works here directly!
    client := imdb.NewClient()
    movie, err := client.GetMovie(r.URL.Query().Get("id"))
    
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(movie)
}
