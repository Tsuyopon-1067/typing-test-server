package src

import (
	"fmt"
	"time"
)

type RegisterScore struct {
    Keystrokes int       `json:"keystrokes"`
    Accuracy   float64   `json:"accuracy"`
    Score      float64   `json:"score"`
    StartedAt  time.Time `json:"startedAt"`
    EndedAt    time.Time `json:"endedAt"`
}

func (r *RegisterScore) Text() string {
	return fmt.Sprintf("Keystrokes: %d,\nAccuracy: %f\nScore: %f\nStartedAt: %s\nEndedAt: %s", r.Keystrokes, r.Accuracy, r.Score, r.StartedAt, r.EndedAt)
}