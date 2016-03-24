package main

import (
	"flag"
	core "github.com/azharaa/carsSingapore/cars/core"
	"log"
	"strconv"
)

func main() {
	dateOfRegisteration := flag.String("d", "", "Enter the vehical's date as \"Jan 2, 2010\" format ")
	todaysDate := flag.String("td", "", "Enter the todays's date as \"Jan 2, 2010\" format ")
	omvS := flag.String("omv", "", "Enter the vehical's omv ")
	priceS := flag.String("price", "", "Enter the vehical's price ")
	roiS := flag.String("roi", "", "Enter the bank's roi ")
	ccS := flag.String("cc", "", "Enter the CC capacity of the Engine")

	flag.Parse()

	if *dateOfRegisteration == "" {
		log.Fatalln("Must enter the Date of Regiseration of the vehical in \"Jan 2, 2010\" format  ")
	}

	if *todaysDate == "" {
		log.Fatalln("Must enter the Todays date in \"Jan 2, 2010\" format  ")
	}

	if *omvS == "" {
		log.Fatalln("Must enter the OMV of the vehical")
	}

	if *priceS == "" {
		log.Fatalln("Must enter the price of the vehical")
	}

	if *roiS == "" {
		log.Fatalln("Must enter bank's Rate of interest ")
	}
	omv, _ := strconv.ParseInt(*omvS, 10, 0)
	price, _ := strconv.ParseInt(*priceS, 10, 0)
	cc, _ := strconv.ParseInt(*ccS, 10, 0)
	roi, _ := strconv.ParseFloat(*roiS, 32)

	monthsLeft := core.ComputeMonthsLeft(*dateOfRegisteration, *todaysDate)

	parf := core.ComputeParf(int(omv), *dateOfRegisteration)
	emi, down, loan := core.ComputeEMI(monthsLeft, int(price), float32(roi))
	dep := core.ComputeDepriciation((down + loan), parf, monthsLeft)
	rt := core.ComputeRoadTax(int(cc))

	log.Printf("Down Payement %d : Get Back  %d : EMI %d : Depericiation %d : Months Left: %d, Road Tax: %f", down, parf, emi, dep, monthsLeft, rt)

}
