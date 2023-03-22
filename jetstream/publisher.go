package main

import (
	"encoding/json"
	models "github.com/Shaad7/nats-examples/jetstream/model"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	SubjectNameReviewCreated = "REVIEWS.rateGiven"
)

func publishReviews(js nats.JetStreamContext) {
	reviews, err := getReviews()
	if err != nil {
		log.Println(err)
		return
	}

	for _, oneReview := range reviews {
		// create random message intervals to slow down
		r := rand.Intn(1500)
		time.Sleep(time.Duration(r) * time.Millisecond)

		reviewString, err := json.Marshal(oneReview)
		if err != nil {
			log.Println(err)
			continue
		}
		// publish to REVIEWS.rateGiven subject
		_, err = js.Publish(SubjectNameReviewCreated, reviewString)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Publisher  =>  Message:%s\n", oneReview.Text)
		}
	}
}

func getReviews() ([]models.Review, error) {
	rawReviews, _ := os.ReadFile("/home/shaad/go/src/github.com/Shaad7/nats-examples/jetstream/model/review.json")
	var reviewsObj []models.Review
	err := json.Unmarshal(rawReviews, &reviewsObj)

	return reviewsObj, err
}
