package api

import (
	"bufio"
	"fmt"
	"github.com/rivo/sessions"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"seating/app"
	"strings"
)

// data struct for the add attendee template
type inputData struct {
	Industries   []string
	SuccessMsg   string
	Name         string
	KeyErr       string
	BusinessName string
}

// overall data for the application
type AppData struct {
	Industries []string
	Attendees  []Attendee
	Pairs      []pair
}

type Attendee struct {
	name       string
	id         int
	industry   string
	business   string
	pairedWith []int
}

type pair struct {
	seat1 Attendee
	seat2 Attendee
}

func TestJson(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	var output string
	var inv string
	// read each line, check if there are extra spaces at end of value by calling function, if spaces are
	// present add invalid message to the line
	scanner := bufio.NewScanner(strings.NewReader(string(body)))
	for scanner.Scan() {
		text := scanner.Text()
		b := validateNoSpaces(text)
		if !b {
			inv = inv + text + "   *** INVALID *** \n"
			// split the line, add the message to last element, then join back to a single string
			line := strings.Split(scanner.Text(), "\"")
			length := len(line)
			line[length-2] = line[length-2] + "'   *** INVALID ***"

			j := strings.Join(line, "\"")
			output = output + j
		} else {
			output = output + text
		}
	}

	w.Header().Set("Content-type", "Text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(inv))
}

func validateNoSpaces(s string) bool {
	// convert string to []rune
	sr := []rune(s)

	var extractedRunes []rune // used to hold extracted runes, only chars between ""

	var track bool

	// iterate over the rune slice looking for characters between quotes
	for _, c := range sr {

		if string(c) == "\"" {
			if track == false {
				// start tracking
				track = true

			} else {
				track = false
				// add last \" here because the flag has been switched to not track
				extractedRunes = append(extractedRunes, c)
			}
		}

		if track == true {
			extractedRunes = append(extractedRunes, c)
		}
	}

	// convert back to string, the results may contain more than 1 quoted string
	result := string(extractedRunes)

	// split to get each string we want to evaluate and check if there are any blank spaces
	z := strings.Split(result, "\"")
	for _, substr := range z {
		if len(substr) > 0 {
			if strings.HasPrefix(substr, " ") || strings.HasSuffix(substr, " ") {

				return false
			}
		}

	}
	return true
}

func (a *AppData) ResetData(w http.ResponseWriter, r *http.Request) {

	// start new session and create cookie
	session, _ := sessions.Start(w, r, true)
	session.Set("successMsg", "Meeting attendees have been reset")

	a.Attendees = []Attendee{}

	http.Redirect(w, r, "/", http.StatusFound)
}

// AttendeeEntry is the handler to display the user input form.
func (a *AppData) AttendeeEntry(w http.ResponseWriter, r *http.Request) {

	// start the session but do not create a new session if one does not exist as this could indicate either someone mistakenly routed to this endpoint or some kind of attack.
	session, err := sessions.Start(w, r, false)

	var tData inputData
	// check if there is a session, then get the session message and immediately delete it, if not present set default message
	if session != nil {
		msg := session.GetAndDelete("successMsg", "")
		err = session.Destroy(w, r)
		//if err != nil {
		//	// there was an error deleting the session, redirect to main page
		//	http.Redirect(w, r, "/", http.StatusSeeOther)
		//
		//	log.Error(err.Error())
		//	return
		//}
		tData.SuccessMsg = fmt.Sprintf("%v", msg)

	}

	tData.Industries = a.Industries

	err = loadForm(app.InputForm, w, tData)
	if err != nil {
		//log.Debug("error in AttendeeEntry() received from LoadForm()", err.Error())
	}

}

func (a *AppData) ProcessSecretsForm(w http.ResponseWriter, r *http.Request) {

	// start new session and create cookie
	session, err := sessions.Start(w, r, true)

	err = r.ParseForm()
	if err != nil {
		// do something
	}

	name := r.Form.Get("name")
	business := r.Form.Get("business")
	industry := r.Form.Get("industry")

	attendee := Attendee{
		name:     name,
		business: business,
		id:       randomInt(1, 1000),
		industry: industry,
	}

	a.Attendees = append(a.Attendees, attendee)

	err = session.Set("successMsg", fmt.Sprintf("Added %s to meeting", attendee.name))

	http.Redirect(w, r, "/", http.StatusSeeOther)

	//fmt.Printf("name: %s, business: %s, industry: %s", name, business, industry)
	//var tData inputData
	//tData.Industries = SetIndustries()
	//
	//err = loadForm(app.InputForm, w, tData)
	//if err != nil {
	//	//log.Debug("error in AttendeeEntry() received from LoadForm()", err.Error())
	//}

}

func (a *AppData) DisplayAttendees(w http.ResponseWriter, r *http.Request) {

	var s string
	for _, att := range a.Attendees {
		s = s + att.name + "\t" + att.business + "\n"
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(s))
}

// loadForm is a helper function to handling parsing and displaying the forms.
func loadForm(file string, w http.ResponseWriter, data interface{}) error {

	// parse the template file
	t, err := template.New("html").Parse(file)
	if err != nil {
		log.Println("template parsing error: ", err)
		return err
	}

	// load the form
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template executing error: ", err)
		return err
	}

	return nil

}

func (a *AppData) BuildChart(w http.ResponseWriter, r *http.Request) {
	// add a placeholder member if odd number registered
	if len(a.Attendees)%2 != 0 {
		a.Attendees = append(a.Attendees, Attendee{name: "Placeholder"})
	}

	c := make([]Attendee, len(a.Attendees))
	num := copy(c, a.Attendees)
	fmt.Printf("number copied: %d \n", num)

	var s string
	for i := 0; i < 4; i++ {

		for ok := true; ok; ok = len(a.Attendees) > 2 {
			var p pair

			p.seat1 = a.shiftArray()
			p.seat2 = a.selectPartner(p.seat1, c)

			a.Pairs = append(a.Pairs, p)
			addAttendeePairing(p.seat1.id, p.seat2.id, c)
		}

		// select last two no matter the match
		lastPair1 := a.shiftArray()
		lastPair2 := a.shiftArray()
		a.Pairs = append(a.Pairs, pair{
			seat1: lastPair1,
			seat2: lastPair2,
		})

		addAttendeePairing(lastPair1.id, lastPair2.id, c)

		s += printPairs(a.Pairs)

		a.Pairs = []pair{}
		// the slice should be nil at this point, reload original data
		a.Attendees = make([]Attendee, len(c))
		num = copy(a.Attendees, c)
		fmt.Printf("number copied: %d \n", num)
	}

	a.Pairs = []pair{}

	a.clearSeating()

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write([]byte(s))
}

func (a *AppData) clearSeating() {
	for i := 0; i < len(a.Attendees); i++ {
		a.Attendees[i].pairedWith = nil
		fmt.Println(a.Attendees[i])
	}
}

func printPairs(Pairs []pair) string {
	var s string
	for _, p := range Pairs {
		s = s + fmt.Sprintf("%s (%s) \t\t %s (%s) \n", p.seat1.name, p.seat1.industry, p.seat2.name, p.seat2.industry)

	}
	s = s + "\n"

	return s
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

func arrayContains(needle int, haystack []int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}

	return false
}

func (a *AppData) peek(i int) Attendee {
	return a.Attendees[i]
}

func (a *AppData) selectPartner(seat1 Attendee, c []Attendee) Attendee {
	var seat2 Attendee
	//i := randomInt(0, len(d.Attendees))
	//seat2 = d.peek(i)
	for {
		i := randomInt(0, len(a.Attendees))
		seat2 = a.peek(i)

		if seat2.industry != seat1.industry {
			if seat1.pairedWith == nil {
				// remove from slice - swap to end and reslice
				a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
				a.Attendees = a.Attendees[:len(a.Attendees)-1]

				return seat2
			}
			for _, v := range c {
				if v.id == seat2.id {
					if !arrayContains(seat1.id, seat2.pairedWith) {
						// these two have not been paired before
						// remove from slice - swap to end and reslice
						a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
						a.Attendees = a.Attendees[:len(a.Attendees)-1]

						return seat2
					}

				}
			}

		}
	}
}

func (a *AppData) shiftArray() Attendee {
	t := a.Attendees[0]
	a.Attendees = a.Attendees[1:]

	return t
}

func SetIndustries() (industries []string) {
	industries = append(industries, "Accountants & Tax Preparation",
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

	return
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}
