package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func main() {
    db := database{"shoes": 50, "socks": 5}
    mux := http.NewServeMux()
    mux.HandleFunc("/list", db.list)
    mux.HandleFunc("/price", db.price)
    mux.HandleFunc("/update", db.update)
    log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func (db database) create(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    price := req.URL.Query().Get("price")
    if _, ok := db[item]; ok {
        newPrice, err := strconv.ParseFloat(price, 32)
        if err != nil {
            panic(err)
        }
        if newPrice < 0 {
            fmt.Fprintf(w, "price set fail: %q\n", newPrice)
            return
        }
        db[item] = dollars(float32(newPrice))
    }

    fmt.Fprintf(w, "update success %s: %s\n", item, price)
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
    for item, price := range db {
        fmt.Fprintf(w, "%s: %s\n", item, price)
    }
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
    item := req.URL.Query().Get("item")
    price, ok := db[item]
    if !ok {
        w.WriteHeader(http.StatusNotFound) // 404
        fmt.Fprintf(w, "no such item: %q\n", item)
        return
    }
    fmt.Fprintf(w, "%s\n", price)
}
