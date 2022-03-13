package main
 
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"math/rand"
	"time"
)


func main() {
	//sets seed for random number generator based on current time
	rand.Seed(time.Now().UnixNano())
	
	//sets random number as variable v
	var v int = rand.Intn(51-1) + 1
	
	//prints the random number
	//fmt.Println(v)
	
	//prints the file
	if line, err := rsl("quotes.txt", v); err == nil {
		fmt.Print(`                      __                         __        ._._._.
  ____   ____ _______/  |_    ________ __  _____/  |_  ____| | | |
 /    \ /  _ \\____ \   __\  / ____/  |  \/  _ \   __\/ __ \ | | |
|   |  (  <_> )  |_> >  |   < <_|  |  |  (  <_> )  | \  ___/\|\|\|
|___|  /\____/|   __/|__|    \__   |____/ \____/|__|  \___  >_____
     \/       |__|              |__|                      \/\/\/\/`)

		fmt.Println("")
		fmt.Println(line)
	} else {
		fmt.Println("rsl:", err)
	}
}
//errors and shit
func rsl(fn string, n int) (string, error) {
	if n < 1 {
		return "", fmt.Errorf("invalid request: line %d", n)
	}
	f, err := os.Open(fn)
	if err != nil {
		return "", err
	}
	defer f.Close()
	bf := bufio.NewReader(f)
	var line string
	for lnum := 0; lnum < n; lnum++ {
		line, err = bf.ReadString('\n')
		if err == io.EOF {
			switch lnum {
			case 0:
				return "", errors.New("no lines in file")
			case 1:
				return "", errors.New("only 1 line")
			default:
				return "", fmt.Errorf("only %d lines", lnum)
			}
		}
		if err != nil {
			return "", err
		}
	}
	if line == "" {
		return "", fmt.Errorf("line %d empty", n)
	}
	return line, nil
}
