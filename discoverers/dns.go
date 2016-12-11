package discoverers

import (
	"fmt"
	"net"
	"sort"
	"strings"
)

type DNS struct{}

func NewDNS() *DNS {
	return &DNS{}
}

func (d *DNS) Discover(s string, ns string) ([]string, error) {
	_, srvRecords, err := net.LookupSRV("", "", strings.Join([]string{s, ns}, "."))
	if err != nil {
		return nil, err
	}

	records := []string{}
	for _, srvRecord := range srvRecords {
		endpoint := fmt.Sprintf("%s", srvRecord.Target[:len(srvRecord.Target)-1])
		records = append(records, endpoint)
	}

	sort.Strings(records)

	return records, nil
}
