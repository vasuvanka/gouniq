package gouniq

import "testing"

func TestUniqCodeSuccess(t *testing.T) {
	ID, err := NewID(10)
	if err != nil {
		t.Error(err)
	}
	if len(ID) != 10 {
		t.Errorf("id length mismatch - %d", len(ID))
	} else {
		t.Logf("success : %s", ID)
	}
}

func TestUniqCodeFailure(t *testing.T) {
	_, err := NewID(0)
	if err == nil {
		t.Error(err)
	} else {
		t.Logf("expected failure : %s", err.Error())
	}
}
