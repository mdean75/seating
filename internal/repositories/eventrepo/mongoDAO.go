package eventrepo

import (
	"context"
	"fmt"
	"seating/internal/app/domain"
	"seating/internal/app/ports"
	"seating/internal/db"
	"seating/internal/repositories/attendeerepo"
	"seating/internal/repositories/grouprepo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDataStore struct {
	*db.MongoConn
	db  *mongo.Database
	col *mongo.Collection
}

func NewDAO(dbconn *db.MongoConn, db, col string) ports.EventRepository {
	dbx := dbconn.Client.Database(db)
	conx := dbx.Collection(col)

	return &MongoDataStore{dbconn, dbx, conx}
}

type Event struct {
	ID            string                  `bson:"_id,omitempty"`
	Date          time.Time               `bson:"date"`
	GroupID       string                  `bson:"groupId"`
	Group         grouprepo.Group         `bson:"group"`
	Attendees     []attendeerepo.Attendee `bson:"attendees"`
	PairingRounds [][]Pair                `bson:"pairingRounds"`
}

//type PairingRound struct {
//	Pairs []Pair
//}

type Pair struct {
	Seat1 attendeerepo.Attendee `bson:"seat1"`
	Seat2 attendeerepo.Attendee `bson:"seat2"`
}

func NewMongoEventFromDomain(domainEvent domain.Event) Event {
	return Event{
		ID:      domainEvent.ID,
		Date:    domainEvent.Date,
		GroupID: domainEvent.GroupID,
		Group:   grouprepo.NewMongoGroupFromDomain(domainEvent.Group),
		//Attendees: domainEvent.Attendees,
	}
}

func convertMongoEventToDomain(event Event) domain.Event {
	var domainAttendees []domain.Attendee

	for _, attendee := range event.Attendees {
		var pairedWithAttendees []domain.Attendee
		for _, pairedWithAttendee := range attendee.PairedWith {
			//var pairedWithAttendees []domain.Attendee
			pairedWithAttendees = append(pairedWithAttendees, domain.Attendee{
				ID:          pairedWithAttendee.ID,
				Name:        pairedWithAttendee.Name,
				CompanyName: pairedWithAttendee.CompanyName,
				Industry:    pairedWithAttendee.Industry,
				//PairedWith:     pairedWithAttendee.,
				//PairedWithID:   nil,
				//PairedWithName: nil,
			})

			domainAttendees = append(domainAttendees, domain.Attendee{
				ID:          attendee.ID,
				Name:        attendee.Name,
				CompanyName: attendee.CompanyName,
				Industry:    attendee.Industry,
				PairedWith:  pairedWithAttendees,
				//PairedWithID:   nil,
				//PairedWithName: nil,
			})
		}
		//domainAttendees = append(domainAttendees, domain.)
	}

	var domainPairingRounds []domain.PairingRound

	for _, round := range event.PairingRounds {
		var p domain.PairingRound
		for _, pair := range round {
			seat1 := domain.Attendee{
				ID:          pair.Seat1.ID,
				Name:        pair.Seat1.Name,
				CompanyName: pair.Seat1.CompanyName,
				Industry:    pair.Seat1.Industry,
			}

			seat2 := domain.Attendee{

				ID:          pair.Seat2.ID,
				Name:        pair.Seat2.Name,
				CompanyName: pair.Seat2.CompanyName,
				Industry:    pair.Seat2.Industry,
			}

			p.Pairs = append(p.Pairs, domain.NewPair(seat1, seat2))

			//domaa := domain.PairingRound{
			//	Pairs: domain.Pair{
			//		Seat1: seat1,
			//		Seat2: seat2,
			//	},
			//	Attendees: nil,
			//}
		}
		domainPairingRounds = append(domainPairingRounds, p)
	}

	//var domAttendees []domain.Attendee
	//for _, attendee := range event.Attendees {
	//	domAttendees = append(domAttendees, domain.NewAttendee(attendee.ID, attendee.Name, attendee.CompanyName, attendee.Industry))
	//}
	//return domain.NewEvent(event.ID, event.GroupID, event.Date, grouprepo.ConvertMongoGroupToDomain(event.Group), domainAttendees...)

	//return domain.NewEvent(event.ID, event.GroupID, event.Date, grouprepo.ConvertMongoGroupToDomain(event.Group))
	return domain.Event{
		ID:           event.ID,
		Date:         event.Date,
		GroupID:      event.GroupID,
		Group:        grouprepo.ConvertMongoGroupToDomain(event.Group),
		Attendees:    domainAttendees,
		PairingRound: domainPairingRounds,
	}

}

func (m *MongoDataStore) Save(event domain.Event) (ports.ID, error) {

	e := NewMongoEventFromDomain(event)
	if e.Attendees == nil {
		e.Attendees = make([]attendeerepo.Attendee, 0)
	}

	res, err := m.col.InsertOne(context.TODO(), e)
	if err != nil {
		return "", err
	}

	objID := res.InsertedID.(primitive.ObjectID)

	return ports.ID(objID.Hex()), nil
}

func (m *MongoDataStore) Get(eventID string) (domain.Event, error) {
	var event Event
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return domain.Event{}, err
	}

	err = m.col.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&event)
	if err != nil {
		return domain.Event{}, err
	}

	return convertMongoEventToDomain(event), nil

}

var ErrUnableToDeleteResource error = fmt.Errorf("unable to delete the specified resource")

func (m *MongoDataStore) Delete(eventID string) error {
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	result, err := m.col.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return ErrUnableToDeleteResource
	}

	return nil
}

func (m *MongoDataStore) SaveRound(eventID string, pairs []domain.Pair) error {
	id, err := primitive.ObjectIDFromHex(eventID)
	if err != nil {
		return err
	}

	// convert domain pairs to mongo pairs
	var upsertPairs []Pair
	for _, p := range pairs {
		upsertPairs = append(upsertPairs,
			Pair{
				Seat1: attendeerepo.NewAttendee(p.Seat1.ID, p.Seat1.Name, p.Seat1.CompanyName, p.Seat1.Industry),
				Seat2: attendeerepo.NewAttendee(p.Seat2.ID, p.Seat2.Name, p.Seat2.CompanyName, p.Seat2.Industry)})
	}

	updateQuery := bson.M{"$addToSet": bson.M{"pairingRounds": upsertPairs}}

	result, err := m.col.UpdateByID(context.TODO(), id, updateQuery)
	if err != nil {
		fmt.Println("pairing rounds", err)
		return err
	}

	if result.ModifiedCount == 0 {
		fmt.Println("pairing list not upserted", err)
		return fmt.Errorf("record %s was not updated", id)
	}

	// now we need to update the attendees 'paired with'
	for _, p := range upsertPairs {

		//matchQuery := bson.M{"_id": id, "attendees.attendeeId": p.Seat1.ID}
		updatePairedWithQuery := bson.M{"$push": bson.M{"attendees.$.pairedWith": p.Seat2}}

		if p.Seat1.Name != "Placeholder" {
			result, err := m.col.UpdateOne(context.TODO(), bson.M{"_id": id, "attendees.attendeeId": p.Seat1.ID}, updatePairedWithQuery)
			if err != nil {
				fmt.Println("update seat1 error", err)
				return err
			}

			if result.ModifiedCount == 0 {
				fmt.Println("update seat1 no match", result)
				//fmt.Println("match query: ", matchQuery)
				//fmt.Println("update query", updateQuery)
				return fmt.Errorf("record %s was not updated", id)
			}
		}

		if p.Seat2.Name != "Placeholder" {
			matchQuery2 := bson.M{"_id": id, "attendees.attendeeId": p.Seat2.ID}
			updatePairedWithQuery2 := bson.M{"$push": bson.M{"attendees.$.pairedWith": p.Seat1}}

			result, err = m.col.UpdateOne(context.TODO(), matchQuery2, updatePairedWithQuery2)
			if err != nil {
				fmt.Println("update seat2 error", err)
				return err
			}

			if result.ModifiedCount == 0 {
				fmt.Println("update seat2 no match", result)
				return fmt.Errorf("record %s was not updated", id)
			}
		}

	}

	return nil
}
