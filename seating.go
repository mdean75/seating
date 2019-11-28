package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"seating/api"
	"strconv"
	"time"
)

type data struct {
	Industries []string
	Attendees  []Attendee
	Pairs      []pair
}

type Attendee struct {
	name       string
	id         int
	industry   string
	pairedWith []int
}

type pair struct {
	seat1 Attendee
	seat2 Attendee
}

func main() {

	var a api.AppData
	a.Industries = api.SetIndustries()
	r := mux.NewRouter()

	r.HandleFunc("/", a.AttendeeEntry).Methods(http.MethodGet)
	r.HandleFunc("/", a.ProcessAttendeeEntry).Methods(http.MethodPost)
	r.HandleFunc("/attendees", a.DisplayAttendees).Methods(http.MethodGet)
	r.HandleFunc("/seating", a.BuildChart).Methods(http.MethodGet)
	r.HandleFunc("/reset-attendees", a.ResetData).Methods(http.MethodGet)

	r.HandleFunc("/demo", a.Demo).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:         "0.0.0.0:3000",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Println("api listening on port " + srv.Addr)
	log.Fatal(srv.ListenAndServe())

}

func (d *data) setIndustries(indust *[]string) {
	*indust = append(*indust, "Accountants & Tax Preparation",
		"Advertising",
		"Advertising: Direct Mail",
		"Advertising: Promotional Products",
		"Aesthestics",
		"Air Duct Cleaning & Chimney Sweep",
		"Alterations",
		"Ambulance",
		"Apartment Complexes",
		"Appliances: Sales/Service/Parts",
		"Art Studio",
		"Automobile: Body & Dent Repair",
		"Automobile: Detailing",
		"Automobile: Sales",
		"Automobile: Services & Repair",
		"Bakery",
		"Banks",
		"Banquet Facilities",
		"Barber Shop",
		"Beverage Distributors",
		"Bookkeeping Services",
		"BRC: Business Attorney",
		"BRC: CPA",
		"Brewery",
		"Bridal Shop",
		"Building Materials",
		"Business Development",
		"Business Emergency Planning",
		"Business Networking Organization",
		"Business Services",
		"Car Wash",
		"Career Coaching",
		"Catering Services",
		"Chamber of Commerce",
		"Cheerleading",
		"Child Care & Preschools",
		"Chiropractors",
		"Chocolate & Gifts",
		"Church",
		"Civic Organizations",
		"Cleaning Services: Commercial",
		"Cleaning Services: Residential",
		"Coffee House",
		"Community",
		"Computer: IT/Service/Security",
		"Concrete Work: Residential & Commercial",
		"Construction Supply",
		"Construction: Commercial",
		"Construction: Residential",
		"Consultants",
		"Consumer Lending",
		"Counseling Services",
		"Credit Card Processing",
		"Credit Unions",
		"Dance & Gymnastics",
		"Debt Counseling & Repair",
		"Dental Health",
		"Digital Media",
		"Document Management",
		"Document Shredding",
		"Education: Colleges",
		"Education: Music",
		"Education: Public & Private Schools",
		"Educational Services",
		"Elected Officials",
		"Electrical",
		"Embroidery",
		"Emergency Response & Recovery Planning",
		"Employment Agency/Service",
		"Engineering Services",
		"Entertainment",
		"Equipment Rental",
		"Event Planner",
		"Executive Collaboration Suites",
		"Financial Services",
		"Fire Protection",
		"Fitness Club & Gym",
		"Flooring",
		"Florist",
		"Food and Beverage Supply",
		"Food Truck",
		"Funeral Homes",
		"Furniture",
		"Furniture: Outdoor",
		"General Contracting",
		"Glass & Window Repair",
		"Golf Course",
		"Graphic Design",
		"Hardware Stores",
		"Health & Wellness",
		"Healthcare Services",
		"Heating & Air Services",
		"Home Decor & Accessories",
		"Home Health Agencies",
		"Home Improvement",
		"Home Inspections",
		"Hospitals",
		"Hotels & Motels",
		"Human Resource Services",
		"HWP: Grocery",
		"In-Home Podiatry",
		"Individual",
		"Industrial Supplies",
		"Insurance Broker: Chamber Benefit Plan",
		"Insurance Services",
		"Insurance Services: Commercial",
		"Internet Marketing",
		"IT",
		"IT: Back Up/Recovery/Security",
		"Jewelry",
		"Kitchen and Bath",
		"Landscaping & Lawn Service",
		"Landscaping & Lawn Service: Commercial",
		"Leadership Development & Coaching",
		"Legal Services",
		"Life Coach",
		"Locksmith",
		"Mailing Services",
		"Manufacturing",
		"Marketing: Mobile/On-line",
		"Marketing: Sales Promotions",
		"Marketing: Videography/Photography",
		"Martial Arts Academy",
		"Massage Therapy",
		"Mattress Store",
		"Meat Market",
		"Media",
		"Medicare",
		"Memory Care Unit",
		"Mental Health Services",
		"Mold Remediation",
		"Mortgage Services",
		"Moving & Storage",
		"Newspaper & Magazines",
		"Non-Profit Organization",
		"Nutritional Supplement",
		"Occupational Therapy",
		"Office Equipment & Copiers",
		"Optometrists & Ophthamologists",
		"Pain Management",
		"Painting & Supplies",
		"Party Rentals & Inflatables",
		"Payroll Services",
		"Personal Trainer & Nutrition Counseling",
		"Pest Control",
		"Pet Care",
		"Pharmacy",
		"Photo Restoration",
		"Photographers",
		"Physical Therapy",
		"Physicians",
		"Picture Framing",
		"Plumbing Services",
		"Printers & Publishers",
		"Professional Services",
		"Property Management",
		"Radio Station",
		"Real Estate: Commercial",
		"Real Estate: Residential",
		"Recreation & Sports",
		"Recycling",
		"Restaurants",
		"Restaurants: Bar & Grill",
		"Restaurants: Fast Food",
		"Restaurants: Frozen Desserts",
		"Restoration: Fire & Flood",
		"Retail Shopping",
		"Roofing",
		"RV Sales & Repair",
		"Salon & Spa",
		"Screen Printing",
		"Screen Repair",
		"Sealing",
		"Security Services",
		"Self Storage",
		"Senior Living: Assisted",
		"Senior Living: Independent",
		"Senior Living: Skilled Nursing",
		"Senior Services",
		"Service Organization",
		"Services for the Disabled",
		"Siding & Exteriors",
		"Sign Manufacturer",
		"Smoke Shop",
		"Tanning Salon",
		"Technology",
		"Telecommunications Services",
		"Text Message Marketing",
		"Title Companies",
		"Trampoline Park",
		"Travel Services",
		"Truck Services",
		"Trucking Company",
		"Urgent Care",
		"Utilities",
		"Veterinarian",
		"Web Site Design",
		"Website-Video",
		"Wholesale Clubs",
		"Wholesale Distributor",
		"Window Treatments",
		"Wine Bar",
		"Winery")
}

func (d *data) processInput(action string) {
	switch action {
	case "1":
		d.addAttendee()

	case "2":
		d.displayAttendeeList()

	case "3":
		d.buildChart()
	}
}

func (d *data) addAttendee() {
	var att Attendee

	fmt.Println("Add Attendee")
	fmt.Println("\nName: \t\t")
	att.name = readInput()
	d.displayIndustries()
	att.industry = d.getIndustry()

	att.id = randomInt(1, 1000)
	d.Attendees = append(d.Attendees, att)
}

func (d *data) getIndustry() string {
	input := readInput()
	i, _ := strconv.Atoi(input)
	return d.Industries[i]
}

func (d *data) displayIndustries() {
	fmt.Println("Select Attendees' Industry")
	for k, v := range d.Industries {
		fmt.Println(k, ": ", v)
	}
}

func (d *data) displayAttendeeList() {
	for _, a := range d.Attendees {
		fmt.Println(a)
	}
}

func (d *data) buildChart() {
	// ensure an even pairing
	if len(d.Attendees)%2 != 0 {
		d.Attendees = append(d.Attendees, Attendee{name: "Placeholder"})
	}

	c := make([]Attendee, len(d.Attendees))
	copy(c, d.Attendees)

	for ok := true; ok; ok = len(d.Attendees) > 2 {
		var p pair

		p.seat1 = d.shiftArray()
		p.seat2 = d.selectPartner(p.seat1, c)

		d.Pairs = append(d.Pairs, p)
		addAttendeePairing(p.seat1.id, p.seat2.id, c)
	}

	// select last two no matter the match
	lastPair1 := d.shiftArray()
	lastPair2 := d.shiftArray()
	d.Pairs = append(d.Pairs, pair{
		seat1: lastPair1,
		seat2: lastPair2,
	})

	addAttendeePairing(lastPair1.id, lastPair2.id, c)

	//fmt.Println(d.Pairs)
	//fmt.Println("\n\n", c)
	printPairs(d.Pairs)

	// the slice should be nil at this point, reload original data
	d.Attendees = c
}

func printPairs(pairs []pair) {
	for _, v := range pairs {
		fmt.Println(v.seat1.name + ", " + v.seat2.name)
	}
	fmt.Print("\n\n")
}

func addAttendeePairing(seat1 int, seat2 int, list []Attendee) {
	for i := range list {
		if list[i].id == seat1 {
			list[i].pairedWith = append(list[i].pairedWith, seat2)
		}
		if list[i].id == seat2 {
			list[i].pairedWith = append(list[i].pairedWith, seat1)
		}
	}
}

func (d *data) shiftArray() Attendee {
	t := d.Attendees[0]
	d.Attendees = d.Attendees[1:]

	return t
}

func (d *data) selectPartner(seat1 Attendee, c []Attendee) Attendee {
	var seat2 Attendee
	//i := randomInt(0, len(d.Attendees))
	//seat2 = d.peek(i)
	for {
		i := randomInt(0, len(d.Attendees))
		seat2 = d.peek(i)

		if seat2.industry != seat1.industry {
			if seat1.pairedWith == nil {
				// remove from slice - swap to end and reslice
				d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
				d.Attendees = d.Attendees[:len(d.Attendees)-1]

				return seat2
			}
			for _, v := range c {
				if v.id == seat2.id {
					if !arrayContains(seat1.id, seat2.pairedWith) {
						// these two have not been paired before
						// remove from slice - swap to end and reslice
						d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
						d.Attendees = d.Attendees[:len(d.Attendees)-1]

						return seat2
					}

				}
			}

		}
	}
	//if seat2.industry != seat1.industry {
	//	// remove from slice - swap to end and reslice
	//	d.Attendees[i] = d.Attendees[len(d.Attendees)-1]
	//	d.Attendees = d.Attendees[:len(d.Attendees)-1]
	//}
	return seat2

}

func arrayContains(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func (d *data) peek(i int) Attendee {
	return d.Attendees[i]
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func displayMenu() {
	fmt.Println("\n\nSelect a menu option")
	fmt.Println("\nEnter Attendee \t\t\t ... 1")
	fmt.Println("\nDisplay list of Attendees \t ... 2")
	fmt.Println("\nCreate seating charts \t\t ... 3")
	fmt.Println("\nQuit \t\t\t\t ... Q")
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
