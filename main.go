package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/tmendonca28/gophersafe"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: defang <filename>")
        return
    }

    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        defangedLine := defang.DefangURL(line)
        fmt.Println(defangedLine)
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
}