package utils

const maxCoordinate = 100.0

func ApplyVelocityComponent(pos, vel *float64, energyLoss, iterations float64){
	*pos += *vel / iterations

	if *pos < 0 {
		*pos = -*pos
		*vel = -*vel
	}

	if *pos > maxCoordinate {
		*pos = 2 * maxCoordinate - *pos
		*vel = -(*vel * energyLoss)
	}
}