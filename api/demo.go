package api

import "net/http"

func (a *AppData) Demo(w http.ResponseWriter, r *http.Request) {
	a.LoadDemoData()

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *AppData) DemoAPI(w http.ResponseWriter, r *http.Request) {
	a.LoadDemoData()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func (a *AppData) LoadDemoData() {

	a.Attendees = []Attendee{
		{
			Name:     "Bob Davis",
			ID:       a.generateID(),
			Industry: "Office Equipment & Copiers",
			Business: "Cartridge World",
		},
		{
			Name:     "Elisa Zieg",
			ID:       a.generateID(),
			Industry: "Non-Profit Organization",
			Business: "The Sparrow's Nest Maternity Home",
		},
		{
			Name:     "Ben Motil",
			ID:       a.generateID(),
			Industry: "Community",
			Business: "City of O'Fallon",
		},
		{
			Name:     "Kris Kinsinger",
			ID:       a.generateID(),
			Industry: "Window Treatments",
			Business: "Two Blind Guys",
		},
		{
			Name:     "Anna Alt",
			ID:       a.generateID(),
			Industry: "Home Improvement",
			Business: "LSL Finishes",
		},
		{
			Name:     "Danielle Fristoe",
			ID:       a.generateID(),
			Industry: "Health & Wellness",
			Business: "Melaleuka Wellness Company",
		},
		{
			Name:     "Betty Bauer",
			ID:       a.generateID(),
			Industry: "Insurance Services",
			Business: "Compass Health Consultants - Healthcare Solutions Team",
		},
		{
			Name:     "Nicolas Ippolito",
			ID:       a.generateID(),
			Industry: "Accountants & Tax Preparation",
			Business: "Managerial Accounting Service",
		},
		{
			Name:     "Jim Mason",
			ID:       a.generateID(),
			Industry: "Travel Services",
			Business: "The Cruise & Travel Planner",
		},
		{
			Name:     "Kathy Fleming",
			ID:       a.generateID(),
			Industry: "Human Resource Services",
			Business: "People Solutions Center",
		},
		{
			Name:     "Jodie Uhlemeyer",
			ID:       a.generateID(),
			Industry: "Pain Management",
			Business: "Arch Advanced Pain Management - Find A Better You",
		},
		{
			Name:     "Brad Henningfeld",
			ID:       a.generateID(),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Matt Buetow",
			ID:       a.generateID(),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Ann Lubieswki",
			ID:       a.generateID(),
			Industry: "Travel Services",
			Business: "Travel by Ann",
		},
		{
			Name:     "Rodney Schrum",
			ID:       a.generateID(),
			Industry: "Text Message Marketing",
			Business: "Ameritext",
		},
		{
			Name:     "Angie Harness",
			ID:       a.generateID(),
			Industry: "Real Estate: Residential",
			Business: "Keller Williams - Angie Harness",
		},
		{
			Name:     "Heidi Martin",
			ID:       a.generateID(),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Linda Otto",
			ID:       a.generateID(),
			Industry: "Retail Shopping",
			Business: "Patty-Cakes of St. Louis",
		},
		{
			Name:     "Mary Agan",
			ID:       a.generateID(),
			Industry: "Insurance Services",
			Business: "Kathy Kilo Peterson State Farm",
		},
		{
			Name:     "Caitlyn Baratti",
			ID:       a.generateID(),
			Industry: "Insurance Services",
			Business: "Haight Insurance Agency, LLC",
		},
		{
			Name:     "Lindsey Helland",
			ID:       a.generateID(),
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Stephanie DiCiro",
			ID:       a.generateID(),
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Katie Worzel",
			ID:       a.generateID(),
			Industry: "Health & Wellness",
			Business: "Cornerstone Care",
		},
		{
			Name:     "Rick Nies",
			ID:       a.generateID(),
			Industry: "Home Improvement",
			Business: "MidWest SoftWash",
		},
		{
			Name:     "Skip Stephens",
			ID:       a.generateID(),
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Wil Skaggs",
			ID:       a.generateID(),
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Deanna Hoffman",
			ID:       a.generateID(),
			Industry: "Retail Shopping",
			Business: "Touchstone Crystal Jewelry by Swarovski",
		},
		{
			Name:     "Dan Tripp",
			ID:       a.generateID(),
			Industry: "Coffee House",
			Business: "Alpha & Omega Roasting Company",
		},
		{
			Name:     "Greg Kinder",
			ID:       a.generateID(),
			Industry: "Civic Organizations",
			Business: "O'Fallon Lions Club",
		},
		{
			Name:     "Kim Henson",
			ID:       a.generateID(),
			Industry: "Marketing: Sales Promotions",
			Business: "Circle of Marketing",
		},
		{
			Name:     "Brian Richardson",
			ID:       a.generateID(),
			Industry: "Radio Station",
			Business: "99.9 FM KFAV & 730 AM KWRE Kaspar Broadcasting",
		},
		{
			Name:     "Shelley Barr",
			ID:       a.generateID(),
			Industry: "Radio Station",
			Business: "104.5 FM KSLQ",
		},
		{
			Name:     "Rich Johns",
			ID:       a.generateID(),
			Industry: "Restoration: Fire & Flood",
			Business: "CATCO Catastrophe Cleaning & Restoration Company",
		},
		{
			Name:     "Jennifer Begley",
			ID:       a.generateID(),
			Industry: "Banks",
			Business: "Reliance Bank",
		},
		{
			Name:     "Katy Kruze",
			ID:       a.generateID(),
			Industry: "Radio Station",
			Business: "K-Wulf 101.7FM",
		},
	}
}
