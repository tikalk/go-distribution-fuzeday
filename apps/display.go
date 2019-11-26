package apps

import (
	"encoding/json"
	"go-distribution-fuzeday/messaging"
	"go-distribution-fuzeday/models"
	"net/http"
	"sync"
)

//var GlobalDisplayChannel = make(chan *models.DisplayStatus, 1000)
//var displayInputChannel = make(chan *models.DisplayStatus, 1000)
var displayInputChannel chan *models.DisplayStatus = nil

//func getDisplay(w http.ResponseWriter, r *http.Request) {
//	//message := r.URL.Path
//	//message = strings.TrimPrefix(message, "/")
//	//message = "Hello " + message
//
//	message := <- getDisplayInputChannel()
//
//	w.Write([]byte(message))
//}

func LaunchDisplay(port int, externalWaitGroup *sync.WaitGroup) {

	displayInput := getDisplayInputChannel()

	gameField := models.NewGameField()

	// HTTP Server
	//TODO Challenge (4):
	//	1. launch HTTP server here on 8080
	//	2. requests to "/display" should return a json representation of the updated gameField
	//	3. requests to "/client/" should return static files from directory "display_client". Use http.FileServer...
	// 	------
	// 	Tip: use http.HandleFunc and http.ListenAndServe
	http.HandleFunc("/display", func(w http.ResponseWriter, r *http.Request){
		buf, _ := json.Marshal(gameField)
		w.Write(buf)
	})
	fs := http.FileServer(http.Dir("display_client"))
	http.Handle("/client/", http.StripPrefix("/client", fs))
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			panic(err)
		}
	}()

	// Game Field updater
	//TODO Challenge (4):
	//	1. iterate over display channel
	//	2. update gamefield on each consumed value
	//	------
	//	Tip: use iteration over channel range
	for {
		display := <- displayInput
		gameField.Update(display)
	}

	//displayInput = displayInput // only to prevent "unused variable error", remove after implementation
	//gameField = gameField       // only to prevent "unused variable error", remove after implementation

	if externalWaitGroup != nil {
		externalWaitGroup.Done()
	}
}

func getDisplayInputChannel() chan *models.DisplayStatus {
	//TODO Challenge (2):
	//  get []byte input channel from messaging,
	//  create an internal goroutine that consumes messages from it,
	//  de-serialize them to return type and populates return DIRECTIONAL channel
	//return GlobalDisplayChannel

	if displayInputChannel == nil {
		displayInputChannel = make(chan *models.DisplayStatus, 1000)

		go func() {
			for {
				// Read from messaging
				var displayStatus *models.DisplayStatus = nil
				response := <- messaging.GetInputChannel("display")
				json.Unmarshal(response, &displayStatus)
				displayInputChannel <- displayStatus
			}
		}()
	}

	return displayInputChannel
}
