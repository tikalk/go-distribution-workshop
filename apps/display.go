package apps

import (

	"net/http"
	"sync"
	"github.com/tikalk/go-distribution-workshop/messaging"
	"fmt"
)

func LaunchDisplay(port int, externalWaitGroup *sync.WaitGroup){
	defer messaging.Stop()


	go func(){
		// TODO Challenge: serve gameField here, via `/display` end point

		fs := http.FileServer(http.Dir("display_client"))
		http.Handle("/client/", http.StripPrefix("/client", fs))

		http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	}()

	// TODO Challenge: read from display channel and update game field


	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}
}

