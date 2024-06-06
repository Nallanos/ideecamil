package calcul

type Vecteur struct {
	X float64
	Y float64
}

// Structure pour le point d'arrivée
type ArrivalPoint struct {
	X float64
	Y float64
}

// Structure principale qui encapsule Vecteur et ArrivalPoint
type FinalInput struct {
	Vecteur      Vecteur
	ArrivalPoint ArrivalPoint
}

type Result struct {
	X float64
	Y float64
}

func Operation(body FinalInput) Result {
	result := Result{
		X: body.ArrivalPoint.X - body.Vecteur.X,
		Y: body.ArrivalPoint.Y - body.Vecteur.Y,
	}

	return result
}
