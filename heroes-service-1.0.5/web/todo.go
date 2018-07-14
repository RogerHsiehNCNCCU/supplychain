package web

import "time"

type Todo struct {
    Name      string    `json:"name"`
    Completed bool      `json:"completed"`
    Due       time.Time `json:"due"`
}

type Todos []Todo

type UploadDataInputs struct {
    Name      string    `json:"name"`
    Efficacy  string    `json:"headache"`
    Color     string    `json:"green"`
    Owner     string    `json:"POC0"`
}
