package action

import (
	"reflect"
	"testing"
)

func TestNextActionsCanUseChoice(t *testing.T) {
	expectedActions := Choices{[]Action{GetSheep, GetCow, GetHorse, GetTimber, GetBrick}}
	if !reflect.DeepEqual(Grocer.NextActions(), expectedActions) {
		t.Errorf("did not get grocer actions")
	}
}

func TestChoiceCanAccepAnyAnimal(t *testing.T) {
	c := choices(getAnyAnimal())
	if !reflect.DeepEqual(choices(GEtShhep, GetCow, GetHorse), expectedActions) {
		t.Errorf("did not get grocer actions")
	}
}
