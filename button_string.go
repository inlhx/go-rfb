// Code generated by "stringer -type=Button"; DO NOT EDIT.

package rfb

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BtnLeft-1]
	_ = x[BtnMiddle-2]
	_ = x[BtnRight-4]
	_ = x[BtnFour-8]
	_ = x[BtnFive-16]
	_ = x[BtnSix-32]
	_ = x[BtnSeven-64]
	_ = x[BtnEight-128]
	_ = x[BtnNone-0]
}

const (
	_Button_name_0 = "BtnNoneBtnLeftBtnMiddle"
	_Button_name_1 = "BtnRight"
	_Button_name_2 = "BtnFour"
	_Button_name_3 = "BtnFive"
	_Button_name_4 = "BtnSix"
	_Button_name_5 = "BtnSeven"
	_Button_name_6 = "BtnEight"
)

var (
	_Button_index_0 = [...]uint8{0, 7, 14, 23}
)

func (i Button) String() string {
	switch {
	case i <= 2:
		return _Button_name_0[_Button_index_0[i]:_Button_index_0[i+1]]
	case i == 4:
		return _Button_name_1
	case i == 8:
		return _Button_name_2
	case i == 16:
		return _Button_name_3
	case i == 32:
		return _Button_name_4
	case i == 64:
		return _Button_name_5
	case i == 128:
		return _Button_name_6
	default:
		return "Button(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
