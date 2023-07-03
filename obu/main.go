package main

import (
	"log"
	"math"
	"math/rand"
	"time"
"github.com/anurag4way/go-concurrency/types"
	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var sendtimeinterval = time.Second



func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}
func genLatLong() (float64, float64) {

	return genCoord(), genCoord()
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func main() {
	obuIDs := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := types.OBUDATA{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(sendtimeinterval)
	}
}
func init() {
	rand.Seed(time.Now().UnixNano())
}
