package core

import (
	"time"
)

const (
	dateFmt = "January 2, 2006"
)

func ComputeParf(omv int, dateOfRegistration string) int {
	t, _ := time.ParseInLocation(dateFmt, dateOfRegistration, time.UTC)

	cutOfDate, _ := time.ParseInLocation(dateFmt, "January 1, 2008", time.UTC)
	if cutOfDate.After(t) {
		return int(omv * 55 / 100)
	}

	return int(omv * 50 / 100)
}

func ComputeEMI(monthsLeft int, amount int, rateOfInterest float32) (emi, downPayment, loanAmount int) {

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

func ComputeMonthsLeft(dateOfRegistration string, todaysDate string) int {
	t0, _ := time.ParseInLocation(dateFmt, todaysDate, time.UTC)
	t1, _ := time.ParseInLocation(dateFmt, dateOfRegistration, time.UTC)

	y, m, _ := diff(t1, t0)
	totalElapsedMonths := (y * 12) + m

	return 120 - totalElapsedMonths
}

func ComputeDepriciation(cost int, parf int, monthsLeft int) int {
	effectiveValue := cost - parf
	return effectiveValue / monthsLeft
}

func ComputeRoadTax(cc int) float32 {
	var halfYearly float32
	if cc <= 1600 {
		halfYearly = (250.00 + .375*(float32(cc)-1000.00)) * .782
	} else if cc <= 3000 {
		halfYearly = (475.00 + 0.75*(float32(cc)-1600.00)) * 0.782
	} else {
		halfYearly = (1525.00 + (float32(cc) - 1600.00)) * 0.782
	}
	return halfYearly * 2.00
}
