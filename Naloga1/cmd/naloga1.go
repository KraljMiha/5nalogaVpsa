package main

import (
	"context"
	f "fmt"
	"log"
	"os"

	"github.com/KraljMiha/5nalogaVpsa/redovalnica"

	"github.com/urfave/cli/v3"
)

func izpisRedovalnice(slovar map[string]redovalnica.Student) {
	f.Println("REDOVALNICA:")
	for vpisnaStevilka, _ := range slovar {
		f.Printf("%s - %s %s: ", vpisnaStevilka, slovar[vpisnaStevilka].Ime, slovar[vpisnaStevilka].Priimek)
		f.Print(slovar[vpisnaStevilka].Ocene, "\n")
	}

}

func main() {

	cmd := &cli.Command{
		Name:  "Redovalnica",
		Usage: "Obdeluje podatke študentov v redovalnici.",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "Definira najmanjše število ocen potrebnih za pozitivno oceno",
				Value: 5,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "Najmanjša možna ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "Največja možna ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")
			omejitve := redovalnica.NewOmejitve(stOcen, minOcena, maxOcena)
			return runRedovalnica(omejitve)
		},
	}
	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func runRedovalnica(omejitve redovalnica.Omejitve) error {
	var slovar = make(map[string]redovalnica.Student)

	slovar["63210001"] = redovalnica.Student{Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8, 7}}
	slovar["63210002"] = redovalnica.Student{Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8}}
	slovar["63210003"] = redovalnica.Student{Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3, 5}}

	redovalnica.DodajOceno("63210001", 9, slovar, omejitve)  // dodajanje ocene Ana	  2x
	redovalnica.DodajOceno("63210001", 9, slovar, omejitve)  // dodajanje ocene Ana
	redovalnica.DodajOceno("63210002", 10, slovar, omejitve) // dodajanje ocene Boris 3x
	redovalnica.DodajOceno("63210002", 10, slovar, omejitve) // dodajanje ocene Boris
	redovalnica.DodajOceno("63210002", 10, slovar, omejitve) // dodajanje ocene Boris
	redovalnica.DodajOceno("63210003", 8, slovar, omejitve)  // dodajanje ocene Janez 4x
	redovalnica.DodajOceno("63210003", 8, slovar, omejitve)  // dodajanje ocene Janez
	redovalnica.DodajOceno("63210003", 8, slovar, omejitve)  // dodajanje ocene Janez
	redovalnica.DodajOceno("63210003", 8, slovar, omejitve)  // dodajanje ocene Janez

	// Ana ima 6 ocen, Boris 7, Janez 8

	f.Println()
	f.Println("Izpis redovalnice po dodajanju:")
	izpisRedovalnice(slovar)
	f.Println()

	f.Println("Dodajanje ocen 3 in 12, glede na omejitve in neobstoječ:")
	redovalnica.DodajOceno("63210003", 3, slovar, omejitve)  // dodajanje ocene 3 Janezu
	redovalnica.DodajOceno("63210003", 12, slovar, omejitve) // dodajanje ocene 11 janezu
	redovalnica.DodajOceno("63210004", 7, slovar, omejitve)  // dodajanje ocene neobstoječemu

	f.Println()
	f.Println("Izpis končnega uspeha in ocen:")
	for vpisnaStevilka, _ := range slovar {
		redovalnica.IzpisVsehOcen(vpisnaStevilka, slovar)
		redovalnica.IzpisiKoncniUspeh(vpisnaStevilka, slovar, omejitve)
		f.Println()
	}

	return nil
}
