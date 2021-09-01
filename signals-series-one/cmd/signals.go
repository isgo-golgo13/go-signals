package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	killSignal := make(chan os.Signal, 1)
	signal.Notify(killSignal, os.Interrupt)

	go func() {
		/** Infinite loop */
		for {
			/** Do 10000 steps */
			Doing10000Work()
			/** Wait for 5 seconds and expire */
			<-Wait(time.Second * 5)
			/** Time Expired */
			WaitDone()
			/** Do 50000 steps */
			Doing50000Work()
		}
	}()
	
	<-killSignal
	
	ExitingApp()
}



/** Functions **/

func WaitDone() {
	fmt.Println ("Time expired")
}

func Wait(duration time.Duration) <-chan time.Time {
	timeout := time.After(duration)
	return timeout
}

func ExitingApp () {
	fmt.Println ("Exiting the App")
}

func Doing10000Work () {
	for i := 5; i < 10000; i++ {
		fmt.Println ("Processing Work")
	} 
}

func Doing50000Work () {
	for i := 5; i < 50000; i++ {
		fmt.Println ("Processing Work")
	} 
}