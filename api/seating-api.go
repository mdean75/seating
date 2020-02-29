package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rivo/sessions"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"seating/app"
	"time"
)

// data struct for the add attendee template
type inputData struct {
	Industries   []string `json:"industries"`
	SuccessMsg   string   `json:"successMsg"`
	Name         string   `json:"name"`
	KeyErr       string   `json:"keyErr"`
	BusinessName string   `json:"businessName"`
	ListCount    int      `json:"listCount"`
}

// overall data for the application
type AppData struct {
	Industries []string   `json:"industries"`
	Attendees  []Attendee `json:"attendees"`
	Pairs      []Pair     `json:"pairs"`
	ListCount  int        `json:"listCount"`
}

type Attendee struct {
	Name           string   `json:"name" bson:"name"`
	ID             int      `json:"id" bson:"id"`
	Industry       string   `json:"industry" bson:"industry"`
	Business       string   `json:"business" bson:"business"`
	PairedWith     []int    `json:"pairedWith" bson:"pairedWith"`
	PairedWithName []string `json:"pairedWithName" bson:"pairedWithName"`
}

type Pair struct {
	Seat1 Attendee `json:"seat1"`
	Seat2 Attendee `json:"seat2"`
}

func (a *AppData) ResetData(w http.ResponseWriter, r *http.Request) {

	// start new session and create cookie
	session, _ := sessions.Start(w, r, true)
	session.Set("successMsg", "Meeting attendees have been reset")

	a.Attendees = []Attendee{}
	a.ListCount = 0

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
	tData.ListCount = a.ListCount

	err = loadForm(app.InputForm, w, tData)
	if err != nil {
		//log.Debug("error in AttendeeEntry() received from LoadForm()", err.Error())
	}

}

func (a *AppData) ProcessAttendeeEntry(w http.ResponseWriter, r *http.Request) {

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
		Name:     name,
		Business: business,
		ID:       a.generateID(),
		Industry: industry,
	}

	if len(a.Attendees) > 0 {
		if a.Attendees[len(a.Attendees)-1].Name == "Placeholder" {
			a.Attendees = a.Attendees[:len(a.Attendees)-1]
		}
	}

	a.Attendees = append(a.Attendees, attendee)

	err = session.Set("successMsg", fmt.Sprintf("Added %s to meeting", attendee.Name))

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func (a *AppData) AddAttendeeAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var att Attendee
		err := json.NewDecoder(r.Body).Decode(&att)
		if err != nil {
			resp := map[string]string{"error": err.Error()}

			//resp := map[string]string{"response": "successfully added attendee"}
			b, _ := json.Marshal(resp)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(b)

			return
		}

		att.ID = a.generateID() //randomInt(1, 1000)
		a.Attendees = append(a.Attendees, att)

		resp := map[string]string{"response": "successfully added attendee"}
		b, err := json.Marshal(resp)

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}

}

func (a *AppData) DisplayAttendeesAPI(w http.ResponseWriter, r *http.Request) {

	b, _ := json.Marshal(a.Attendees)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (a *AppData) ResetAttendeesAPI(w http.ResponseWriter, r *http.Request) {

	a.Attendees = []Attendee{}
	a.ListCount = 0

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func (a *AppData) DisplayPairsAPI(w http.ResponseWriter, r *http.Request) {

	m := struct {
		Pairs []Pair
	}{Pairs: a.Pairs}

	b, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (a *AppData) GetAppData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a)
	}
}

func (a *AppData) GetListCount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(a.ListCount)
	}
}

func (a *AppData) GetIndustries() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := struct {
			Industries []string
		}{Industries: a.Industries}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(m)
	}
}

func (a *AppData) DisplayAttendees(w http.ResponseWriter, r *http.Request) {

	var s string
	for _, att := range a.Attendees {
		//s = s + att.Name + att.Business + "\n"
		name := att.Name
		if len(name) < 20 {
			numSpaces := 20 - len(name)
			i := 0
			for i < numSpaces {
				name = name + " "
				i++
			}
		}
		s = s + name + att.Business + "\n"
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

func setSessionError(msg, url string, w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Start(w, r, true)
	if err != nil {

	}
	err = session.Set("successMsg", msg)
	http.Redirect(w, r, url, http.StatusSeeOther)

}
func (a *AppData) BuildChart(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func(ctx context.Context) {
		defer cancel()
		// ensure enough attendees have been entered
		if len(a.Attendees) < 5 {
			// start new session and create cookie
			msg := "Unable to build seating charts, not enough attendees!"
			setSessionError(msg, "/", w, r)

			return
		}

		// add a placeholder member if odd number registered
		if len(a.Attendees)%2 != 0 {
			a.Attendees = append(a.Attendees, Attendee{Name: "Placeholder"})
		}

		c := make([]Attendee, len(a.Attendees))
		num := copy(c, a.Attendees)
		fmt.Printf("number copied: %d \n", num)

		var s string

		for ok := true; ok; ok = len(a.Attendees) > 2 {
			var p Pair

			p.Seat1 = a.shiftArray()
			p.Seat2 = a.selectPartner(p.Seat1, c)

			a.Pairs = append(a.Pairs, p)
			addAttendeePairing(p.Seat1, p.Seat2, c)
		}

		// select last two no matter the match
		lastPair1 := a.shiftArray()
		lastPair2 := a.shiftArray()
		a.Pairs = append(a.Pairs, Pair{
			Seat1: lastPair1,
			Seat2: lastPair2,
		})

		addAttendeePairing(lastPair1, lastPair2, c)

		s += printPairs(a.Pairs)

		a.Pairs = []Pair{}
		// the slice should be nil at this point, reload original data
		a.Attendees = make([]Attendee, len(c))
		num = copy(a.Attendees, c)
		fmt.Printf("number copied: %d \n", num)

		a.ListCount++
		a.Pairs = []Pair{}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-type", "application/json")
		w.Write([]byte(s))

	}(ctx)

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println(ctx.Err())
			setSessionError("unable to generate report, cannot find unique pairing", "/", w, r)
		}
	}

}

func (a *AppData) BuildChartAPI(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	go func(ctx context.Context) {
		defer cancel()
		// ensure enough attendees have been entered
		if len(a.Attendees) < 5 {
			m := map[string]string{"error": "Unable to build seating charts, not enough attendees!"}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(m)

			return
		}

		// add a placeholder member if odd number registered
		if len(a.Attendees)%2 != 0 {
			a.Attendees = append(a.Attendees, Attendee{Name: "Placeholder"})
		}

		c := make([]Attendee, len(a.Attendees))
		num := copy(c, a.Attendees)
		fmt.Printf("number copied: %d \n", num)

		//var s string

		for ok := true; ok; ok = len(a.Attendees) > 2 {
			var p Pair

			p.Seat1 = a.shiftArray()
			p.Seat2 = a.selectPartner(p.Seat1, c)

			a.Pairs = append(a.Pairs, p)
			addAttendeePairing(p.Seat1, p.Seat2, c)
		}

		// select last two no matter the match
		lastPair1 := a.shiftArray()
		lastPair2 := a.shiftArray()
		a.Pairs = append(a.Pairs, Pair{
			Seat1: lastPair1,
			Seat2: lastPair2,
		})

		addAttendeePairing(lastPair1, lastPair2, c)

		//m := struct {
		//	Pairs []Pair
		//}{Pairs: a.Pairs}

		b, err := json.Marshal(a.Pairs)
		if err != nil {
			fmt.Println(err)
		}

		a.Pairs = []Pair{}
		// the slice should be nil at this point, reload original data
		a.Attendees = make([]Attendee, len(c))
		num = copy(a.Attendees, c)
		fmt.Printf("number copied: %d \n", num)

		a.ListCount++
		a.Pairs = []Pair{}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Cache-Control", "no-store")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	}(ctx)

	select {
	case <-ctx.Done():
		if ctx.Err() == context.DeadlineExceeded {
			// TODO: TAKE OUT THE SESSION MESSAGE IN FAVOR OF BETTER LOGGING
			fmt.Println(ctx.Err())
			setSessionError("unable to generate report, cannot find unique pairing", "/", w, r)
		}
	}

}

//func (a *AppData) clearSeating() {
//	for i := 0; i < len(a.Attendees); i++ {
//		a.Attendees[i].PairedWith = nil
//		fmt.Println(a.Attendees[i])
//	}
//}

func printPairs(Pairs []Pair) string {
	var s string
	for _, p := range Pairs {

		p1 := fmt.Sprintf("%s (%s)", p.Seat1.Name, p.Seat1.Industry)
		if len(p1) < 60 {
			numSpaces := 60 - len(p1)
			i := 0
			for i < numSpaces {
				p1 = p1 + " "
				i++
			}
		}
		p2 := fmt.Sprintf("%s (%s)", p.Seat2.Name, p.Seat2.Industry)
		s = s + p1 + p2 + "\n"

	}
	s = s + "\n"

	return s
}

func addAttendeePairing(seat1 Attendee, seat2 Attendee, list []Attendee) {
	for i := range list {
		if list[i].ID == seat1.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat2.ID)
			list[i].PairedWithName = append(list[i].PairedWithName, seat2.Name)
		}
		if list[i].ID == seat2.ID {
			list[i].PairedWith = append(list[i].PairedWith, seat1.ID)
			list[i].PairedWithName = append(list[i].PairedWithName, seat1.Name)
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
	//Seat2 = d.peek(i)
	for {
		i := randomInt(0, len(a.Attendees))
		seat2 = a.peek(i)

		if seat2.Industry != seat1.Industry {
			if seat1.PairedWith == nil {
				// remove from slice - swap to end and reslice
				a.Attendees[i] = a.Attendees[len(a.Attendees)-1]
				a.Attendees = a.Attendees[:len(a.Attendees)-1]

				return seat2
			}
			for _, v := range c {
				if v.ID == seat2.ID {
					if !arrayContains(seat1.ID, seat2.PairedWith) {
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

func (a *AppData) generateID() int {
	var i int
	var b bool

	for !b {
		i = randomInt(1, 1000)
		b = a.isUniqueID(i)
		fmt.Println()
	}
	return i
}

func (a *AppData) isUniqueID(id int) bool {
	for _, attendee := range a.Attendees {
		if attendee.ID == id {
			return false // this id already exists, return false
		}
	}
	return true
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
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}
