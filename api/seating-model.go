/**
filename:	seating-model.go
purpose:	provides models for the application
update:		1/28/2022
comments:	3/3/2020 all of the models as well as the function to load the industries data have been moved to this file.
			1/28/2022 adding controller with db connection
*/

package api

import "go.mongodb.org/mongo-driver/mongo"

// TODO: I THINK THIS STRUCT AND ANYTHING THAT USES IT CAN GO AWAY
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
	Industries []string      `json:"industries"`
	Attendees  []Attendee    `json:"attendees"`
	Pairs      []Pair        `json:"pairs"`
	ListCount  int           `json:"listCount"`
	Conn       *mongo.Client `json:"-"`
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

// SetIndustries is used to populate the data for the industry option select on the form
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
