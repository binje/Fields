// Code generated by "stringer -type=Season"; DO NOT EDIT.

package time

import "strconv"

const _Season_name = "JunePreperationsNovemberInventoryingDecemberPreperationsMayInventorying"

var _Season_index = [...]uint8{0, 16, 36, 56, 71}

func (i Season) String() string {
	if i < 0 || i >= Season(len(_Season_index)-1) {
		return "Season(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Season_name[_Season_index[i]:_Season_index[i+1]]
}
