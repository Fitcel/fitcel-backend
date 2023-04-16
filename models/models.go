package models

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type Model struct {
	DB *gorm.DB
}
type dietType uint

const (
	beginner dietType = iota
	moderate
	advanced
)

func (s dietType) String() string {
	switch s {
	case beginner:
		return "Beginner"
	case moderate:
		return "Moderate"
	case advanced:
		return "Advanced"
	}
	return "unknown" // under ideal circumstances we should never reach this.
}

func (s dietType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *dietType) UnmarshalJSON(data []byte) (err error) {
	var diet string
	if err := json.Unmarshal(data, &diet); err != nil {
		return err
	}
	if *s, err = func(dietString string) (dietType, error) {
		dietString = strings.TrimSpace(strings.ToLower(dietString))
		switch dietString {
		case "beginner":
			return beginner, nil
		case "moderate":
			return moderate, nil
		case "advanced":
			return advanced, nil
		default:
			return dietType(0), fmt.Errorf("%q is not a valid card suit", dietString)
		}
	}(diet); err != nil {
		return err
	}
	return nil
}
