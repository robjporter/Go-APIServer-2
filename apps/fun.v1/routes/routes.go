package routes

import (
	"../controllers"

	"github.com/go-zoo/bone"
)

func Prefetch() {}

func GetRouteURL() string {
	return "/api/v1/fun"
}

func Routes() *bone.Mux {
	tmp := bone.New()

	tmp.GetFunc("/", controllers.Home)

	tmp.GetFunc("/logo/github", controllers.LogoGithub)
	tmp.GetFunc("/logo/greenapple", controllers.LogoGreenApple)
	tmp.GetFunc("/logo/cisco1", controllers.LogoCisco1)

	tmp.GetFunc("/animal/simpleowl", controllers.AnimalSimpleOwl)
	tmp.GetFunc("/animal/animals", controllers.AnimalAnimals)
	tmp.GetFunc("/animal/cat", controllers.AnimalCat)
	tmp.GetFunc("/animal/bear1", controllers.AnimalBear1)
	tmp.GetFunc("/animal/bear2", controllers.AnimalBear2)
	tmp.GetFunc("/animal/dog", controllers.AnimalDog)
	tmp.GetFunc("/animal/elephant", controllers.AnimalElephant)
	tmp.GetFunc("/animal/lion", controllers.AnimalLion)
	tmp.GetFunc("/animal/monkey", controllers.AnimalMonkey)
	tmp.GetFunc("/animal/penguin", controllers.AnimalPenguin)
	tmp.GetFunc("/animal/sheep", controllers.AnimalSheep)

	tmp.GetFunc("/object/clouds", controllers.ObjectClouds)
	tmp.GetFunc("/object/mackeyboard", controllers.ObjectMacKeyboard)

	tmp.GetFunc("/characters/minions", controllers.CharactersMinions)
	tmp.GetFunc("/characters/goofy", controllers.CharactersGoofy)
	tmp.GetFunc("/characters/coder", controllers.CharactersCoder)

	tmp.GetFunc("/characters/simpsons/apu", controllers.CharactersApu)
	tmp.GetFunc("/characters/simpsons/bart", controllers.CharactersBart)
	tmp.GetFunc("/characters/simpsons/comicbookguy", controllers.CharactersComic)
	tmp.GetFunc("/characters/simpsons/homer", controllers.CharactersHomer)
	tmp.GetFunc("/characters/simpsons/homer2", controllers.CharactersHomer2)
	tmp.GetFunc("/characters/simpsons/itchy", controllers.CharactersItchy)
	tmp.GetFunc("/characters/simpsons/krusty", controllers.CharactersKrusty)
	tmp.GetFunc("/characters/simpsons/lisa", controllers.CharactersLisa)
	tmp.GetFunc("/characters/simpsons/maggie", controllers.CharactersMaggie)
	tmp.GetFunc("/characters/simpsons/marge", controllers.CharactersMarge)
	tmp.GetFunc("/characters/simpsons/mrburns", controllers.CharactersMrBurns)
	tmp.GetFunc("/characters/simpsons/ned", controllers.CharactersNedFlanders)
	tmp.GetFunc("/characters/simpsons/ralph", controllers.CharactersRalphWiggum)
	tmp.GetFunc("/characters/simpsons/smithers", controllers.CharactersSmithers)

	return tmp
}
