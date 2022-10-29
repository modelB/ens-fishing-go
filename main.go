package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	f, err := os.Open("top10milliondomains.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reader := csv.NewReader(f)
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}

	for _, row := range rows {
		time.Sleep(4000000000)
		// fmt.Println(row[1])
		fullDomain := row[1]
		possibleDomains := strings.Split(fullDomain, ".")
		for j, domain := range possibleDomains {
			if j < 2  && len(domain) > 3 {
				ethName := domain + ".eth"
				resp, err := http.Get("https://etherscan.io/enslookup-search?search=" + ethName)
				if err != nil {
					log.Fatalln(err)
				}
				// //We Read the response body on the line below.
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatalln(err)
				}
				//Convert the body to type string
				sb := string(body)

                // fmt.Println(sb)

				if strings.Contains(sb, "is either not registered on ENS") {
					fmt.Println(fullDomain)
					fmt.Println(ethName)
				} else if strings.Contains(sb, "security") {
                    fmt.Println("SECURITY FAIL")
                    os.Exit(1)
                }
			}
		}

	}

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
}
