package main
import (
    "fmt"
    "strings"
    "os"
)
const (
    cyan = "\x1b[36m"
    reset = "\x1b[0m"
)
func fs(path string, input string) {
    dirs, err := os.ReadDir(path)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, dir := range dirs {
        if strings.Contains(dir.Name(), input) {
            hl := strings.ReplaceAll(dir.Name(), input, cyan+input+reset)
            if dir.IsDir() {
                fmt.Println(hl + "/")
            } else {
                fmt.Println(hl)
            }
        }
    }
}
func main() {
    fs(".", os.Args[1])
}
