// main.go
// https://www.codingame.com/ide/puzzle/mad-pod-racing
// main is the main package for the racing coding game game
package main

import (
	"fmt"
	"math"
)

const (
	checkpointRadius    = 600
	podCollisionRadius  = 400
	boostExecuteKeyword = "BOOST"
	boostCount          = 1
	friction            = 0.85
)

var headingError = 90

// Vec2 represents a 2D vector
type Vec2 struct {
	x, y float64
}

// NewVec2 creates a new 2D vector
func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

// Add adds two 2D vectors
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.x + other.x, v.y + other.y}
}

// Sub subtracts two 2D vectors
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.x - other.x, v.y - other.y}
}

// Dot returns the dot product of two 2D vectors
func (v Vec2) Dot(other Vec2) float64 {
	return v.x*other.x + v.y*other.y
}

// Multiply multiplies two 2D vectors
func (v Vec2) Multiply(other Vec2) Vec2 {
	return Vec2{v.x * other.x, v.y * other.y}
}

// Cross returns the cross product of two 2D vectors
func (v Vec2) Cross(other Vec2) float64 {
	return v.x*other.y - v.y*other.x
}

// Divide divides two 2D vectors
func (v Vec2) Divide(other Vec2) Vec2 {
	return Vec2{v.x / other.x, v.y / other.y}
}

// Magnitude returns the magnitude of a 2D vector
func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.Dot(v))
}

// Equal returns true if two 2D vectors are equal
func (v Vec2) Equal(other Vec2) bool {
	return v.x == other.x && v.y == other.y
}

// NotEqual returns true if two 2D vectors are not equal
func (v Vec2) NotEqual(other Vec2) bool {
	return v.x != other.x || v.y != other.y
}

// Angle returns the angle of a 2D vector
func (v Vec2) Angle() float64 {
	return math.Atan2(v.y, v.x)
}

// Normalize returns a normalized 2D vector
func (v Vec2) Normalize() Vec2 {
	mag := v.Magnitude()
	if mag == 0 {
		return Vec2{0, 0}
	}
	return Vec2{v.x / mag, v.y / mag}
}

// Normal returns the normal vector of a 2D vector
func (v Vec2) Normal() Vec2 {
	mag := v.Magnitude()
	return Vec2{-v.y / mag, v.x / mag}
}

// Distance returns the distance between two 2D vectors
func (v Vec2) Distance(other Vec2) float64 {
	dx := v.x - other.x
	dy := v.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

// String returns a string representation of a 2D vector
func (v Vec2) String() string {
	return fmt.Sprintf("%d %d", int(v.x), int(v.y))
}

// StateTrack represents the state of the track
type StateTrack struct {
	currentLap          int
	checkPointPositions []Vec2
	allCheckpointsFound bool
}

// NewStateTrack creates a new StateTrack
func NewStateTrack() StateTrack {
	return StateTrack{
		currentLap:          0,
		checkPointPositions: make([]Vec2, 3),
		allCheckpointsFound: false,
	}
}

// Update updates the state of the track
func (s *StateTrack) Update(x, y int) {
	if s.allCheckpointsFound {
		return
	}

	newCheckpoint := NewVec2(float64(x), float64(y))
	if len(s.checkPointPositions) == 0 {
		s.checkPointPositions = append(s.checkPointPositions, newCheckpoint)
	} else if !s.checkPointPositions[len(s.checkPointPositions)-1].Equal(newCheckpoint) {
		if s.checkPointPositions[0].Equal(newCheckpoint) {
			s.allCheckpointsFound = true
			s.currentLap++
		} else {
			s.checkPointPositions = append(s.checkPointPositions, newCheckpoint)
		}
	}
}

// Boost represents the state of the boost
type Boost struct {
	boostsRemaining int
}

// NewBoost creates a new Boost
func NewBoost() Boost {
	return Boost{boostsRemaining: boostCount}
}

// TryBoosting tries to boost
func (b *Boost) TryBoosting() bool {
	if b.boostsRemaining <= 0 {
		return false
	}

	b.boostsRemaining--
	return true
}

func main() {
	state := NewStateTrack()
	boost := NewBoost()
	for {
		// Read input values
		var x, y, nextCheckpointX, nextCheckpointY, nextCheckpointDist, nextCheckpointAngle int
		fmt.Scan(&x, &y, &nextCheckpointX, &nextCheckpointY, &nextCheckpointDist, &nextCheckpointAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)

		state.Update(nextCheckpointX, nextCheckpointY)

		vector := NewVec2(float64(nextCheckpointX), float64(nextCheckpointY))
		// Thrusting
		if boost.TryBoosting() {
			fmt.Println(vector.String(), boostExecuteKeyword)
		} else {
			var thrust int
			heading := math.Abs(float64(nextCheckpointAngle))
			if heading > float64(headingError) {
				thrust = 0
			} else {
				thrust = 100
			}
			fmt.Println(vector.String(), thrust)
		}
	}
}
