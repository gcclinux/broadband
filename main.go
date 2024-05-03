package main

import (
	"fmt"
	"gcclinux/broadband/handler"
)

func main() {

	l, d, u, i := handler.GetBroadbandSpeed()
	w := handler.GetOutboundIP()
	c := handler.GetConfig()

	if c.Verbose[0] {
		fmt.Printf("\nWanIP: %s, Latency: %f, Download: %f, Upload: %f, (%s)\n", w, l, d, u, i)
	}

	if c.SaveToDB[0] {
		err := handler.SaveStaticAddress(w, l, d, u, i)
		if err != nil {
			fmt.Println("Error in saving data to database:", err)
		}
	}

}
