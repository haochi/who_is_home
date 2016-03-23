# Installation
```bash
go install github.com/haochi/who_is_home
```

## Use as a library
```bash
go get github.com/haochi/who_is_home
```

## Usage

```bash
  -file string
    	Known MAC addresses csv file (default "knownMacAddresses.csv")
  -network string
    	Network (default "192.168.0.0/24")
  -tries int
    	Number of runs (default 1)
```

The file needs to be a CSV file with owner name, device name, and the device's MAC address. See `knownMacAddresses.csv.example` for an example.
