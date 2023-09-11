package schemas

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quiz struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	LiveID   string             `json:"live_id" validate:"required"`
	Question string             `json:"question" validate:"required"`
	Answers  [4]Answer          `json:"answers" validate:"required"`
}

type Answer struct {
	Text    string `json:"text" validate:"required"`
	Correct bool   `json:"correct" validate:"required"`
	Options string `json:"options" validate:"required"`
}

type AnswerResponse struct {
	Text    string `json:"text"`
	Id      string `json:"id"`
	Options string `json:"options"`
}

type AnswerStatResponse struct {
	Text         string `json:"text"`
	Id           string `json:"id"`
	UserResponse int64  `json:"user_response"`
	Options      string `json:"options"`
	Correct      bool   `json:"correct"`
	YourAnswer   bool   `json:"your_answer"`
}

type AnswerRequest struct {
	Qid string `json:"quiz_id" validate:"required"`
	Aid string `json:"answer_id" validate:"required"`
}
