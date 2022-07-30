package time

import "fmt"

type Calendar struct {
	month         month
	halfYear      int
	endOfTheWorld int
}

func NewCalendar() *Calendar {
	return &Calendar{
		month:         july,
		halfYear:      1,
		endOfTheWorld: 9,
	}
}

type Season int

const (
	JunePreperations Season = iota
	NovemberInventorying
	DecemberPreperations
	MayInventorying
)

// TODO is month exported?
type month int

const (
	july month = iota
	august
	september
	october
	november
	january
	february
	march
	april
	may
)

func (c Calendar) printDate() {
	fmt.Printf("\nYear:%d-%s\n", c.halfYear, c.month)
}

func (c Calendar) EndOfTheWorld() bool {
	return c.halfYear > c.endOfTheWorld
}

func (c *Calendar) NextMonth() {
	c.month++
	if c.month > may {
		c.month = july
	}
	c.printDate()
}

func (c Calendar) Season() Season {
	switch c.month {
	case july, august, september, october:
		return JunePreperations
	case november:
		return NovemberInventorying
	case january, february, march, april:
		return DecemberPreperations
	case may:
		return MayInventorying
	default:
		panic("Season unknown")
	}
}
