package sctructs

import (
	"testing"
)

func TestUser_AddScore(t *testing.T) {
	tt := []struct {
		name  string
		item  User
		want  float64
		error string
		score int
	}{
		{
			name: "ok",
			item: User{
				Scores: []Score{{Score: 5}, {Score: 2}},
			},
			want:  3.6666666666666665,
			score: 4,
		},
		{
			name: "scores nil",
			item: User{
				Scores: nil,
			},
			want:  4,
			score: 4,
		},
		{
			name:  "adding to empty scores",
			item:  User{Scores: []Score{}},
			want:  4,
			score: 4,
		},
		{
			name:  "score less than 0",
			item:  User{Scores: []Score{{Score: 5}}},
			want:  0,
			error: "score cant be less than 0 or bigger than 5",
			score: -1,
		},
		{
			name:  "score bigger than 5",
			item:  User{Scores: []Score{{Score: 5}}},
			want:  0,
			error: "score cant be less than 0 or bigger than 5",
			score: 10,
		},
	}

	for _, v := range tt {
		t.Run(v.name, func(t *testing.T) {
			err := v.item.AddScore(0, v.score)
			if v.item.AverageScore != v.want || (err != nil && v.error != err.Error()) {
				t.Errorf("want '%v', got '%v'", v.want, v.item.AverageScore)
				t.Errorf("want error: '%v', got error: '%v'", v.error, err)
			}
		})
	}

}
