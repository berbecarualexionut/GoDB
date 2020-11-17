package main

import(
	"fmt"
	"os"
	"strings"
	"log"
	"bufio"
	"net/http"
	"github.com/berbecarualexionut/webserver"
	"log"
)

var Store = make(map[string]string)

func main() {
	http.HandleFunc("/get", webserver.GetHandle)
	http.HandleFunc("/post", webserver.PostHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
