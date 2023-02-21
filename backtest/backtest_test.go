package backtest

import (
	"testing"
)

// test btnReadConf_Clicked

func TestBtnReadConf_Clicked(t *testing.T) {
	// test the output
	instr, bd, _, _, _, _ := BtnReadConf_Clicked("../config/Manual/")
	if instr != "sh510050" {

		t.Errorf("instr is %s, want sh510050", instr)
	}
	if bd != "2023.01.18T09:35:00.000" {
		t.Errorf("bd is %s, want  2023.01.18T09:35:00.000,", bd)
	}
	// test the error

}
