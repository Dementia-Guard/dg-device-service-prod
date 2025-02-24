package models

type Coordinate struct {
	Latitude  float64 `json:"latitude" firestore:"latitude"`
	Longitude float64 `json:"longitude" firestore:"longitude"`
}

type Patient struct {
	Name            string      `json:"name" firestore:"name"`
	Email           string      `json:"email" firestore:"email"`
	Telephone       int64       `json:"telephone" firestore:"telephone"`
	CaregiverEmail  string      `json:"caregiverEmail" firestore:"caregiverEmail"`
	DgHomeId        string      `json:"dgHomeId" firestore:"dgHomeId"`
	DgWearId        string      `json:"dgWearId" firestore:"dgWearId"`
	BoundaryCoordinates []Coordinate `json:"boundaryCoordinates" firestore:"boundaryCoordinates"`
}