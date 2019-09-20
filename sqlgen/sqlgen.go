package sqlgen

var convertTimeToUTC = true

func ConvertTimeToUTC(toUTC bool) {
	convertTimeToUTC = toUTC
}
