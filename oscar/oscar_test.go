package oscar_test

import (
	"hello/oscar"
	"testing"
)

func TestActorWhoGotMoreThanOne(t *testing.T) {
	t.Run("oscar more than one", func(t *testing.T) {
		oscar.ActorWhoGotMoreThanOne("./oscar_age_male.csv")
	})
}
