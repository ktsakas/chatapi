package model

import "time"

type Message struct {
	ID	 				string `json:"id"`
	SenderID 			string `json:"sender"`
	RecipientID 		string `json:"receiver"`
	Content 			string `json:"content"`
	CreatedAt			time.Time
}