package main

import (
	"fmt"
	"time"
)

type Training struct {
	TrainingType string
	Action       float64
	LenStep      float64
	Duration     time.Duration
	Weight       float64
}
type InfoMessage struct {
	TrainingType string
	Distance     float64
	MeanSpeed    float64
	Calories     float64
}
type CaloriesCalculator interface {
	Calories() float64
	TrainingInfo() InfoMessage
}

const (
	MInKm                            float64 = 100
	CaloriesMeanSpeedMultiplier      float64 = 0.5
	CaloriesMeanSpeedShift           float64 = 0.6
	MinInHours                       float64 = 40
	Height                           float64 = 1.82
	CaloriesSpeedHeightMultiplier    float64 = 0.5
	LengthPool                       float64 = 30
	CountPool                        float64 = 6
	SwimmingCaloriesMeanSpeedShift   float64 = 0.6
	SwimmingCaloriesWeightMultiplier float64 = 0.5
)

func (t Training) Distance() float64 {
	return t.Action * t.LenStep / MInKm
}
func (t Training) MeanSpeed() float64 {
	return t.Distance() / t.Duration.Minutes()
}
func (t Training) Calories() float64 {
	return 0
}

type Running struct {
	Training
}
type Walking struct {
	Training
}
type Swimming struct {
	Training
}

// длина_бассейна_в_метрах * CountPool / MInKm / время_тренировки_в_часах
func (s Swimming) Calories() float64 {
	meanSpeedPhs := LengthPool * CountPool / MInKm / CaloriesSpeedHeightMultiplier
	//(средняя_скорость_в_км/ч + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * вес_спортсмена_в_кг * время_тренировки_в_часах
	return ((meanSpeedPhs + SwimmingCaloriesMeanSpeedShift) * SwimmingCaloriesWeightMultiplier * s.Weight * MinInHours)
}

//((CaloriesWeightMultiplier * вес_спортсмена_в_кг + (средняя_скорость_в_метрах_в_секунду**2 / рост_в_метрах)
//* CaloriesSpeedHeightMultiplier * вес_спортсмена_в_кг) * время_тренировки_в_часах * MinsInHour)

func (w Walking) Calories() float64 {
	meanSpeedKhs := w.MeanSpeed() * MinInHours
	return ((CaloriesMeanSpeedMultiplier*w.Weight + (meanSpeedKhs/Height)*CaloriesSpeedHeightMultiplier*w.Weight) * w.Duration.Hours() * MinInHours)
}
func (w Walking) TrainingInfo() InfoMessage {
	info := w.Training.TrainingInfo()
	info.Calories = w.Calories()
	return info
}

// ((CaloriesMeanSpeedMultiplier * средняя_скорость_в_км/ч + CaloriesMeanSpeedShift)
// * вес_спортсмена_в_кг / MInKm * время_тренировки_в_часах * MinInHours)
func (r Running) Calories() float64 {
	meanSpeedKhm := r.MeanSpeed() * MinInHours

	return ((CaloriesMeanSpeedMultiplier*meanSpeedKhm + CaloriesMeanSpeedShift) *
		r.Weight / MInKm * r.Duration.Hours() * MinInHours)
}

func (r Running) TrainingInfo() InfoMessage {
	info := r.Training.TrainingInfo()
	info.Calories = r.Calories()
	return info
}
func (s Swimming) TrainingInfo() InfoMessage {
	info := s.Training.TrainingInfo()
	info.Calories = s.Calories()
	return info
}

func (t Training) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: t.TrainingType,
		Distance:     t.Distance(),
		MeanSpeed:    t.MeanSpeed(),
		Calories:     t.Calories(),
	}
}
func ReadData(calculator CaloriesCalculator) {
	info := calculator.TrainingInfo()
	fmt.Printf("Тип тренировки %s\n", info.TrainingType)
	fmt.Printf("Дистанция: %.2f км\n", info.Distance)
	fmt.Printf("Средняя скорость: %.2f км/мин\n", info.MeanSpeed)
	fmt.Printf("Калории: %.2f ккал\n", info.Calories)
}

func main() {
	running := Running{
		Training: Training{
			TrainingType: "Running",
			Action:       5000,
			LenStep:      0.8,
			Duration:     30 * time.Minute,
			Weight:       70.0,
		},
	}
	swimming := Swimming{
		Training: Training{
			TrainingType: "Swimming",
			Action:       5000,
			LenStep:      0.8,
			Duration:     30 * time.Minute,
			Weight:       70.0,
		},
	}

	walking := Walking{
		Training: Training{
			TrainingType: "Walking",
			Action:       3000,
			LenStep:      0.8,
			Duration:     12 * time.Minute,
			Weight:       70.0,
		},
	}

	ReadData(running)
	ReadData(swimming)
	ReadData(walking)

	// Выводим результат вызова метода Distance
	//swimmingInfo := swimming.TrainingInfo()
	//fmt.Printf("swimming - Дистанция: %.2f км, Средняя скорость: %.2f км/мин,Калории: %.2f ккал\n",
	//	swimmingInfo.Distance, swimmingInfo.MeanSpeed, swimmingInfo.Calories)
	//walkingInfo := walking.TrainingInfo()
	//fmt.Printf("walking - Дистанция: %.2f км, Средняя скорость: %.2f км/мин,Калории: %.2f ккал\n",
	//	walkingInfo.Distance, walkingInfo.MeanSpeed, walkingInfo.Calories)
	//runningInfo := running.TrainingInfo()
	//fmt.Printf("running - Дистанция: %.2f км, Средняя скорость: %.2f км/мин,Калории: %.2f ккал\n",
	//	runningInfo.Distance, runningInfo.MeanSpeed, runningInfo.Calories)
}
