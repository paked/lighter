package core

import (
	"github.com/paked/engi"
	"github.com/paked/lighter/components"
	"github.com/paked/lighter/systems"
	"log"
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
}

func (l *Lighter) Setup() {
	rand.Seed(time.Now().UnixNano())
	log.Println("Welcome to lighter, a game made by @paked_ for Ludum Dare 31")
	l.AddSystem(&engi.RenderSystem{})
	l.AddSystem(&systems.ControlSystem{})
	l.AddSystem(&engi.CollisionSystem{})
	l.AddSystem(&systems.LightSystem{})
	l.AddSystem(&systems.GuardAISystem{})
	l.AddSystem(&systems.KeySystem{})
	l.AddSystem(&systems.VisionSystem{})
	l.AddSystem(&systems.StickySystem{})

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

	l.AddEntity(NewKey())
}

func NewPlayer() *engi.Entity {
	player := engi.NewEntity([]string{"RenderSystem", "ControlSystem", "CollisionSystem", "VisionSystem"})
	player.Pattern = "player"
	render := engi.NewRenderComponent(engi.Files.Image("player"), engi.Point{2, 2}, "player")
	space := engi.SpaceComponent{Position: engi.Point{400, 400}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	control := components.ControlComponent{Scheme: systems.CONTROL_SCHEME_WASD}
	speed := components.SpeedComponent{}
	collision := engi.CollisionComponent{Main: true, Extra: engi.Point{}, Solid: true}
	key := components.KeyComponent{}

	player.AddComponent(&render)
	player.AddComponent(&space)
	player.AddComponent(&control)
	player.AddComponent(&speed)
	player.AddComponent(&collision)
	player.AddComponent(&key)

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
	guard := engi.NewEntity([]string{"RenderSystem", "GuardAISystem", "VisionSystem"})
	guard.Pattern = "guard"
	render := engi.NewRenderComponent(engi.Files.Image("guard"), engi.Point{2, 2}, "guard")
	space := engi.SpaceComponent{Position: engi.Point{engi.Width() * rand.Float32(), 100}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	destination := components.DestinationComponent{}
	link := engi.LinkComponent{target}
	vision := components.VisionComponent{true, 0}
	guard.AddComponent(&render)
	guard.AddComponent(&space)
	guard.AddComponent(&link)
	guard.AddComponent(&destination)
	guard.AddComponent(&vision)

	sight := engi.NewEntity([]string{"RenderSystem", "StickySystem"})
	sight.Pattern = "sight"
	renderS := engi.NewRenderComponent(engi.Files.Image("sight"), engi.Point{2, 2}, "sight")
	spaceS := engi.SpaceComponent{Position: space.Position, Width: 64 * render.Scale.X, Height: 64 * render.Scale.Y}
	linkS := engi.LinkComponent{guard}
	sight.AddComponent(&renderS)
	sight.AddComponent(&spaceS)
	sight.AddComponent(&linkS)
	return guard, sight
}

func NewKey() *engi.Entity {
	key := engi.NewEntity([]string{"RenderSystem", "CollisionSystem", "KeySystem"})
	key.Pattern = "key"
	render := engi.NewRenderComponent(engi.Files.Image("key"), engi.Point{2, 2}, "guard")
	space := engi.SpaceComponent{Position: engi.Point{100, 100}, Width: 16 * render.Scale.X, Height: 16 * render.Scale.Y}
	link := engi.LinkComponent{}
	collision := engi.CollisionComponent{}
	// keyC := components.KeyComponent{}

	key.AddComponent(&render)
	key.AddComponent(&space)
	key.AddComponent(&link)
	// key.AddComponent(&keyC)
	key.AddComponent(&collision)
	return key
}
