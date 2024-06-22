package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strconv"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        t := template.New("main")
        t.Funcs(template.FuncMap{
            "pageNum": func(page string) (int, error) {
                pageNumInt, err := strconv.Atoi(page)
                return pageNumInt, err
            },
        })

        switch r.URL.Path {
        case "/":
            mainHandler(w, r, t)
        default:
            http.Error(w, "Page not found", http.StatusNotFound)
        }
    })
    log.Println("Server listening on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func mainHandler(w http.ResponseWriter, r *http.Request, t *template.Template) {
    // Check if a "page" query parameter is present in the URL
    pageNum := r.URL.Query().Get("page")
    if pageNum == "" {
        // Render index. html (start of the story) if no page query param is present
        data := map[string]string{"title": "Start of the Story"}
        err := t.ExecuteTemplate(w, "index.html", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    } else {
        pageNumInt, _ := strconv.Atoi(pageNum)
        data := map[string]string{"title": "Page " + pageNum}
        err := t.ExecuteTemplate(w, "page"+pageNum+".html", data)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
}
