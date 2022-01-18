package GoTicker

import (
	"fmt"
	"time"
)

func Run() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("get!")
		}

	}
}
