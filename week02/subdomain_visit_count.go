package leetcode

import (
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	result := make([]string, 0)
	if len(cpdomains) == 0 {
		return result
	}
	domainCountMap := make(map[string]int, 0)
	for _, domain := range cpdomains {
		countDomain := strings.Split(domain, " ")
		allDomains := strings.Split(countDomain[1], ".")
		temp := make([]string, 0)
		for i := len(allDomains) - 1; i >= 0; i-- {
			temp = append([]string{allDomains[i]}, temp...)
			ld := strings.Join(temp, ".")
			count, _ := strconv.Atoi(countDomain[0])
			if val, ok := domainCountMap[ld]; !ok {
				domainCountMap[ld] = count
			} else {
				domainCountMap[ld] = count + val
			}
		}
	}
	for k, v := range domainCountMap {
		t := strings.Join([]string{strconv.Itoa(v), k}, " ")
		result = append(result, t)
	}
	return result
}
