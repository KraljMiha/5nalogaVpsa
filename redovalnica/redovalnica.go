// Package redovalnica provides code for operating with student's homework marks.
//
// It contains methods, which can be used in formulating a markbook for students.
//
// Example usage from Naloga1.go:
//
// slovar["63210001"] = redovalnica.Student{Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8, 7}}
// redovalnica.DodajOceno("63210001", 9, slovar, omejitve)
//
//	for vpisnaStevilka, _ := range slovar {
//			redovalnica.IzpisVsehOcen(vpisnaStevilka, slovar)
//			redovalnica.IzpisiKoncniUspeh(vpisnaStevilka, slovar, omejitve)
//			f.Println()
//		}
//
// Package redovalnica does the following operations:
// 		- Adds a mark to a student
// 		- Prints all grades of a student
// 		- Prints a student's final grade

package redovalnica

import (
	f "fmt"
)

// Creating a student
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

// Omejitve is created for checking validity of added grades
type Omejitve struct {
	StOcen   int
	MinOcena int
	MaxOcena int
}

// Creating Omejitve
func NewOmejitve(st int, min int, max int) Omejitve {
	return Omejitve{StOcen: st, MinOcena: min, MaxOcena: max}
}

// Povprecje calculates averages of a student and prints errors of possible invalid operations
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

// DodajOceno attempts to add a grade to a certain student, depending of the grade's validity
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

// IzpisVsehOcen prints all of the grades of one student
func IzpisVsehOcen(vpisnaStevilka string, slovar map[string]Student) {
	f.Printf("Študent ima sledeče ocene: %v\n", slovar[vpisnaStevilka].Ocene)
}

// IzpisiKoncniUspeh prints out the final grade
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
