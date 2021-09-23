package main

import (
	"fmt"
	"time"
)

type personnage struct { //creation d'une classe
	nom        string
	classe     string
	niveau     int
	pvmax      int
	pvactuel   int
	inventaire []string
	skill      []string
}

func (p *personnage) init(nom string, classe string, pvmax int, pvactuel int, niveau int, inventaire []string, skill string) { //initialise des personnages
	p.nom = nom
	p.classe = classe
	p.pvmax = pvmax
	p.niveau = niveau
	p.pvactuel = pvactuel
	p.inventaire = inventaire
	p.skill = []string{"coup de point"}
}
func (p *personnage) displayInfo() { // affiche les attribut des personnages
	fmt.Println("nom:", p.nom)
	fmt.Println("classe:", p.classe)
	fmt.Println("viemaximun:", p.pvmax)
	fmt.Println("PV:", p.pvactuel)
	fmt.Println("INVENTAIRE:", p.inventaire)
	fmt.Println("skill:", p.skill)
}
func (p *personnage) popovie() { // soigne le perso
	if p.pvmax == p.pvactuel { //si le personnage a toutes ses vies il ne peut pas se soigner
		fmt.Print("tu a deja toutes tes vies")
	} else {
		if len(p.inventaire) == 0 { //si l'inventaire du personnage est vide il peut pas se soigner
			fmt.Print("ton inventaire est vide")
		} else {
			for i := 0; i < len(p.inventaire); i++ { // parcours de l'inventaire du personnage a la recherche de popo de soin
				if p.inventaire[i] == "popovie" {
					if p.pvactuel+50 > p.pvmax { // verifie si la santé du personnage ne sera pas superieur a celle maximun autorisé lors du heal
						p.pvactuel = p.pvactuel + p.pvmax - p.pvactuel
						p.removeInventory("popovie") //appelle de la fonction remove qui supprimme la popo de soin consommé
						fmt.Println(p.nom, ":", p.pvactuel, "/", p.pvmax)
						break
					} else {
						p.pvactuel = p.pvactuel + 50
						p.removeInventory("popovie")
						fmt.Println(p.nom, ":", p.pvactuel, "/", p.pvmax)
						break
					}
				} else {
					fmt.Println("Plus de popo")
				}
			}
		}
	}
}

func (p *personnage) accessInventory() { // permet d'affiche le contenu d'un inventaire
	for i := len(p.inventaire); i <= len(p.inventaire); i++ {
		if len(p.inventaire) != 0 {
			fmt.Println("Inventaire : \n", "--------------------------------")
			fmt.Println(p.inventaire)
		} else if len(p.inventaire) == 0 {
			fmt.Println("Inventaire : \n", "--------------------------------")
			fmt.Println("L'inventaire est vide ")
		}
	}
}

func (p *personnage) removeInventory(itemremove string) { //supprime un objet de l'inventaire d'un personnage
	for i := 0; i <= len(p.inventaire); i++ {
		if itemremove == p.inventaire[i] {
			p.inventaire[i] = ""
			break
		}
	}
}
func (p *personnage) addInventory(itemadd string) bool {
	if len(p.inventaire) < 10 {
		p.inventaire = append(p.inventaire, itemadd) // ont ajoute dans l'inventaire du personnage un nouvelle item pour l'instant inconnue
	} else {
		fmt.Println("inventaire complet")
		return false
	}
	return true
}

func (p *personnage) pnj(i int) { // pnj vendeurs qui vend pas
	if i == 0 {
		p.addInventory("popovie")
	} else if i == 1 {
		p.addInventory("poison")
	} else {
		if p.addInventory("Livre de Sort: Boule de feu") == false {
			fmt.Println("Plus de Place ☺")

		} else {
			p.removeInventory("Livre de Sort: Boule de feu")
			p.spellBook("Boule de feu")
		}
	}
}

func (p *personnage) dead() { //verifie si le perso est mort
	if p.pvactuel < 0 {
		fmt.Println(p.nom, ": a succombé(e)")
		p.pvactuel = p.pvmax * 50 / 100
	}
}
func (p *personnage) poison() { // retire 30 hp aux personnages
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(p.nom, ":", p.pvactuel)
		p.pvactuel = p.pvactuel - 10
		if p.pvactuel < 0 {
			p.dead()
			break
		}

	}
}
func (p *personnage) spellBook(talentcaché string) { //attribue des compétemces en fonctions des livres achetés
	for i := 0; i < len(p.skill); i++ {
		if p.skill[i] == talentcaché {
			fmt.Println("Tu possédes déjà ce talent")
			break
		} else {
			p.skill = append(p.skill, talentcaché)
			break
		}
	}
}
func main() {
	var p1 personnage
	p1.init("jackouille", "fripouille", 150, 10, 1, []string{"popovie", "popovie", "poison", "popovie", "popovie", "popovie", "popovie", "popovie", "popovie"}, "coup de point")
	p1.poison()
	p1.pnj(3)
	p1.displayInfo()
	fmt.Println()
}
