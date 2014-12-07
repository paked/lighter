package core

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/systems"
	"math/rand"
	"time"
)

type Lighter struct {
	engi.World
}

func (l Lighter) Preload() {
	engi.Files.Add("player", "assets/player.png")
	engi.Files.Add("lightsource", "assets/lightsource.png")
	engi.Files.Add("shade", "assets/shade.png")
	engi.Files.Add("guard", "assets/enemy.png")
	engi.Files.Add("key", "assets/key.png")
	engi.Files.Add("sight", "assets/sight.png")
	engi.Files.Add("tileset", "assets/tileset.png")
	engi.Files.Add("playersheet", "assets/Hero.png")
	engi.Files.Add("arrows", "assets/arrows.png")
	engi.Files.Add("enemysheet", "assets/enemysheet.png")
}

func (l *Lighter) Setup() {
	rand.Seed(time.Now().UnixNano())
	l.AddSystem(&systems.KeySystem{})
	l.AddSystem(&engi.RenderSystem{})
	l.AddSystem(&engi.CollisionSystem{})
	l.AddSystem(&engi.AnimationSystem{})
	l.AddSystem(&systems.ControlSystem{})
	l.AddSystem(&systems.LightSystem{})
	l.AddSystem(&systems.GuardAISystem{})
	l.AddSystem(&systems.VisionSystem{})
	l.AddSystem(&systems.StickySystem{})
	l.AddSystem(&systems.PuzzleSystem{})

	w := int(engi.Height() / 32)
	h := int(engi.Width() / 32)
	array := make([][]string, w)
	for i := range array {
		array[i] = make([]string, h)
	}

	gameMap := engi.NewEntity([]string{"RenderSystem"})
	for y := 0; y < w; y++ {
		for x := 0; x < h; x++ {
			tile := "1"

			if rand.Float32() > .99 {
				tile = "2"
			}

			if rand.Float32() > .95 {
				tile = "3"
			}

			if rand.Float32() > .95 {
				tile = "4"
			}

			if y == 0 {
				tile = "5"
			}

			if x == 0 {
				tile = "6"
			}

			if x == (w - 1) {
				tile = "7"
			}

			if x == 0 && y == 0 {
				tile = "8"
			}

			if x == (w-1) && y == 0 {
				tile = "9"
			}

			array[y][x] = tile

		}
	}

	tilemap := engi.NewTilemap(array, engi.Files.Image("tileset"))
	mapRender := engi.NewRenderComponent(tilemap, engi.Point{2, 2}, "map")
	mapSpace := engi.SpaceComponent{engi.Point{0, 0}, 0, 0}
	gameMap.AddComponent(&mapRender)
	gameMap.AddComponent(&mapSpace)

	l.AddEntity(gameMap)

	l.AddEntity(NewPlayer())

	for x := float32(0); x < 1; x += .25 {
		for y := float32(0); y < 1; y += .25 {
			li, sh := NewLightAndShade(x*engi.Width()+64, y*engi.Height()+64)
			l.AddEntity(li)
			l.AddEntity(sh)
		}
	}

	for i := 0; i < 5; i++ {
		g, s := NewGuardAndSite(nil)
		l.AddEntity(g)
		l.AddEntity(s)
	}

	// l.AddEntity(NewKey())

	l.AddEntity(NewPuzzle())
}
func NewPuzzle() *engi.Entity {
	p := engi.NewEntity([]string{"RenderSystem", "PuzzleSystem"})
	render := engi.NewRenderComponent(nil, engi.Point{2, 2}, "puzzle")
	space := engi.SpaceComponent{Position: engi.Point{400 - 32, 400 - 32}, Width: 32, Height: 32}
	puzzle := components.PuzzleComponent{}

	p.AddComponent(&render)
	p.AddComponent(&space)
	p.AddComponent(&puzzle)

	return p
}
func NewPlayer() *engi.Entity {
	player := engi.NewEntity([]string{"RenderSystem", "ControlSystem", "CollisionSystem", "VisionSystem", "AnimationSystem"})
	player.Pattern = "player"
	spritesheet := engi.NewSpritesheet("playersheet", 16)
	render := engi.NewRenderComponent(spritesheet.Cell(0), engi.Point{2, 2}, "player")
	space := engi.SpaceComponent{Position: engi.Point{400, 400}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	control := components.ControlComponent{Scheme: systems.CONTROL_SCHEME_WASD}
	speed := components.SpeedComponent{}
	collision := engi.CollisionComponent{Main: true, Extra: engi.Point{}, Solid: true}
	key := components.KeyComponent{}

	animation := engi.NewAnimationComponent()
	animation.Rate = .2
	animation.S = spritesheet

	animation.AddAnimation("default", []int{4})
	animation.AddAnimation("up", []int{0, 1, 2, 3})
	animation.AddAnimation("down", []int{4, 5, 6, 7})
	animation.AddAnimation("left", []int{8, 9, 10, 11})
	animation.AddAnimation("right", []int{8, 9, 10, 11})
	animation.SelectAnimation("default")

	player.AddComponent(&render)
	player.AddComponent(&space)
	player.AddComponent(&control)
	player.AddComponent(&speed)
	player.AddComponent(&collision)
	player.AddComponent(&key)
	player.AddComponent(animation)

	return player
}

func NewLightAndShade(x, y float32) (*engi.Entity, *engi.Entity) {
	light := engi.NewEntity([]string{"RenderSystem", "CollisionSystem", "LightSystem"})
	light.Pattern = "light"
	render := engi.NewRenderComponent(engi.Files.Image("lightsource"), engi.Point{2, 2}, "light")

	offset := engi.Point{rand.Float32() * (engi.Width() / 20), rand.Float32() * (engi.Height() / 20)}
	if rand.Float32() > .5 {
		offset.X *= -1
	}

	if rand.Float32() > .5 {
		offset.Y *= -1
	}

	space := engi.SpaceComponent{Position: engi.Point{x + offset.X, y + offset.Y}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	collision := engi.CollisionComponent{Main: false, Solid: false}
	light.AddComponent(&render)
	light.AddComponent(&space)
	light.AddComponent(&collision)

	shade := engi.NewEntity([]string{"RenderSystem"})
	shade.Pattern = "shade"
	texture := engi.Files.Image("shade")
	shadeRender := engi.NewRenderComponent(texture, engi.Point{1, 1}, "shade")
	shadeSpace := engi.SpaceComponent{Position: engi.Point{(space.Position.X + space.Width/2) - (texture.Width() / 2), (space.Position.Y + space.Height/2) - (texture.Height() / 2)}, Width: texture.Width(), Height: texture.Height()}
	shade.Exists = false
	shade.AddComponent(&shadeRender)
	shade.AddComponent(&shadeSpace)

	link := engi.LinkComponent{shade}
	light.AddComponent(&link)
	return light, shade
}

func NewGuardAndSite(target *engi.Entity) (*engi.Entity, *engi.Entity) {
	guard := engi.NewEntity([]string{"RenderSystem", "GuardAISystem", "AnimationSystem"})
	guard.Pattern = "guard"
	s := engi.NewSpritesheet("enemysheet", 16)
	render := engi.NewRenderComponent(s.Cell(0), engi.Point{2, 2}, "guard")
	space := engi.SpaceComponent{Position: engi.Point{engi.Width() * rand.Float32(), 100}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	destination := components.DestinationComponent{}
	link := engi.LinkComponent{target}
	animation := engi.NewAnimationComponent()
	animation.Rate = .2
	animation.S = s

	animation.AddAnimation("default", []int{0})
	animation.AddAnimation("down", []int{0, 1})
	animation.AddAnimation("left", []int{2, 3})
	animation.AddAnimation("right", []int{4, 5})
	animation.AddAnimation("up", []int{6, 7})
	animation.SelectAnimation("default")

	guard.AddComponent(&render)
	guard.AddComponent(&space)
	guard.AddComponent(&link)
	guard.AddComponent(&destination)
	guard.AddComponent(&components.SpeedComponent{})
	guard.AddComponent(animation)

	sight := engi.NewEntity([]string{"RenderSystem", "StickySystem", "AnimationSystem", "VisionSystem"})
	sight.Pattern = "sight"
	spritesheet := engi.NewSpritesheet("sight", 64)
	renderS := engi.NewRenderComponent(spritesheet.Cell(0), engi.Point{2, 2}, "sight")
	spaceS := engi.SpaceComponent{Position: space.Position, Width: 64 * render.Scale.X, Height: 64 * render.Scale.Y}
	linkS := engi.LinkComponent{guard}
	vision := components.VisionComponent{true, 0}
	animationS := engi.NewAnimationComponent()
	animationS.Rate = .1
	animationS.S = spritesheet

	animationS.AddAnimation("default", []int{4})
	animationS.AddAnimation("attack", []int{0, 1, 2, 3, 2, 1, 0})
	animationS.SelectAnimation("default")

	sight.AddComponent(&renderS)
	sight.AddComponent(&spaceS)
	sight.AddComponent(&linkS)
	sight.AddComponent(animationS)
	sight.AddComponent(&vision)
	return guard, sight
}

func NewKey() *engi.Entity {
	key := engi.NewEntity([]string{"RenderSystem", "CollisionSystem", "KeySystem"})
	key.Pattern = "key"
	render := engi.NewRenderComponent(engi.Files.Image("key"), engi.Point{2, 2}, "guard")
	space := engi.SpaceComponent{Position: engi.Point{100, 100}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	link := engi.LinkComponent{}
	collision := engi.CollisionComponent{}

	key.AddComponent(&render)
	key.AddComponent(&space)
	key.AddComponent(&link)
	key.AddComponent(&collision)
	return key
}
