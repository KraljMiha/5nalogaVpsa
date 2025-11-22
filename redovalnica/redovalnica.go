package redovalnica

import (
	f "fmt"
)

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

type Omejitve struct {
	StOcen   int
	MinOcena int
	MaxOcena int
}

func NewOmejitve(st int, min int, max int) Omejitve {
	return Omejitve{StOcen: st, MinOcena: min, MaxOcena: max}
}

func povprecje(vpisnaStevilka string, ocene []int, slovar map[string]Student, o Omejitve) float64 {
	_, obstaja := slovar[vpisnaStevilka]
	if !obstaja {
		f.Println("Student ne obstaja")
		return float64(-1.0)
	}

	if len(ocene) < o.StOcen {
		f.Printf("Student ima manj kot %d ocen\n", o.StOcen)
		return float64(0.0)
	}

	sum := 0
	for i := range ocene {
		sum += ocene[i]
	}
	result := float64(sum) / float64(len(ocene))
	return float64(result)
}

func DodajOceno(vpisnaStevilka string, ocena int, slovar map[string]Student, o Omejitve) {
	student, obstaja := slovar[vpisnaStevilka]
	if ocena < o.MinOcena || ocena > o.MaxOcena {
		f.Printf("Ocena ni med %d in %d\n", o.MinOcena, o.MaxOcena)
	} else if !obstaja {
		f.Println("Študenta ni na seznamu")
	} else {
		student.Ocene = append(slovar[vpisnaStevilka].Ocene, ocena)
		slovar[vpisnaStevilka] = student
	}
}

// Nevem zakaj sm to tko napisu
/* func IzpisVsehOcen(vpisnaStevilka string, slovar map[string]Student, o Omejitve) {
	f.Printf("Študent z vpisno številko: %s ima zaključno oceno: %f\n", vpisnaStevilka, povprecje(vpisnaStevilka, slovar[vpisnaStevilka].Ocene, slovar, o))
} */

func IzpisVsehOcen(vpisnaStevilka string, slovar map[string]Student) {
	f.Printf("Študent ima sledeče ocene: %v\n", slovar[vpisnaStevilka].Ocene)
}

func IzpisiKoncniUspeh(vpisnaStevilka string, slovar map[string]Student, o Omejitve) {
	avg := povprecje(vpisnaStevilka, slovar[vpisnaStevilka].Ocene, slovar, o)
	f.Printf("%s %s: povprečna ocena %.1f", slovar[vpisnaStevilka].Ime, slovar[vpisnaStevilka].Priimek, avg)
	switch {
	case avg >= 9:
		f.Println(" -> Odličen študent!")
	case avg >= 6 && avg < 9:
		f.Println(" -> Povprečen študent")
	default:
		f.Println(" -> Neuspešen študent")
	}
}
