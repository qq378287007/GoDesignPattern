package main

import "fmt"

type Train interface {
	Arrive()
	Depart()
	PermitArrival()
}

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}

//货运列车类
type FreightTrain struct {
	Mediator Mediator
}

//火车到达
func (g *FreightTrain) Arrive() {
	if !g.Mediator.CanArrive(g) {
		fmt.Println("FreightTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("FreightTrain: Arrived")
}

//火车离开
func (g *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving")
	g.Mediator.NotifyAboutDeparture()
}

//允许到达
func (g *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted")
	g.Arrive()
}

//客运火车类
type PassengerTrain struct {
	Mediator Mediator
}

//火车到达
func (p *PassengerTrain) Arrive() {
	if !p.Mediator.CanArrive(p) {
		fmt.Println("PassengerTrain: Arrival blocked, waiting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

//火车离开
func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving")
	p.Mediator.NotifyAboutDeparture()
}

//允许到达
func (p *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, arriving")
	p.Arrive()
}

type StationManager struct {
	isPlatformFree bool
	trainQueue     []Train
}

func NewStationManger() *StationManager {
	return &StationManager{isPlatformFree: true}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.isPlatformFree {
		s.isPlatformFree = false
		return true
	}
	s.trainQueue = append(s.trainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.isPlatformFree {
		s.isPlatformFree = true
	}
	if len(s.trainQueue) > 0 {
		firstTrainInQueue := s.trainQueue[0]
		s.trainQueue = s.trainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}

func main() {
	//声明具体中介者
	stationManager := NewStationManger()

	//声明客运火车
	passengerTrain := &PassengerTrain{Mediator: stationManager}

	//声明货运火车
	freightTrain := &FreightTrain{Mediator: stationManager}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
