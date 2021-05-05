package sctructs

import "errors"

type Item struct {
	ID           int
	Name         string
	Category     string
	Scores       []Score
	AverageScore float64

	Cost   int
	Author string
}

func (i *Item) AddScore(userID, score int) error {
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
