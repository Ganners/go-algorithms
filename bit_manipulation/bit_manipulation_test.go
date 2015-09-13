package main

import (
	"strconv"
	"testing"
)

func TestInsertInto(t *testing.T) {

	fixture := (1 << 10) | (1 << 6) | (1 << 3) | (1 << 2)
	out, err := insertInto(
		(1<<10)|(1<<4)|(1<<5), // Set 4th and 5th bit to prove they get overwritten
		(1<<4)|(1<<1)|(1<<0),
		2,
		6,
	)
	if err != nil {
		t.Errorf("Did not expect error, received: %s", err)
	}

	if out != fixture {
		t.Errorf("Output %s did not match expected %s",
			strconv.FormatInt(int64(out), 2),
			strconv.FormatInt(int64(fixture), 2),
		)
	}
}
