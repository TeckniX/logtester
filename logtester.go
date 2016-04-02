package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup(count int) {
	fmt.Printf("Should have had %d log messages\n", count)
}

func main() {

	msPtr := flag.Int("ms", 5000, "milliseconds to wait between logging")
	flag.Parse()

	i := 1
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	go func() {
		<-c
		cleanup(i)
		os.Exit(1)
	}()

	for ; true; i++ {
		t := time.Now()
		fmt.Fprintf(os.Stderr, "%v stderr msg: %d\n", t, i)
		fmt.Fprintf(os.Stdout, "%v stdout msg: %d\n", t, i)
		time.Sleep(time.Duration(*msPtr) * time.Millisecond)
	}

}
