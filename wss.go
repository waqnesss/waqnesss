package main

import (
	"fmt"
)

type cow struct {
	voice string
}

type duck struct {
	voice string
	color string
}

type goat struct {
	voice string
	color string
}

type sheep struct {
	color string
	voice string
}

type jamal struct {
	nation string
	voice  string
}

type jaguar struct {
	voice string
	color string
	year  int
}

type dog struct {
	voice string
	color string
}

func (c *cow) setVoice(voice string) {
	c.voice = voice
}

func (d *duck) setVoice(voice string) {
	d.voice = voice
}

func (g *goat) setVoice(voice string) {
	g.voice = voice
}

func (s *sheep) setVoice(voice string) {
	s.voice = voice
}

func (j *jamal) setVoice(voice string) {
	j.voice = voice
}

func (j *jaguar) setVoice(voice string) {
	j.voice = voice
}

func (d *dog) setVoice(voice string) {
	d.voice = voice
}

func shepHerd(v animals, voice string) {
	v.setVoice(voice)
}

type animals interface {
	setVoice(voice string)
}

func main() {
	cow := cow{voice: "muuuuu"}
	duck := duck{color: "white", voice: "cra-cra"}
	goat := goat{color: "black", voice: "meeeeeee"}
	sheep := sheep{color: "brown", voice: "beeeeeee"}
	jamal := jamal{nation: "afganets", voice: "Allahu-Akbar"}
	jaguar := jaguar{color: "red", year: 2011, voice: "vvrrrr-ru-tu-tu-tu"}
	dog := dog{color: "blue", voice: "rraw-raw"}

	shepHerd(&cow, "\nмуууу")
	fmt.Println("\nКорова на месте!", cow.voice)

	shepHerd(&duck, "\nкря-кря")
	fmt.Println("\nУтка на месте!", duck.voice)

	shepHerd(&goat, "\nмеееее")
	fmt.Println("\nКоза на месте!", goat.voice)

	shepHerd(&sheep, "\nбеееее")
	fmt.Println("\nОвца на месте!", sheep.voice)

	shepHerd(&jamal, "\nАллаху-Акбар")
	fmt.Println("\nДжамал не отстает от животных!", jamal.voice)

	shepHerd(&jaguar, "\nвврррр-ру-ту-ту-ту")
	fmt.Println("\nЭто вообщето машина, хотя тож сойдет", jaguar.voice)

	shepHerd(&dog, "\nррав-рав")
	fmt.Println("\nСобака на месте!", dog.voice)
}
