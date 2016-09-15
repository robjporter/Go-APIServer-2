package controllers

import (
	"net/http"

	"../../../render"
)

func TMP(w http.ResponseWriter, req *http.Request) {
	templates := render.GetBaseTemplates()
	templates = append(templates, "apps/fun.v1/views/home.html")
	//templates := []string{"apps/fun.v1/views/home.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Home(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("HOME!"))
}

func LogoGithub(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/github.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoGreenApple(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/greenapple.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func LogoCisco1(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/logo/cisco1.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalAnimals(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/animals.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalCat(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/cat.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalBear1(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/bear1.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalBear2(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/bear2.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalDog(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/dog.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalElephant(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/elephant.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalLion(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/lion.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalMonkey(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/monkey.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalPenguin(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/penguin.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalSheep(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/sheep.animal.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AnimalSimpleOwl(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/animal/simpleowl.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ObjectClouds(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/objects/clouds.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func ObjectMacKeyboard(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/objects/mackeyboard.object.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersGoofy(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/goofy.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersHomer2(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/homer2.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMinions(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/minions.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersCoder(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/coder.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersApu(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/apu.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersBart(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/bart.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersComic(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/comicbookguy.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersHomer(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/homer.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersItchy(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/itchy.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersKrusty(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/krusty.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersLisa(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/lisa.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMaggie(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/maggie.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMarge(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/marge.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersMrBurns(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/mrburns.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersNedFlanders(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/nedflanders.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersRalphWiggum(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/ralphwiggum.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func CharactersSmithers(w http.ResponseWriter, req *http.Request) {
	templates := []string{"apps/fun.v1/views/characters/smithers.characters.html"}
	err := render.RenderTemplate(w, templates, "base", map[string]string{"Title": "Home"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
