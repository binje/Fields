package time

type Month int

const (
	July Month = iota
	August
	September
	October
	November
	January
	February
	March
	April
	May
)

type Season int

const (
	JunePreperations Season = iota
	NovemberInventorying
	DecemberPreperations
	MayInventorying
)

func (m Month) Season() Season {
	switch m {
	case July, August, September, October:
		return JunePreperations
	case November:
		return NovemberInventorying
	case January, February, March, April:
		return DecemberPreperations
	case May:
		return MayInventorying
	default:
		panic("Season unknown")
	}
}
