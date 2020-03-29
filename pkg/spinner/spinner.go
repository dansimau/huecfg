package spinner

import (
	"time"
)

var (
	chars   = []string{`/`, `-`, `\`, `|`}
	running bool
	stop    chan struct{}
)

func init() {
	stop = make(chan struct{})
}

func Start() {
	if running {
		return
	}

	running = true

	go func() {
		for {
			for i := 0; i < len(chars); i++ {
				select {
				case <-stop:
					return
				default:
					print("\r")
					print(chars[i])
					time.Sleep(100 * time.Millisecond)
				}
			}
		}
	}()
}

func Stop() {
	if !running {
		return
	}

	stop <- struct{}{}
	print("\r")
	running = false
}
