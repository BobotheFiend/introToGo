package main

func main() {}

func SumSeries(seriesRange int) int {
	if seriesRange == 0 {
		return 0
	}

	return seriesRange + SumSeries(seriesRange-1)
}
