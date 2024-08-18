package main

import (
	"bufio"
	"fmt"
	"os"
	"github.com/tmendonca28/gophersafe/defang"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: gophersafe <filename>")
        return
    }
    
    // Opening input file
    file, err := os.Open(os.Args[1])
    if err != nil {
        fmt.Println("Error opening file", err)
        return
    }
    defer file.Close()

    // Create output file
    outputFile, err := os.Create("defanged_" + os.Args[1])
    if err != nil {
        fmt.Println("Error creating output file", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(outputFile)
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        // Defang line
        defangedLine := defang.DefangURL(line)
        writer.WriteString(defangedLine + "\n")
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error reading file:", err)
    }
    writer.Flush()

}