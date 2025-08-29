package command

import (
	"fmt"
	"time"
)

func HandlerAggregate(s *State, time_between_reqs string) error {
	timeDuration, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return err
	}

	fmt.Printf("\nCollecting feeds every %v\n\n", timeDuration)
	fmt.Printf("——————————————————————————\n\n")

	ticker := time.NewTicker(timeDuration)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}
