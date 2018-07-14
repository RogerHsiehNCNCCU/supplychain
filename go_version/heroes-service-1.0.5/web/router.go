package web

import (
    "net/http"

    "github.com/gorilla/mux"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func (app *Application) NewRouter() *mux.Router {

    var routes = Routes{
        Route{
            "Index",
            "GET",
            "/",
            app.Index,
        },
        Route{
            "TodoIndex",
            "GET",
            "/todos",
            app.TodoIndex,
        },
        Route{
            "TodoShow",
            "GET",
            "/todos/{todoId}",
            app.TodoShow,
        },
        Route{
            "InitLedger",
            "GET",
            "/InitLedger",
            app.InitLedger,
        },
        Route{
            "QueryTransactionByKey",
            "GET",
            "/QueryTransactionByKey",
            app.QueryTransactionByKey,
        },
        Route{
            "QueryAllPoCs",
            "GET",
            "/QueryAllPoCs",
            app.QueryAllPoCs,
        },
        Route{
            "UploadData",
            "POST",
            "/UploadData",
            app.UploadData,
        },
    }
    
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
        Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

//    router := mux.NewRouter().StrictSlash(true)
//        router.
//        Methods("GET").
//            Path("/Initledger").
//            Name("InitLedger").
//            Handler(app.InitLedger)
    
    return router
}

//var routes = Routes{
//    Route{
//        "Index",
//        "GET",
//        "/",
//        Index,
//    },
//    Route{
//        "TodoIndex",
//        "GET",
//        "/todos",
//        TodoIndex,
//    },
//    Route{
//        "TodoShow",
//        "GET",
//        "/todos/{todoId}",
//        TodoShow,
//    },
//    Route{
//        "InitLedger",
//        "GET",
//        "/InitLedger",
//        InitLedger,
//    },
//}