package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

type Participant struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	RSVP  string `json:"rvsp" bson:"rvsp"`
}

type Meeting struct {
	Id           string        `json:"id" bson:"id"`
	Title        string        `json:"title" bson:"title"`
	Participants []Participant `json:"participants" bson:"participants"`
	Starttime    string        `json:"starttime" bson:"starttime"`
	Endtime      string        `json:"endtime" bson:"endtime"`
	Timestamp    time.Time     `json:"timestamp" bson:"timestamp"`
}
type Allmeetings struct {
	Heading  string
	Meetings []Meeting
}

var allMeetings = Allmeetings{}

func allMettingHnadler(w http.ResponseWriter, r *http.Request) {
	// allMeetings.Heading = "Total number of meetings"
	// allMeetings.Meetings = append(allMeetings.Meetings, meeting)
	// allMeetingsJSON, err := json.Marshal(allMeetings)
	// w.Write(allMeetingsJSON)
}

func meetingHandler(response http.ResponseWriter, request *http.Request) {

	response.Header().Set("content-type", "application/json")
	var meeting Meeting
	_ = json.NewDecoder(request.Body).Decode(&meeting)
	collection := client.Database("harshith").Collection("meeting")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// result, _ := collection.InsertOne(ctx, meeting)
	collection.InsertOne(ctx, meeting)
	// json.NewEncoder(response).Encode(result)
	meetingJSON, err := json.Marshal(meeting)
	if err != nil {
		panic(err)
	}
	response.WriteHeader(http.StatusOK)
	response.Write(meetingJSON)

}

func getmeetingByIdHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println("GET params were:", request.URL)
}
func main() {
	http.HandleFunc("/meetings", meetingHandler)

	http.HandleFunc("/meeting/{id}", getmeetingByIdHandler)
	fmt.Println("Starting the application...")
	// var options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, clientOptions)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}

}
