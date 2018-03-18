package types

type SenseBox struct {
	Id              string           `json:"_id"`
	CreatedAt       JsonNullTime     `json:"createdAt"`
	UpdatedAt       JsonNullTime     `json:"updatedAt"`
	Name            JsonNullString   `json:"name"`
	CurrentLocation SenseBoxLocation `json:"currentLocation"`
	Exposure        JsonNullString   `json:"exposure"`
	Sensors         []SenseBoxSensor `json:"sensors"`
	Model           JsonNullString   `json:"model"`
	Description     JsonNullString   `json:"description"`
	Loc             []SenseBoxLoc    `json:"loc"`
}

type SenseBoxLocation struct {
	Timestamp   JsonNullTime    `json:"timestamp"`
	Coordinates []JsonNullFloat `json:"coordinates"`
	Type        JsonNullString  `json:"type"`
}

type SenseBoxLoc struct {
	Geometry SenseBoxGeometry `json:"geometry"`
	Type     JsonNullString   `json:"type"`
}

type SenseBoxGeometry struct {
	Timestamp   JsonNullTime    `json:"timestamp"`
	Coordinates []JsonNullFloat `json:"coordinates"`
	Type        JsonNullString  `json:"type"`
}

type SenseBoxSensor struct {
	Id              string              `json:"_id"`
	Title           JsonNullString      `json:"title"`
	Unit            JsonNullString      `json:"unit"`
	SensorType      JsonNullString      `json:"sensorType"`
	Icon            JsonNullString      `json:"icon"`
	LastMeasurement SenseBoxMeasurement `json:"lastMeasurement"`
}

type SenseBoxMeasurement struct {
	Value     JsonNullFloat `json:"value"`
	CreatedAt JsonNullTime  `json:"createdAt"`
}
