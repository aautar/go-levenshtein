package main

import "fmt"
import "bufio"
import "os"
import "levenshtein"

func main() {

    consoleReader := bufio.NewReader(os.Stdin)
    fmt.Print("What word are you looking for? ")

    input, err := consoleReader.ReadString('\n')
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    // From http://www.cl.cam.ac.uk/~mgk25/ucs/examples/quickbrown.txt
    // Translates to: "Wrongful practicing of xylophone music tortures every larger dwarf"
    xylophoneTortureWords := [8]string{"Falsches", "Üben", "von", "Xylophonmusik", "quält", "jeden", "größeren", "Zwerg"}

    minDist := int(^uint(0) >> 1)
    minDistIndex := -1
    for i:=0; i<len(xylophoneTortureWords); i++ {
        dist := levenshtein.Distance(xylophoneTortureWords[i], input)
        if(dist < minDist) {
            minDist = dist
            minDistIndex = i
        }
    }

    if minDist == 0 {
        fmt.Println("WORD FOUND: " + xylophoneTortureWords[minDistIndex])
    } else {
        fmt.Println("Did you mean " + xylophoneTortureWords[minDistIndex] + " ?")
    }

}
