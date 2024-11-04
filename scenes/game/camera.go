package game

import (
	"space-shooter/config"
	"space-shooter/scenes/game/component"
)

type Camera struct {
	X      float64
	Y      float64
	config *config.AppConfig
}

func NewCamera(x, y float64, config *config.AppConfig) *Camera {
	return &Camera{
		X:      x,
		Y:      y,
		config: config,
	}
}

func (self *Camera) FocusTarget(target component.PositionData) {
	self.X = -target.X + float64(self.config.ScreenWidth)/2.0
	self.Y = -target.Y + float64(self.config.ScreenHeight)/2.0
}