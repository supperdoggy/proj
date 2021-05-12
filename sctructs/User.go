package sctructs

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID           int `json:"id"`
	Username     string `gorm:"unique;not null" json:"username"`
	Email        string `gorm:"unique;not null" json:"email"`
	HashedPass   string `gorm:"unique;not null" json:"hashed_pass"`
	Scores       []Score `json:"scores"`
	AverageScore float64 `json:"average_score"`

	BirthDate time.Time `json:"birth_date"`
}

func (i *User) AddScore(userID, score int) error {
	if score < 0 || score > 5 {
		return errors.New("score cant be less than 0 or bigger than 5")
	}
	scoreStruct := Score{
		Score:  uint8(score),
		UserID: userID,
		ItemID: i.ID,
	}
	i.Scores = append(i.Scores, scoreStruct)

	var avrg float64
	for _, v := range i.Scores {
		avrg += float64(v.Score)
	}
	avrg /= float64(len(i.Scores))
	i.AverageScore = avrg
	return nil
}
