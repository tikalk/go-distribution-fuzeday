package models

import (
	"fmt"
	"go-distribution-fuzeday/utils"
	"math"
	"math/rand"
	"sync"
	"time"
)

const kickThreshold = 6
const kickVelocityThreshold = 4

type Player struct {
	X      float64
	Y      float64
	TeamID Team
	ID     string
	Name   string

	MaxVelocity float64
	LastKick    time.Time

	ball *Ball

	ballChannel chan *Ball // TODO Challenge: replace with directional input and output channels (<-chan and chan<-)

	idleV     float64
	idleVx    float64
	idleVy    float64
	idleAngle float64
}

func (p *Player) GetDisplayStatus() *DisplayStatus {
	res := &DisplayStatus{}
	res.X = p.X
	res.Y = p.Y
	res.ItemID = p.ID
	res.ItemLabel = p.Name
	res.TeamID = p.TeamID
	res.ItemType = TypePlayer
	res.LastUpdated = time.Now()

	return res
}

func reportDisplay(item DisplayStatusProvider, channel chan *DisplayStatus) {
	if channel == nil || item == nil {
		return
	}

	channel <- item.GetDisplayStatus()
}

func (p *Player) Activate(displayChannel chan *DisplayStatus, wg *sync.WaitGroup) {

	p.ballChannel = GetBallChannel()

	go p.setIdleKinematics()

	// Closing distance to ball
	// TODO Challenge: launch a goroutine that calls p.runToBall every 200 milliseconds or so...

	// reporting player display
	// TODO Challenge: launch a goroutine that calls reportDisplay() every 200 milliseconds or so...

	// launching main life cycle
	// TODO Challenge: call p.mainLifeCycle in a goroutine and implement it internally

}

func (p *Player) setIdleKinematics() {
	nextDelay := 0 * time.Second
	for {
		select {
		case <-time.After(nextDelay):

			p.idleV = 0.5 + 0.5*rand.Float64()
			p.idleAngle = math.Pi * 2 * rand.Float64()
			p.idleVx = math.Cos(p.idleAngle) * p.idleV
			p.idleVy = math.Sin(p.idleAngle) * p.idleV
			nextDelay = time.Duration(5.0+rand.Float64()*6.0) * time.Second
		}
	}
}

func (p *Player) mainLifeCycle(displayChannel chan *DisplayStatus, wg *sync.WaitGroup) {

	// TODO Tip: a ticker returns a channel that is automatically populated with a time message every defined interval
	//ticker := time.NewTicker(10 * time.Second)

	//TODO Challenge:
	// 1. iterate endlessly
	// 2. consume from ball channel
	// 3. decide if player is able to kick and applyKick, otherwise sleep for 20ms and applyKinematics to the ball
	// 4. reportDisplay and publish ball back to the channel
	// ----------------
	// * Pay attention to an initial delay before game starts (preferred, for distributed queues to initiate)
	// * Bonus: if waiting for more than 30 seconds for the ball message, check if player ever got the ball. If no, log and wait for another 30 seconds. If not - assume another player got killed with the ball, and throw another one to the channel
	// * consider utilize "select-case" mechanism

	wg.Done()
}

func (p *Player) getDistanceToBall(ball *Ball) float64 {
	return math.Sqrt(math.Pow(p.X-ball.X, 2) + math.Pow(p.Y-ball.Y, 2))
}

func (p *Player) runToBall() {

	// TODO Algorithm: make view threshold (50) random so that a distant player sees that ball after a period of time

	// once every N seconds - the player gets a longer view and can see the ball. Once saw the ball -
	// he keeps the "long view" mode for a longer period

	if p.ball != nil {
		dist := p.getDistanceToBall(p.ball)
		if dist < 30 && time.Now().Sub(p.LastKick) > 8*time.Second {
			vel := 0.05 + rand.Float64()*p.MaxVelocity
			p.X += (p.ball.X - p.X) * vel
			p.Y += (p.ball.Y - p.Y) * vel
		} else {
			p.idleMovement()
		}
		p.log(fmt.Sprintf("Current Position: (%f, %f), Ball Position: (%f, %f)", p.X, p.Y, p.ball.X, p.ball.Y))

	} else {
		p.idleMovement()
		p.log(fmt.Sprintf("Current Position: (%f, %f), No ball...", p.X, p.Y))

	}

}

func (p *Player) idleMovement() {
	utils.ApplyVelocityComponent(&p.X, &p.idleVx, 1, 1)
	utils.ApplyVelocityComponent(&p.Y, &p.idleVy, 1, 1)
}

func (p *Player) log(message string) {
	if message[0:1] == "\n" {
		message = message[1:]
		fmt.Printf("\n%s (%s): %s", p.Name, p.TeamID, message)
	} else {
		fmt.Printf("\r%s (%s): %s", p.Name, p.TeamID, message)
	}
}

func (p *Player) applyKick() {
	rand.Seed(time.Now().UnixNano())
	angle := 2 * math.Pi * rand.Float64()

	// TODO Algorithm: put the ball NEAR the threshold so sometimes he might re-kick the ball
	p.ball.X = p.X + 1.1*kickThreshold*math.Cos(angle)
	p.ball.Y = p.Y + 1.1*kickThreshold*math.Sin(angle)

	v := 1 + rand.Float64()*2
	p.ball.Vx = v * math.Cos(angle)
	p.ball.Vy = v * math.Sin(angle)
	p.ball.HolderID = p.ID
	p.ball.HolderTeam = p.TeamID
	p.ball.LastKick = time.Now()

	p.LastKick = time.Now()

	p.log(fmt.Sprintf("\nKick!!! (angle: %d degrees, velocity: %f)\n", int(180*angle/math.Pi), v))
}
