package main

import(
	"PRFLR.SDK.GO/PRFLR"
	"time"
	"strconv"
)

func main() {
	//init PRFLR
	PRFLR.Init("prflr://VlgjLdisTrCEKUY1jsLvoMleAD9q09dp@prflr.org:4000", "GO.SDK.Test")

	//Start test loop
	for i := 1; i <= 10; i++ {
		//Start Timer
		timer := PRFLR.New("Test.IT")
		//Do some logic
		time.Sleep(987 * time.Millisecond)
		//Finish
		timer.End("Step:"+strconv.Itoa(i))
	}
}
