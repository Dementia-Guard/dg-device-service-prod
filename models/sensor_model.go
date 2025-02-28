package models

// SensorData represents the sensor data structure
type SensorData struct {
	// ID                    SensorID `bson:"_id " json:"id"`
	ID map[string]any `bson:"_id" json:"_id"`
	AccelerometerVariance float64 `bson:"accelerometer_variance" json:"accelerometer_variance"`
	Activity              string  `bson:"activity" json:"activity"`
	BloodOxygenAvg        float64 `bson:"blood_oxygen_avg" json:"blood_oxygen_avg"`
	BloodOxygenMax        float64 `bson:"blood_oxygen_max" json:"blood_oxygen_max"`
	BloodOxygenMin        float64 `bson:"blood_oxygen_min" json:"blood_oxygen_min"`
	GyroscopeVariance     float64 `bson:"gyroscope_variance" json:"gyroscope_variance"`
	PulseRateAvg          float64 `bson:"pulse_rate_avg" json:"pulse_rate_avg"`
	PulseRateMax          int     `bson:"pulse_rate_max" json:"pulse_rate_max"`
	PulseRateMin          int     `bson:"pulse_rate_min" json:"pulse_rate_min"`
	StepCountSum          int     `bson:"step_count_sum" json:"step_count_sum"`
	TemperatureAvg        float64 `bson:"temperature_avg" json:"temperature_avg"`
	TemperatureMax        float64 `bson:"temperature_max" json:"temperature_max"`
	TemperatureMin        float64 `bson:"temperature_min" json:"temperature_min"`
}
