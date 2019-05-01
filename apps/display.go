package main

import (
	"github.com/Tikal/distributionWorkshop/messaging"
	"github.com/Tikal/distributionWorkshop/models"
	"encoding/json"
	"net/http"
	"fmt"
)

func main(){
	defer messaging.Stop()

	fmt.Println(http.Dir("display_client"))

	displayInput := getDisplayChannel()

	gameField := models.NewGameField()

	go func(){
		http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request) {
			gfSer, err := json.Marshal(gameField)
			if err != nil {
				// TODO handle gracefully
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(gfSer)
		})
		fs := http.FileServer(http.Dir("display_client"))
		http.Handle("/client/", http.StripPrefix("/client", fs))

		http.ListenAndServe(":8080", nil)
	}()

	for ds := range displayInput {
		gameField.Update(ds)
	}
}

func getDisplayChannel() <- chan *models.DisplayStatus  {
	rawInput, _ := messaging.GetInputChannel(messaging.DisplayChannelName)
	res := make(chan *models.DisplayStatus)

	// Display channel population, executed in function closure
	go func(){
		for val := range rawInput {
			ds := &models.DisplayStatus{}
			err := json.Unmarshal(val, ds)
			if err == nil {
				res <- ds
			}
		}
	}()
	return res
}
