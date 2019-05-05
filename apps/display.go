package apps

import (

	"encoding/json"
	"net/http"
	"sync"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"github.com/tikalk/go-distribution-workshop/models"
	"fmt"
)

func LaunchDisplay(port int, externalWaitGroup *sync.WaitGroup){

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

		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()

	for ds := range displayInput {
		gameField.Update(ds)
	}

	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}
}

func getDisplayChannel() <- chan *models.DisplayStatus  {
	rawInput := messaging.GetInputChannel(messaging.DisplayChannelName)
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
