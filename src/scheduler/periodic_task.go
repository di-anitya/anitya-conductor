package scheduler

import (
	"fmt"
	"time"
)

func periodicFunction(tick time.Time) {
	fmt.Println("Tick at: ", tick)
}

func main() {
	for t := range time.NewTicker(1 * time.Second).C {
		periodicFunction(t)
	}
}
