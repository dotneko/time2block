# time2block

time2block is a commandline tool to calculate the time between two blocks given a block time

## Prerequisites

Go version 1.20

## Installation

```
git clone https://github.com/dotneko/time2block.git

cd time2block
go build .
```
## Usage

```
time2block [start block] [target block] [block time]
```

Flags:
```
-d, --detail    Show estimated date/time and details
-r, --raw       Show time left in seconds in raw format
-l, --local     Show estimated local date/time until target block reached
-t, --time      Show estimated UTC date/time until target block reached
```

