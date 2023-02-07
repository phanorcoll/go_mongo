// Package domain provides domain  
// Will contain structs reflecting the database tables.
package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct    Represets the fields stored in the DB
type User struct {
	ID       primitive.ObjectID `bson:_id`
	Username string             `json:"username"`
	Password string             `json:"password"`
}
