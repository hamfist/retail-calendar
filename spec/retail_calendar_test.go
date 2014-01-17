package spec

import (
	"github.com/modcloth-labs/retail-calendar"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RetailCalendar", func() {
	var date = time.Date(2010, 10, 15, 0, 0, 0, 0, time.UTC)
	It("calculates the startTime of the given date's fiscal month", func() {
		startTime, _ := retail_calendar.GetBeginningAndEndTimeForFiscalMonth(date)
		Expect(startTime.Year()).To(Equal(2010))
		Expect(startTime.Month()).To(Equal(time.October))
		Expect(startTime.Day()).To(Equal(10))
	})

	It("calculates the endTime of the given date's fiscal month", func() {
		_, endTime := retail_calendar.GetBeginningAndEndTimeForFiscalMonth(date)
		Expect(endTime.Year()).To(Equal(2010))
		Expect(endTime.Month()).To(Equal(time.November))
		Expect(endTime.Day()).To(Equal(7))
	})

	It("calculates the endTime at midnight", func() {
		_, endTime := retail_calendar.GetBeginningAndEndTimeForFiscalMonth(date)
		Expect(endTime.Hour()).To(Equal(0))
		Expect(endTime.Minute()).To(Equal(0))
		Expect(endTime.Second()).To(Equal(0))
	})

	It("calculates the startTime at midnight", func() {
		_, startTime := retail_calendar.GetBeginningAndEndTimeForFiscalMonth(date)
		Expect(startTime.Hour()).To(Equal(0))
		Expect(startTime.Minute()).To(Equal(0))
		Expect(startTime.Second()).To(Equal(0))
	})
})
