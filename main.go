
package main
// import package 
import (
    "encoding/json"
    "log"
    "math/rand"
    "net/http"
    "strconv"
    "sync"
)

type Item struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Value string `json:"value"`
}

var (
    items = make(map[int]Item)
    mu    sync.Mutex
)

func getItems(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    var list []Item
    for _, v := range items {
        list = append(list, v)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(list)
}

func getItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    item, ok := items[id]
    if !ok {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

func createItem(w http.ResponseWriter, r *http.Request) {
    var item Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    item.ID = rand.Intn(1000000)
    items[item.ID] = item
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    var item Item
    if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    if _, ok := items[id]; !ok {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }
    item.ID = id
    items[id] = item
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
    idStr := r.URL.Query().Get("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }
    mu.Lock()
    defer mu.Unlock()
    if _, ok := items[id]; !ok {
        http.Error(w, "Item not found", http.StatusNotFound)
        return
    }
    delete(items, id)
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getItems(w, r)
        case http.MethodPost:
            createItem(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    http.HandleFunc("/item", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            getItem(w, r)
        case http.MethodPut:
            updateItem(w, r)
        case http.MethodDelete:
            deleteItem(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
