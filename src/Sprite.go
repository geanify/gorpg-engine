package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	texture          *sdl.Texture
	frame            *sdl.Rect
	maxFrames        uint
	animations       map[string]*Animation
	frameIndex       uint
	currentAnimation string
}

func (sprite *Sprite) nextFrame() {
	currentAnimation := sprite.animations[sprite.currentAnimation]
	currentAnimation.nextFrame()
}

func (sprite *Sprite) setAnimation(animationName string) {
	sprite.currentAnimation = animationName
}

func (sprite *Sprite) getFrame() *sdl.Rect {
	currentAnimation := sprite.animations[sprite.currentAnimation]
	return currentAnimation.getFrame()
}

func (sprite *Sprite) setTextureColorMode(r uint8, g uint8, b uint8) {
	sprite.texture.SetColorMod(r, g, b)
}

func (sprite *Sprite) SetBlendModeAdd() {
	sprite.texture.SetBlendMode(sdl.BLENDMODE_ADD)
}

func (sprite *Sprite) SetBlendModeBlend() {
	sprite.texture.SetBlendMode(sdl.BLENDMODE_BLEND)
}

func (sprite *Sprite) SetBlendModeMod() {
	sprite.texture.SetBlendMode(sdl.BLENDMODE_MOD)
}
