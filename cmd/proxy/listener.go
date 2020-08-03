package main

import (
	"fmt"
	"log"

	bolt "github.com/boltdb/bolt"
	nats "github.com/nats-io/nats.go"
)

// listens for scaling events sent from the controller
func listener(nc *nats.Conn, scaledUp, scaledDown *chanMgr, db *bolt.DB) {
	scaledUpCh := scaledUp.newReader()
	scaledDownCh := scaledDown.newReader()
	for {
		select {
		case msg := <-scaledUpCh:
			log.Printf("Listener got scaled up event")
			// if it's a scale-up, then add to the DB, otherwise, delete it
			err := db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("containers"))
				if err != nil {
					return fmt.Errorf("Couldn't create 'containers' bucket (%s)", err)
				}
				containerURL := msg.Data
				bucket.Put(containerURL, nil)
				return nil
			})
			if err != nil {
				log.Printf("Could not record scale up event")
			}
		case msg := <-scaledDownCh:
			log.Printf("Listener got scaled down event")
			err := db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("containers"))
				if err != nil {
					return fmt.Errorf("Couldn't create 'containers' bucket (%s)", err)
				}
				containerURL := msg.Data
				bucket.Delete(containerURL)
				return nil
			})
			if err != nil {
				log.Printf("Could not record scale down event")
			}
		}
	}
}