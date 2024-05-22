module github.com/go-go

go 1.22.3

require (
	github.com/faiface/pixel v0.10.0
	golang.org/x/image v0.13.0
	  github.com/go-go/game v0.0.0
        github.com/go-go/game/graphics v0.0.0
        github.com/go-go/game/models v0.0.0
)

require (
	github.com/faiface/glhf v0.0.0-20211013000516-57b20770c369 // indirect
	github.com/faiface/mainthread v0.0.0-20171120011319-8b78f0a41ae3 // indirect
	github.com/go-gl/gl v0.0.0-20211210172815-726fda9656d6 // indirect
	github.com/go-gl/glfw/v3.3/glfw v0.0.0-20221017161538-93cebf72946b // indirect
	github.com/go-gl/mathgl v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
)
replace github.com/go-go/game => ./game
replace github.com/go-go/game/graphics => ./game/graphics
replace github.com/go-go/game/models => ./game/models