// Code generated by "stringer -type=Tool"; DO NOT EDIT.

package main

import "strconv"

const _Tool_name = "FishTrapsFleshingBeamsWeavingLoomsSlaughteringTablesSpadesShovelPairsShovelsPotteryWheelsOvensAxesWorkbenches"

var _Tool_index = [...]uint8{0, 9, 22, 34, 52, 58, 69, 76, 89, 94, 98, 109}

func (i Tool) String() string {
	if i < 0 || i >= Tool(len(_Tool_index)-1) {
		return "Tool(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Tool_name[_Tool_index[i]:_Tool_index[i+1]]
}
