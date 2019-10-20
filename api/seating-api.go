package api

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"seating/app"
)

// data struct for the add attendee template
type inputData struct {
	Industries   []string
	KeyErr       string
	Name         string
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

// SecretsForm is the handler to display the user input form.
func (a *AppData) SecretsForm(w http.ResponseWriter, r *http.Request) {

	var tData inputData
	tData.Industries = a.Industries

	err := loadForm(app.InputForm, w, tData)
	if err != nil {
		//log.Debug("error in SecretsForm() received from LoadForm()", err.Error())
	}

}

func (a *AppData) ProcessSecretsForm(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
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

	http.Redirect(w, r, "/", http.StatusSeeOther)

	//fmt.Printf("name: %s, business: %s, industry: %s", name, business, industry)
	//var tData inputData
	//tData.Industries = SetIndustries()
	//
	//err = loadForm(app.InputForm, w, tData)
	//if err != nil {
	//	//log.Debug("error in SecretsForm() received from LoadForm()", err.Error())
	//}

}

func (a *AppData) DisplayAttendees(w http.ResponseWriter, r *http.Request) {

	b, err := json.Marshal(a.Attendees)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(b)
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
