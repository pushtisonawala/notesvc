package services

import (
	"context"
	"main/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notesservice struct{}

type Note struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

func NewNotesService() *Notesservice {
	return &Notesservice{}
}

func (n *Notesservice) GetNotes() []Note {
	var notes []Note
	cursor, err := database.NotesCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return []Note{}
	}
	defer cursor.Close(context.Background())
	if err := cursor.All(context.Background(), &notes); err != nil {
		return []Note{}
	}
	if notes == nil {
		return []Note{}
	}
	return notes
}

func (n *Notesservice) CreateNotes(note Note) Note {
	note.ID = primitive.NewObjectID()
	note.CreatedAt = time.Now()
	note.UpdatedAt = time.Now()
	database.NotesCollection.InsertOne(context.Background(), note)
	return note
}

func (n *Notesservice) DeleteNotes(id string) Note {
	objID, _ := primitive.ObjectIDFromHex(id)
	var deletedNote Note
	database.NotesCollection.FindOneAndDelete(context.Background(), bson.M{"_id": objID}).Decode(&deletedNote)
	return deletedNote
}
