package api

import "net/http"

func (a *AppData) Demo(w http.ResponseWriter, r *http.Request) {
	a.LoadDemoData()

	http.Redirect(w, r, "/", http.StatusFound)
}

func (a *AppData) LoadDemoData() {

	a.Attendees = []Attendee{
		{
			name:     "Bob Davis",
			id:       randomInt(1, 1000),
			industry: "Office Equipment & Copiers",
			business: "Cartridge World",
		},
		{
			name:     "Elisa Zieg",
			id:       randomInt(1, 1000),
			industry: "Non-Profit Organization",
			business: "The Sparrow's Nest Maternity Home",
		},
		{
			name:     "Ben Motil",
			id:       randomInt(1, 1000),
			industry: "Community",
			business: "City of O'Fallon",
		},
		{
			name:     "Kris Kinsinger",
			id:       randomInt(1, 1000),
			industry: "Window Treatments",
			business: "Two Blind Guys",
		},
		{
			name:     "Anna Alt",
			id:       randomInt(1, 1000),
			industry: "Home Improvement",
			business: "LSL Finishes",
		},
		{
			name:     "Danielle Fristoe",
			id:       randomInt(1, 1000),
			industry: "Health & Wellness",
			business: "Melaleuka Wellness Company",
		},
		{
			name:     "Betty Bauer",
			id:       randomInt(1, 1000),
			industry: "Insurance Services",
			business: "Compass Health Consultants - Healthcare Solutions Team",
		},
		{
			name:     "Nicolas Ippolito",
			id:       randomInt(1, 1000),
			industry: "Accountants & Tax Preparation",
			business: "Managerial Accounting Service",
		},
		{
			name:     "Jim Mason",
			id:       randomInt(1, 1000),
			industry: "Travel Services",
			business: "The Cruise & Travel Planner",
		},
		{
			name:     "Kathy Fleming",
			id:       randomInt(1, 1000),
			industry: "Human Resource Services",
			business: "People Solutions Center",
		},
		{
			name:     "Jodie Uhlemeyer",
			id:       randomInt(1, 1000),
			industry: "Pain Management",
			business: "Arch Advanced Pain Management - Find A Better You",
		},
		{
			name:     "Brad Henningfeld",
			id:       randomInt(1, 1000),
			industry: "Financial Services",
			business: "Principal Financial Group",
		},
		{
			name:     "Matt Buetow",
			id:       randomInt(1, 1000),
			industry: "Financial Services",
			business: "Principal Financial Group",
		},
		{
			name:     "Ann Lubieswki",
			id:       randomInt(1, 1000),
			industry: "Travel Services",
			business: "Travel by Ann",
		},
		{
			name:     "Rodney Schrum",
			id:       randomInt(1, 1000),
			industry: "Text Message Marketing",
			business: "Ameritext",
		},
		{
			name:     "Angie Harness",
			id:       randomInt(1, 1000),
			industry: "Real Estate: Residential",
			business: "Keller Williams - Angie Harness",
		},
		{
			name:     "Heidi Martin",
			id:       randomInt(1, 1000),
			industry: "Financial Services",
			business: "Principal Financial Group",
		},
		{
			name:     "Linda Otto",
			id:       randomInt(1, 1000),
			industry: "Retail Shopping",
			business: "Patty-Cakes of St. Louis",
		},
		{
			name:     "Mary Agan",
			id:       randomInt(1, 1000),
			industry: "Insurance Services",
			business: "Kathy Kilo Peterson State Farm",
		},
		{
			name:     "Caitlyn Baratti",
			id:       randomInt(1, 1000),
			industry: "Insurance Services",
			business: "Haight Insurance Agency, LLC",
		},
		{
			name:     "Lindsey Helland",
			id:       randomInt(1, 1000),
			industry: "Winery",
			business: "Cedar Lake Cellars",
		},
		{
			name:     "Stephanie DiCiro",
			id:       randomInt(1, 1000),
			industry: "Winery",
			business: "Cedar Lake Cellars",
		},
		{
			name:     "Katie Worzel",
			id:       randomInt(1, 1000),
			industry: "Health & Wellness",
			business: "Cornerstone Care",
		},
		{
			name:     "Rick Nies",
			id:       randomInt(1, 1000),
			industry: "Home Improvement",
			business: "MidWest SoftWash",
		},
		{
			name:     "Skip Stephens",
			id:       randomInt(1, 1000),
			industry: "Fire Protection",
			business: "Cottleville Fire Protection",
		},
		{
			name:     "Wil Skaggs",
			id:       randomInt(1, 1000),
			industry: "Fire Protection",
			business: "Cottleville Fire Protection",
		},
		{
			name:     "Deanna Hoffman",
			id:       randomInt(1, 1000),
			industry: "Retail Shopping",
			business: "Touchstone Crystal Jewelry by Swarovski",
		},
		{
			name:     "Dan Tripp",
			id:       randomInt(1, 1000),
			industry: "Coffee House",
			business: "Alpha & Omega Roasting Company",
		},
		{
			name:     "Greg Kinder",
			id:       randomInt(1, 1000),
			industry: "Civic Organizations",
			business: "O'Fallon Lions Club",
		},
		{
			name:     "Kim Henson",
			id:       randomInt(1, 1000),
			industry: "Marketing: Sales Promotions",
			business: "Circle of Marketing",
		},
		{
			name:     "Brian Richardson",
			id:       randomInt(1, 1000),
			industry: "Radio Station",
			business: "99.9 FM KFAV & 730 AM KWRE Kaspar Broadcasting",
		},
		{
			name:     "Shelley Barr",
			id:       randomInt(1, 1000),
			industry: "Radio Station",
			business: "104.5 FM KSLQ",
		},
		{
			name:     "Rich Johns",
			id:       randomInt(1, 1000),
			industry: "Restoration: Fire & Flood",
			business: "CATCO Catastrophe Cleaning & Restoration Company",
		},
		{
			name:     "Jennifer Begley",
			id:       randomInt(1, 1000),
			industry: "Banks",
			business: "Reliance Bank",
		},
		{
			name:     "Katy Kruze",
			id:       randomInt(1, 1000),
			industry: "Radio Station",
			business: "K-Wulf 101.7FM",
		},
	}
}
