package calculate

func LimitPercent(req float64) float64 {
	if req > 100 {
		return 100
	}
	return req
}
