package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type airport struct {
	name         string
	connections map[string]bool
}

func calculateFlightPath(flights [][]string) []string {
	airports := make(map[string]*airport)

	for _, flight := range flights {
		source, destination := flight[0], flight[1]

		if airports[source] == nil {
			airports[source] = &airport{name: source, connections: make(map[string]bool)}
		}
		if airports[destination] == nil {
			airports[destination] = &airport{name: destination, connections: make(map[string]bool)}
		}

		// connect source and destination airports
		airports[source].connections[destination] = true
	}

	// find the starting and ending airports
	var start, end *airport
	for _, airport := range airports {
		if len(airport.connections) == 0 {
			end = airport
		} else if start == nil || len(start.connections) < len(airport.connections) {
			start = airport
		}
	}

	// build the final flight path
	path := []string{start.name}
	currAirport := start
	for currAirport != end {
		nextAirport := ""
		for connection := range currAirport.connections {
			if airports[connection].name != path[len(path)-1] {
				nextAirport = connection
				break
			}
		}
		if nextAirport == "" {
			panic("could not find a valid flight path")
		}
		path = append(path, nextAirport)
		currAirport = airports[nextAirport]
	}

	return path
}

func flightPathHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type request struct {
		Flights [][]string `json:"flights"`
	}

	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	flightPath := calculateFlightPath(req.Flights)

	respJSON, err := json.Marshal(flightPath)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respJSON)
}

func main() {
	http.HandleFunc("/calculate", flightPathHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}