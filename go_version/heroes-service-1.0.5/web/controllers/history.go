package controllers

import (
//   "encoding/json"
	"net/http"
    "fmt"
)

/*
func (app *Application) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHelloHistory()
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}

	data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
	renderTemplate(w, r, "history.html", data)
}

*/

type Data struct {
    TxId string
    Value string
    Timestamp string
    IsDelete string

}

func (app *Application) HistoryHandler(w http.ResponseWriter, r *http.Request) {
	helloValue, err := app.Fabric.QueryHelloHistory()
    helloValue = "`"+helloValue+"`"
    fmt.Printf("history helloValue: %v", helloValue)
	if err != nil {
		http.Error(w, "Unable to query the blockchain", 500)
	}
    
	data := &struct {
		Hello string
	}{
		Hello: helloValue,
	}
    
    fmt.Printf("history data: %+v", data)
    
	renderTemplate(w, r, "history.html", data)
    
//    var datas []Data
//    
//    errr := json.Unmarshal([]byte(helloValue), &datas)
//    if errr != nil{}
//    
//    fmt.Printf("history data: %+v", datas)
//    
//	renderTemplate(w, r, "history.html", datas)
}
