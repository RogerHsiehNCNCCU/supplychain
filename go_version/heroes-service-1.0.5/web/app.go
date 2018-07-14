package web

import (
	"fmt"
	"github.com/chainHero/heroes-service/web/controllers"
	"net/http"
    
    "log"
)

func Serve(app *controllers.Application) {
	fs := http.FileServer(http.Dir("web/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	http.HandleFunc("/home.html", app.HomeHandler)
    http.HandleFunc("/history.html", app.HistoryHandler)
    //http.HandleFunc("/historyjson.html", app.HistoryJsonHandler)
	http.HandleFunc("/request.html", app.RequestHandler)

    
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home.html", http.StatusTemporaryRedirect)
	})

	fmt.Println("Listening (http://localhost:3000/) ...")
	http.ListenAndServe(":3000", nil)
}

//要大寫才會是exported狀態
func MyServe(app *Application) {

    router := app.NewRouter()

    fmt.Println("Listening (http://localhost:3000/) ...")
    log.Fatal(http.ListenAndServe(":3000", router))
}