package main

import (
	"fmt"

	"log"

	"github.com/dustinevan/chron"
	"github.com/json-iterator/go"
	"golang.org/x/crypto/ssh"
	"github.com/dustinevan/partner_common/providers/ampadmin"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Times struct {
	T1 chron.Day
	T2 chron.TimeExact
	T3 chron.Time
}

func main() {
	ts := Times{T1: chron.Today(), T2: chron.Now(), T3: chron.ThisYear().AddN(2)}
	bytes, err := json.Marshal(&ts)
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(string(bytes))
}


func DialWithTimeout(network, addr string config ) {
	ssh.Dial()

}