package controllers
/*
import (
    "encoding/json"
	"net/http"
)

var helloValue string

func (app *Application) HistoryJsonHandler(w http.ResponseWriter, r *http.Request) {
    var err error
    http.HandleFunc("/historyjson.html", HistoryJsonHandler2)
    helloValue,err = app.Fabric.QueryHelloHistory()
    if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
}

func HistoryJsonHandler2(w http.ResponseWriter, r *http.Request) {
	//helloValue, err := app.Fabric.QueryHelloHistory()
	/*if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}*/

	/*data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
	renderTemplate(w, r, "historyjson.html", data)*/
    
    /*js, err := json.Marshal(helloValue)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
    //renderTemplate(w, r, "historyjson.html", js)
}
*/