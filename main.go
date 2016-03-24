package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const macPrefix = "MAC Address: "
const macPrefixLength = len(macPrefix)
const macLength = 17
const macEndingIndex = macPrefixLength + macLength

type macRecord struct {
	owner string
	name  string
}

func main() {
	var networkAddress = flag.String("network", "192.168.0.0/24", "Network")
	var knownMacFile = flag.String("file", "knownMacAddresses.csv", "Known MAC addresses csv file")
	var tries = flag.Int("tries", 1, "Number of runs")
	var nmapPath = flag.String("nmap", "nmap", "Location of nmap")
	flag.Parse()

	uniqueOwners := make(map[string]bool)
	knownMacs, err := readKnownMacs(*knownMacFile)

	if err != nil {
		log.Println(err)
		return
	}

	for i := 0; i < *tries; i++ {
		macAddresses, err := getMacAddresses(*nmapPath, *networkAddress)

		if err != nil {
			log.Println(err)
			return
		}

		owners := WhoIsHome(macAddresses, knownMacs)

		for _, owner := range owners {
			uniqueOwners[owner] = true
		}
	}

	for owner := range uniqueOwners {
		fmt.Println(owner)
	}
}

func WhoIsHome(macAddresses []string, knownMacs map[string]*macRecord) []string {
	owners := make(map[string]bool)
	uniqueOwners := make([]string, 0)

	for _, mac := range macAddresses {
		if record, ok := knownMacs[mac]; ok {
			if _, ok := owners[record.owner]; !ok {
				owners[record.owner] = true
			}
		}
	}

	for owner := range owners {
		uniqueOwners = append(uniqueOwners, owner)
	}

	return uniqueOwners
}

func readKnownMacs(file string) (map[string]*macRecord, error) {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	macNames := make(map[string]*macRecord, 0)

	for _, record := range records {
		owner, name, mac := record[0], record[1], record[2]
		macNames[mac] = &macRecord{
			owner: owner,
			name:  name,
		}
	}
	return macNames, nil
}

func getMacAddresses(nmapPath, network string) ([]string, error) {
	macs := make([]string, 0)
	nmapResult, err := exec.Command(nmapPath, "-sn", network).Output()

	if err != nil {
		return macs, errors.New(fmt.Sprintf("%s: %s", "Error running nmap", err.Error()))
	}

	lines := strings.Split(string(nmapResult), "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, macPrefix) {
			mac := line[macPrefixLength:macEndingIndex]
			macs = append(macs, mac)
		}
	}

	return macs, nil
}
