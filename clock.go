package main

import (
	"fmt"

	"github.com/inancgumus/screen"
	"github.com/jwnpoh/funks"
	"math/rand"
	"github.com/logrusorgru/aurora"
	"time"
)

const (
	end = `
🮔🮔🮔🮔🮔🮔🮔🮔🮔  🮔🮔🮔🮔🮔🮔🮔🮔🮔  🮔🮔🮔     🮔🮔🮔  🮔🮔🮔🮔🮔🮔🮔🮔🮔   🮔🮔🮔🮔     🮔🮔🮔🮔🮔🮔     🮔🮔     🮔🮔  🮔🮔🮔🮔🮔🮔🮔      🮔🮔  
    🮔🮔         🮔🮔     🮔🮔 🮔🮔 🮔🮔 🮔🮔  🮔🮔            🮔🮔   🮔🮔🮔          🮔🮔     🮔🮔  🮔🮔     🮔🮔    🮔🮔  
    🮔🮔         🮔🮔     🮔🮔 🮔🮔 🮔🮔 🮔🮔  🮔🮔            🮔🮔   🮔🮔🮔          🮔🮔     🮔🮔  🮔🮔     🮔🮔    🮔🮔  
    🮔🮔         🮔🮔     🮔🮔 🮔🮔 🮔🮔 🮔🮔  🮔🮔                 🮔🮔🮔          🮔🮔     🮔🮔  🮔🮔     🮔🮔    🮔🮔  
    🮔🮔         🮔🮔     🮔🮔   🮔🮔  🮔🮔  🮔🮔🮔🮔🮔🮔               🮔🮔🮔🮔🮔      🮔🮔     🮔🮔  🮔🮔🮔🮔🮔🮔🮔      🮔🮔  
    🮔🮔         🮔🮔     🮔🮔       🮔🮔  🮔🮔                      🮔🮔🮔     🮔🮔     🮔🮔  🮔🮔               
    🮔🮔         🮔🮔     🮔🮔       🮔🮔  🮔🮔                      🮔🮔🮔     🮔🮔     🮔🮔  🮔🮔               
    🮔🮔         🮔🮔     🮔🮔       🮔🮔  🮔🮔                      🮔🮔🮔     🮔🮔     🮔🮔  🮔🮔               
    🮔🮔     🮔🮔🮔🮔🮔🮔🮔🮔🮔  🮔🮔       🮔🮔  🮔🮔🮔🮔🮔🮔🮔🮔🮔          🮔🮔🮔🮔🮔🮔🮔🮔     🮔🮔🮔🮔🮔🮔🮔🮔🮔  🮔🮔           🮔🮔  
`
)

func endMsg() {
	rand.Seed(time.Now().UnixNano())

	//fmt.Println("")
	screen.Clear()
	for {
		screen.MoveTopLeft()
		r := rand.Intn((7))
		switch r {
		case 1:
			fmt.Print(aurora.Sprintf(aurora.Red(end)))
		case 2:
			fmt.Print(aurora.Sprintf(aurora.Green(end)))
		case 3:
			fmt.Print(aurora.Sprintf(aurora.Blue(end)))
		case 4:
			fmt.Print(aurora.Sprintf(aurora.Yellow(end)))
		case 5:
			fmt.Print(aurora.Sprintf(aurora.Magenta(end)))
		case 6:
			fmt.Print(aurora.Sprintf(aurora.Cyan(end)))
		case 0:
			fmt.Print(aurora.Sprintf(aurora.White(end)))
		}
		time.Sleep(250 * time.Millisecond)
	}
}

func printClock(d time.Duration) {
    h := int64(d.Hours()) / 10
    hh := int64(d.Hours()) % 10
    m := (int64(d.Minutes()) - (int64(d.Hours()) * 60)) / 10
    mm := int64(d.Minutes()) % 10
    s := (int64(d.Seconds()) - (int64(d.Minutes()) * 60)) / 10
    ss := int64(d.Seconds()) % 10

    digits := map[int64]funks.Digit{
        0: funks.Zero,
        1: funks.One,
        2: funks.Two,
        3: funks.Three,
        4: funks.Four,
        5: funks.Five,
        6: funks.Six,
        7: funks.Seven,
        8: funks.Eight,
        9: funks.Nine,        
    }

    var sep funks.Digit
    if ss%2 != 0 {
        sep = funks.Colon
    } else {
        sep = funks.Colonempty
    }

    clock:= [...]funks.Digit{
        digits[h], digits[hh],
        sep,
        digits[m], digits[mm],
        sep,
        digits[s], digits[ss],
        funks.Blank,
    }
    
    screen.Clear()
    screen.MoveTopLeft()

    for i := range clock{
        for j := range clock{
            fmt.Print(clock[j][i], "   ")
        }
        fmt.Println()
    }
}
