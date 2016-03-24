package core

import (
	"flag"
	"log"
	"strconv"
	"time"
)

const (
	dateFmt = "January 2, 2006"
)

func computeParf(omv int, dateOfRegistration string) int {
	t, _ := time.ParseInLocation(dateFmt, dateOfRegistration, time.UTC)

	cutOfDate, _ := time.ParseInLocation(dateFmt, "January 1, 2008", time.UTC)
	if cutOfDate.After(t) {
		return int(omv * 55 / 100)
	}

	return int(omv * 50 / 100)
}

func computeEMI(monthsLeft int, amount int, rateOfInterest float32) (emi, downPayment, loanAmount int) {

	downPayment = int(amount * 40 / 100)
	loanAmount = amount - downPayment

	interestPm := int((float32(loanAmount) * (rateOfInterest / 100)) / 12)
	totalPayable := int(loanAmount + (interestPm * monthsLeft))

	emi = totalPayable / monthsLeft

	return
}

func diff(t1, t2 time.Time) (years, months, days int) {
	t2 = t2.AddDate(0, 0, 1) // advance t2 to make the range inclusive

	for t1.AddDate(years, 0, 0).Before(t2) {
		years++
	}
	years--

	for t1.AddDate(years, months, 0).Before(t2) {
		months++
	}
	months--

	for t1.AddDate(years, months, days).Before(t2) {
		days++
	}
	days--

	return
}

func computeMonthsLeft(dateOfRegistration string, todaysDate string) int {
	t0, _ := time.ParseInLocation(dateFmt, todaysDate, time.UTC)
	t1, _ := time.ParseInLocation(dateFmt, dateOfRegistration, time.UTC)

	y, m, _ := diff(t1, t0)
	totalElapsedMonths := (y * 12) + m

	return 120 - totalElapsedMonths
}

func computeDepriciation(cost int, parf int, monthsLeft int) int {
	effectiveValue := cost - parf
	return effectiveValue / monthsLeft
}

func main() {

	dateOfRegisteration := flag.String("d", "", "Enter the vehical's date as \"Jan 2, 2010\" format ")
	todaysDate := flag.String("td", "", "Enter the todays's date as \"Jan 2, 2010\" format ")
	omvS := flag.String("omv", "", "Enter the vehical's omv ")
	priceS := flag.String("price", "", "Enter the vehical's price ")
	roiS := flag.String("roi", "", "Enter the bank's roi ")

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
	roi, _ := strconv.ParseFloat(*roiS, 32)

	monthsLeft := computeMonthsLeft(*dateOfRegisteration, *todaysDate)

	log.Println("Number of Months Left in the Car ", monthsLeft)

	parf := computeParf(int(omv), *dateOfRegisteration)
	emi, down, loan := computeEMI(monthsLeft, int(price), float32(roi))
	dep := computeDepriciation((down + loan), parf, monthsLeft)

	log.Println("Duration: ", monthsLeft, " Get Back: ", parf, " Emi: ", emi, " Dep: ", dep)
}
