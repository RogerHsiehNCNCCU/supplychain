package web

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/chainHero/heroes-service/blockchain"
    "github.com/gorilla/mux"
    "io/ioutil"
    "io"
)

type Application struct {
	Fabric *blockchain.FabricSetup
}

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome!")
}

func (app *Application) TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func (app *Application) TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func (app *Application) InitLedger(w http.ResponseWriter, r *http.Request) {
    app.Fabric.ChainCodeID = "heroes-service"
    result,err := app.Fabric.InvokeInitLedger()
    fmt.Fprintln(w, result)
    fmt.Fprintln(w, "InitLedger!")
    if err != nil{
        fmt.Fprintln(w, err)
    }
}

func (app *Application) QueryTransactionByKey(w http.ResponseWriter, r *http.Request) {
    result, err := app.Fabric.QueryTransactionByKey("hello")
    if err != nil {
        panic(err)
    }
    fmt.Fprintln(w, "queryTransactionByKey!"+result)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        panic(err)
    }
}

func (app *Application) QueryAllPoCs(w http.ResponseWriter, r *http.Request) {
    result, err := app.Fabric.QueryAllPoCs()
    if err != nil {
        panic(err)
    }
    //fmt.Fprintln(w, "queryAllPoCs! "+result)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        panic(err)
    }
}

func (app *Application) UploadData(w http.ResponseWriter, r *http.Request) {
    var input UploadDataInputs
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &input); err != nil {
        w.Header().Set("Content-Type", "application/json;   charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    
    inputs := [4]string{input.Name, input.Efficacy, input.Color, input.Owner}
    
    result, err := app.Fabric.UploadData(inputs)
    if err != nil {
        panic(err)
    }
    //fmt.Fprintln(w, "queryAllPoCs! "+result)
    if err := json.NewEncoder(w).Encode(result); err != nil {
        panic(err)
    }
}