package api

import (
	"context"
	"fmt"
	"net/http"
)

// func (a *AppData) Demo(w http.ResponseWriter, r *http.Request) {
// 	a.LoadDemoData()

// 	http.Redirect(w, r, "/", http.StatusFound)
// }

func (a *AppData) DemoAPI(w http.ResponseWriter, r *http.Request) {
	a.LoadDemoData()

	res, err := a.Conn.Database("testdb").Collection("testcol").InsertOne(context.TODO(), map[string]interface{}{"attendees": a.Attendees, "pairs": a.Pairs})
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(res.InsertedID)
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
			ID:       1,
			Industry: "Office Equipment & Copiers",
			Business: "Cartridge World",
		},
		{
			Name:     "Elisa Zieg",
			ID:       2,
			Industry: "Non-Profit Organization",
			Business: "The Sparrow's Nest Maternity Home",
		},
		{
			Name:     "Ben Motil",
			ID:       3,
			Industry: "Community",
			Business: "City of O'Fallon",
		},
		{
			Name:     "Kris Kinsinger",
			ID:       4,
			Industry: "Window Treatments",
			Business: "Two Blind Guys",
		},
		{
			Name:     "Anna Alt",
			ID:       5,
			Industry: "Home Improvement",
			Business: "LSL Finishes",
		},
		{
			Name:     "Danielle Fristoe",
			ID:       6,
			Industry: "Health & Wellness",
			Business: "Melaleuka Wellness Company",
		},
		{
			Name:     "Betty Bauer",
			ID:       7,
			Industry: "Insurance Services",
			Business: "Compass Health Consultants - Healthcare Solutions Team",
		},
		{
			Name:     "Nicolas Ippolito",
			ID:       8,
			Industry: "Accountants & Tax Preparation",
			Business: "Managerial Accounting Service",
		},
		{
			Name:     "Jim Mason",
			ID:       9,
			Industry: "Travel Services",
			Business: "The Cruise & Travel Planner",
		},
		{
			Name:     "Kathy Fleming",
			ID:       10,
			Industry: "Human Resource Services",
			Business: "People Solutions Center",
		},
		{
			Name:     "Jodie Uhlemeyer",
			ID:       11,
			Industry: "Pain Management",
			Business: "Arch Advanced Pain Management - Find A Better You",
		},
		{
			Name:     "Brad Henningfeld",
			ID:       12,
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Matt Buetow",
			ID:       13,
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Ann Lubieswki",
			ID:       14,
			Industry: "Travel Services",
			Business: "Travel by Ann",
		},
		{
			Name:     "Rodney Schrum",
			ID:       15,
			Industry: "Text Message Marketing",
			Business: "Ameritext",
		},
		{
			Name:     "Angie Harness",
			ID:       16,
			Industry: "Real Estate: Residential",
			Business: "Keller Williams - Angie Harness",
		},
		{
			Name:     "Heidi Martin",
			ID:       17,
			Industry: "Financial Services",
			Business: "Principal Financial Group",
		},
		{
			Name:     "Linda Otto",
			ID:       18,
			Industry: "Retail Shopping",
			Business: "Patty-Cakes of St. Louis",
		},
		{
			Name:     "Mary Agan",
			ID:       19,
			Industry: "Insurance Services",
			Business: "Kathy Kilo Peterson State Farm",
		},
		{
			Name:     "Caitlyn Baratti",
			ID:       20,
			Industry: "Insurance Services",
			Business: "Haight Insurance Agency, LLC",
		},
		{
			Name:     "Lindsey Helland",
			ID:       21,
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Stephanie DiCiro",
			ID:       22,
			Industry: "Winery",
			Business: "Cedar Lake Cellars",
		},
		{
			Name:     "Katie Worzel",
			ID:       23,
			Industry: "Health & Wellness",
			Business: "Cornerstone Care",
		},
		{
			Name:     "Rick Nies",
			ID:       24,
			Industry: "Home Improvement",
			Business: "MidWest SoftWash",
		},
		{
			Name:     "Skip Stephens",
			ID:       25,
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Wil Skaggs",
			ID:       26,
			Industry: "Fire Protection",
			Business: "Cottleville Fire Protection",
		},
		{
			Name:     "Deanna Hoffman",
			ID:       27,
			Industry: "Retail Shopping",
			Business: "Touchstone Crystal Jewelry by Swarovski",
		},
		{
			Name:     "Dan Tripp",
			ID:       28,
			Industry: "Coffee House",
			Business: "Alpha & Omega Roasting Company",
		},
		{
			Name:     "Greg Kinder",
			ID:       29,
			Industry: "Civic Organizations",
			Business: "O'Fallon Lions Club",
		},
		{
			Name:     "Kim Henson",
			ID:       30,
			Industry: "Marketing: Sales Promotions",
			Business: "Circle of Marketing",
		},
		{
			Name:     "Brian Richardson",
			ID:       31,
			Industry: "Radio Station",
			Business: "99.9 FM KFAV & 730 AM KWRE Kaspar Broadcasting",
		},
		{
			Name:     "Shelley Barr",
			ID:       32,
			Industry: "Radio Station",
			Business: "104.5 FM KSLQ",
		},
		{
			Name:     "Rich Johns",
			ID:       33,
			Industry: "Restoration: Fire & Flood",
			Business: "CATCO Catastrophe Cleaning & Restoration Company",
		},
		{
			Name:     "Jennifer Begley",
			ID:       34,
			Industry: "Banks",
			Business: "Reliance Bank",
		},
		{
			Name:     "Katy Kruze",
			ID:       35,
			Industry: "Radio Station",
			Business: "K-Wulf 101.7FM",
		},
	}
}
