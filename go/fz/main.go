package main
import (
    "fmt"
    "strings"
    "os"
)

const (
    cyan = "\x1b[46m"
    reset = "\x1b[0m"
)
func fs(path string, input string) {
    rec := false
    for _, arg := range os.Args {
        if arg == "-r" {
            rec = true
        }
    }

    dirs, err := os.ReadDir(path)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, dir := range dirs {
        if strings.Contains(dir.Name(), input) {
            hl := strings.ReplaceAll(dir.Name(), input, cyan+input+reset)
            if rec {
                fmt.Print(path + "/"
            } if dir.IsDir() {
                fmt.Println(hl + "/")
            } else {
                fmt.Println(hl + "/")
            }
        }
        if rec && dir.IsDir() {
            fs(path+"/"+dir.Name(), input)
        }
    }
}
func main() {
    if os.Args[1] == "-r" {
        fs(os.Args[2], os.Args[3])
    } else {
        fs(os.Args[1], os.Args[2])
    }
}
