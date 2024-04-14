package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const (
	DETAIL  = "detail"
	FULL    = "full"
	LOCAL   = "local"
	RAW     = "raw"
	TIME    = "time"
	SECONDS = "seconds"
	BINARY  = "time2block"
	VERSION = "v0.1"
)

var usage = BINARY + ` is a commandline tool to calculate the time between two blocks given a block time
Usage:
    ` + BINARY + `  [start block] [target block] [block time]
Flags:
	-d, --detail    Show estimated date/time and details
	-r, --raw       Show time left in seconds in raw format
	-l, --local     Show estimated local date/time until target block reached
	-t, --time      Show estimated UTC date/time until target block reached

	-h, --help      Show this message
	-v, --version   Show version information
`[1:]

func help() {
	fmt.Print(usage)
	os.Exit(0)
}

func version() {
	fmt.Printf("%s - %s\n", BINARY, VERSION)
	os.Exit(0)
}

func main() {

	var (
		args       []string
		displayFmt string
		blocksLeft int
		seconds    float64
	)
	// Check flags
	for idx, f := range os.Args {
		switch f {
		case "help", "-h", "-help", "--help":
			help()
		case "-v", "--version":
			version()
		case "-d", "--detail":
			args = append(os.Args[1:idx], os.Args[idx+1:]...)
			displayFmt = DETAIL
		case "-l", "--local":
			args = append(os.Args[1:idx], os.Args[idx+1:]...)
			displayFmt = LOCAL
		case "-r", "--raw":
			args = append(os.Args[1:idx], os.Args[idx+1:]...)
			displayFmt = RAW
		case "-t", "--time":
			args = append(os.Args[1:idx], os.Args[idx+1:]...)
			displayFmt = TIME
		default:
			args = os.Args[1:]
			displayFmt = FULL
		}
	}

	startBlock, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	targetBlock, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	blockTime, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	blocksLeft = targetBlock - startBlock
	seconds = float64(blocksLeft) * blockTime
	n := int(math.Floor(seconds))
	secs := n % 60
	n /= 60
	mins := n % 60
	n /= 60
	hours := n % 60
	n /= 24
	if displayFmt == DETAIL {
		currentTime := time.Now()
		fmt.Printf("Blocks remaining until block %d: %d, Block time: %v\n", targetBlock, blocksLeft, blockTime)
		fmt.Printf("Estimated time left  : %d days %02d:%02d:%02ds\n", n, hours, mins, secs)
		fmt.Printf("UTC Time now         : %v\n", currentTime.UTC())
		fmt.Printf("Estimated target time: %v\n", currentTime.Add(time.Duration(seconds)*time.Second).UTC())
		fmt.Printf("Local Time now       : %v\n", currentTime.Local())
		fmt.Printf("Estimated target time: %v\n", currentTime.Add(time.Duration(seconds)*time.Second).Local())
	} else if displayFmt == LOCAL {
		currentTime := time.Now()
		fmt.Printf("Estimated target time: %v\n", currentTime.Add(time.Duration(seconds)*time.Second).Local())
	} else if displayFmt == RAW {
		fmt.Printf("%v", seconds)
	} else if displayFmt == TIME {
		currentTime := time.Now()
		fmt.Printf("Estimated target time: %v\n", currentTime.Add(time.Duration(seconds)*time.Second).UTC())
	} else {
		// displayFmt == FULL
		fmt.Printf("Blocks remaining     : %d (Target %d, avg block time: %v)\n", blocksLeft, targetBlock, blockTime)
		fmt.Printf("Estimated target time: %dd %02dh %02dm %02ds (%vs) \n", n, hours, mins, secs, seconds)
	}

}
