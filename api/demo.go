package api

import "net/http"

func (a *AppData) Demo(w http.ResponseWriter, r *http.Request) {
	a.LoadDemoData()

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *AppData) LoadDemoData() {

	a.Attendees = []Attendee{
		{
			Name:     "Bob Davis",
			ID:       randomInt(1, 1000),
			Industry: "Office Equipment & Copiers",
			Business: "Cartridge World",
		},
		{
			Name:     "Elisa Zieg",
			ID:       randomInt(1, 1000),
			Industry: "Non-Profit Organization",
			Business: "The Sparrow's Nest Maternity Home",
		},
		{
			Name:     "Ben Motil",
			ID:       randomInt(1, 1000),
			Industry: "Community",
			Business: "City of O'Fallon",
		},
		{
			Name:     "Kris Kinsinger",
			ID:       randomInt(1, 1000),
			Industry: "Window Treatments",
			Business: "Two Blind Guys",
		},
		{
			Name:     "Anna Alt",
			ID:       randomInt(1, 1000),
			Industry: "Home Improvement",
			Business: "LSL Finishes",
		},
		{
			Name:     "Danielle Fristoe",
			ID:       randomInt(1, 1000),
			Industry: "Health & Wellness",
			Business: "Melaleuka Wellness Company",
		},
		{
			Name:     "Betty Bauer",
			ID:       randomInt(1, 1000),
			Industry: "Insurance Services",
			Business: "Compass Health Consultants - Healthcare Solutions Team",
		},
		{
			Name:     "Nicolas Ippolito",
			ID:       randomInt(1, 1000),
			Industry: "Accountants & Tax Preparation",
			Business: "Managerial Accounting Service",
		},
		{
			Name:     "Jim Mason",
			ID:       randomInt(1, 1000),
			Industry: "Travel Services",
			Business: "The Cruise & Travel Planner",
		},
		{
			Name:     "Kathy Fleming",
			ID:       randomInt(1, 1000),
			Industry: "Human Resource Services",
			Business: "People Solutions Center",
		},
		{
			Name:     "Jodie Uhlemeyer",
			ID:       randomInt(1, 1000),
			Industry: "Pain Management",
			Business: "Arch Advanced Pain Management - Find A Better You",
		},
		{
			Name:     "Brad Henningfeld",
			ID:       randomInt(1, 1000),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Matt Buetow",
			ID:       randomInt(1, 1000),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Ann Lubieswki",
			ID:       randomInt(1, 1000),
			Industry: "Travel Services",
			Business: "Travel by Ann",
		},
		{
			Name:     "Rodney Schrum",
			ID:       randomInt(1, 1000),
			Industry: "Text Message Marketing",
			Business: "Ameritext",
		},
		{
			Name:     "Angie Harness",
			ID:       randomInt(1, 1000),
			Industry: "Real Estate: Residential",
			Business: "Keller Williams - Angie Harness",
		},
		{
			Name:     "Heidi Martin",
			ID:       randomInt(1, 1000),
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Linda Otto",
			ID:       randomInt(1, 1000),
			Industry: "Retail Shopping",
			Business: "Patty-Cakes of St. Louis",
		},
		{
			Name:     "Mary Agan",
			ID:       randomInt(1, 1000),
			Industry: "Insurance Services",
			Business: "Kathy Kilo Peterson State Farm",
		},
		{
			Name:     "Caitlyn Baratti",
			ID:       randomInt(1, 1000),
			Industry: "Insurance Services",
			Business: "Haight Insurance Agency, LLC",
		},
		{
			Name:     "Lindsey Helland",
			ID:       randomInt(1, 1000),
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Stephanie DiCiro",
			ID:       randomInt(1, 1000),
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Katie Worzel",
			ID:       randomInt(1, 1000),
			Industry: "Health & Wellness",
			Business: "Cornerstone Care",
		},
		{
			Name:     "Rick Nies",
			ID:       randomInt(1, 1000),
			Industry: "Home Improvement",
			Business: "MidWest SoftWash",
		},
		{
			Name:     "Skip Stephens",
			ID:       randomInt(1, 1000),
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Wil Skaggs",
			ID:       randomInt(1, 1000),
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Deanna Hoffman",
			ID:       randomInt(1, 1000),
			Industry: "Retail Shopping",
			Business: "Touchstone Crystal Jewelry by Swarovski",
		},
		{
			Name:     "Dan Tripp",
			ID:       randomInt(1, 1000),
			Industry: "Coffee House",
			Business: "Alpha & Omega Roasting Company",
		},
		{
			Name:     "Greg Kinder",
			ID:       randomInt(1, 1000),
			Industry: "Civic Organizations",
			Business: "O'Fallon Lions Club",
		},
		{
			Name:     "Kim Henson",
			ID:       randomInt(1, 1000),
			Industry: "Marketing: Sales Promotions",
			Business: "Circle of Marketing",
		},
		{
			Name:     "Brian Richardson",
			ID:       randomInt(1, 1000),
			Industry: "Radio Station",
			Business: "99.9 FM KFAV & 730 AM KWRE Kaspar Broadcasting",
		},
		{
			Name:     "Shelley Barr",
			ID:       randomInt(1, 1000),
			Industry: "Radio Station",
			Business: "104.5 FM KSLQ",
		},
		{
			Name:     "Rich Johns",
			ID:       randomInt(1, 1000),
			Industry: "Restoration: Fire & Flood",
			Business: "CATCO Catastrophe Cleaning & Restoration Company",
		},
		{
			Name:     "Jennifer Begley",
			ID:       randomInt(1, 1000),
			Industry: "Banks",
			Business: "Reliance Bank",
		},
		{
			Name:     "Katy Kruze",
			ID:       randomInt(1, 1000),
			Industry: "Radio Station",
			Business: "K-Wulf 101.7FM",
		},
	}
}
