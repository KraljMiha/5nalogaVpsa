package main

import (
	f "fmt"
)

type Student struct {
	ime     string
	priimek string
	ocene   []int
}

func dodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int) {
	if ocena < 0 || ocena > 10 {
		f.Println("Ocena ni med 0 in 10")
		return
	}
	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		f.Println("Študenta ni na seznamu")
		return
	}

	student.ocene = append(student.ocene, ocena)
	studenti[vpisnaStevilka] = student

}

func povprecje(studenti map[string]Student, vpisnaStevilka string) float64 {

	student, obstaja := studenti[vpisnaStevilka]
	if !obstaja {
		f.Println("Student ne obstaja")
		return float64(-1.0)
	}

	if len(student.ocene) < 6 {
		f.Println("Student ima manj kot 6 ocen")
		return float64(0.0)
	}

	sum := 0
	for i := range student.ocene {
		sum += student.ocene[i]
	}
	result := float64(sum) / float64(len(student.ocene))
	return float64(result)
}

func izpisRedovalnice(studenti map[string]Student) {
	f.Println("REDOVALNICA:")
	for kljuc, student := range studenti {
		f.Printf("%s - %s %s: ", kljuc, student.ime, student.priimek)
		f.Print(student.ocene, "\n")
	}

}

func izpisiKoncniUspeh(studenti map[string]Student) {
	for vpisnaStevilka, student := range studenti {
		avg := povprecje(studenti, vpisnaStevilka)
		f.Printf("%s %s: povprečna ocena %.1f", student.ime, student.priimek, avg)
		switch {
		case avg >= 9:
			f.Println(" -> Odličen študent!")
		case avg >= 6 && avg < 9:
			f.Println(" -> Povprečen študent")
		default:
			f.Println(" -> Neuspešen študent")
		}
	}
}

func main() {
	slovar := make(map[string]Student)

	slovar["63210001"] = Student{ime: "Ana", priimek: "Novak", ocene: []int{10, 9, 8}}
	slovar["63210002"] = Student{ime: "Boris", priimek: "Kralj", ocene: []int{6, 7, 5, 8}}
	slovar["63210003"] = Student{ime: "Janez", priimek: "Novak", ocene: []int{4, 5, 3, 5}}

	// test Boris Kralj
	f.Println("Test Boris Kralj")
	povprecje(slovar, "63210002")      // premalo ocen
	povprecje(slovar, "99999999")      // ne obstaja
	dodajOceno(slovar, "63210002", 10) // dodajanje ocene
	dodajOceno(slovar, "63210002", 10) // dodajanje ocene
	povprecje(slovar, "63210002")      // povprecje po dodanih ocenah, ne izpiše nič, ker so vredu

	f.Println()
	f.Println("Izpis redovalnice pred dodajanjem")
	izpisRedovalnice(slovar)

	f.Println()
	f.Println("Izpis redovalnice po dodajanju:")
	dodajOceno(slovar, "63210001", 9) // dodajanje ocene Ana
	dodajOceno(slovar, "63210001", 9) // dodajanje ocene
	dodajOceno(slovar, "63210001", 9) // dodajanje ocene
	dodajOceno(slovar, "63210003", 6) // dodajanje ocene Janez
	dodajOceno(slovar, "63210003", 6) // dodajanje ocene

	izpisRedovalnice(slovar)

	f.Println()
	f.Println("Izpis končnega uspeha")
	izpisiKoncniUspeh(slovar)
}
