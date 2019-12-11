package day11

import (
	"testing"

	"github.com/EmperorLeo/adventofcode2019/mocks"
	"github.com/EmperorLeo/adventofcode2019/util"
	"github.com/golang/mock/gomock"
)

func TestGetPaintedPanels(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// Make a mock computer to pass into the function, because I already know my computer works
	// and /u/topaz2078 didn't provide different intcode instructions as tests
	testComputer := mocks.NewMockIComputer(mockCtrl)

	// Hooray for gomocks... woooo...
	// To be fair.  This is better than not testing at all LOL
	testComputer.EXPECT().Type(gomock.Any()).AnyTimes()
	testComputer.EXPECT().Run().Times(1)
	testComputer.EXPECT().Read().Return(1, true).Times(1)
	testComputer.EXPECT().Read().Return(0, true).Times(1)
	testComputer.EXPECT().Read().Return(1, true).Times(1)
	testComputer.EXPECT().Read().Return(0, true).Times(1)
	testComputer.EXPECT().Read().Return(0, true).Times(1)
	testComputer.EXPECT().Read().Return(1, true).Times(1)
	testComputer.EXPECT().Read().Return(0, true).Times(1)
	testComputer.EXPECT().Read().Return(1, true).Times(3)
	testComputer.EXPECT().Read().Return(0, true).Times(2)
	testComputer.EXPECT().Read().Return(0, false).Times(2)

	result := getPaintedPanels(testComputer, 0)
	if result[util.Coord{0, 0}].color != 1 {
		t.Log("Expected (0,0) to be white, got black.")
		t.Fail()
	}
	if result[util.Coord{-1, -1}].color != 0 {
		t.Log("Expected (-1,-1) to be black, got white.")
		t.Fail()
	}
	if result[util.Coord{-2, -1}].color != 0 {
		t.Log("Expected (0,0) to be black, got white.")
		t.Fail()
	}
	if result[util.Coord{-2, 0}].color != 1 {
		t.Log("Expected (0,0) to be white, got black.")
		t.Fail()
	}
	// This assertion ensures that
	if result[util.Coord{-1, -0}].color != 0 {
		t.Log("Expected (0,0) to be black, got white.")
		t.Fail()
	}
}

func TestMapPanels(t *testing.T) {
	panels := map[util.Coord]panel{
		util.Coord{-1, -2}: panel{1, 1},
		util.Coord{1, 1}:   panel{0, 1},
		util.Coord{-1, 0}:  panel{1, 1},
		util.Coord{0, -1}:  panel{1, 1},
		util.Coord{2, 2}:   panel{0, 1},
		util.Coord{-1, -1}: panel{1, 1},
	}
	grid := mapPanels(panels)

	expectations := map[util.Coord]int{
		util.Coord{0, 0}: 1,
		util.Coord{1, 3}: 0,
		util.Coord{0, 2}: 1,
		util.Coord{1, 1}: 1,
		util.Coord{3, 4}: 0,
		util.Coord{0, 1}: 1,
	}

	for k, v := range expectations {
		if grid[k.Y][k.X] != v {
			t.Logf("Got color %d at (%d, %d), expected %d\n", grid[k.Y][k.X], k.X, k.Y, v)
			t.Fail()
		}
	}
}
