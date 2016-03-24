# Installation
```bash
go install github.com/haochi/who_is_home
```

## Requirements

* Needs `nmap` installed
    * `apt-get install nmap`
    * `brew install nmap`
    * etc ...
* Needs to run with `sudo`

## Use as a library
```bash
go get github.com/haochi/who_is_home
```

## Usage

* Running the binary: `./who_is_home`
* Running with `go run`: `sudo go run main.go`

### Options
```bash
  -file string
    	Known MAC addresses csv file (default "knownMacAddresses.csv")
  -network string
    	Network (default "192.168.0.0/24")
  -nmap string
    	Location of nmap (default "nmap")
  -tries int
    	Number of runs (default 1)
```

The file needs to be a CSV file with owner name, device name, and the device's MAC address. See `knownMacAddresses.csv.example` for an example.
