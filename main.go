package main

import "fmt"

type profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

type Profile struct {
	Name   string
	Health int
	Power  int
	Exp    int
}

func main() {
	Profiles := MakeProfile("Naruto")
	fmt.Println("Name :", Profiles.Name)
	fmt.Println("Health :", Profiles.Health)
	fmt.Println("Power :", Profiles.Power)
	fmt.Println("Name :", Profiles.Exp)
	fmt.Println("===HEROES POWER UP===")
	powerUp := PowerUP(Profiles, 2)
	fmt.Println("Name :", powerUp.Name)
	fmt.Println("Health :", powerUp.Health)
	fmt.Println("Power :", powerUp.Power)
	fmt.Println("Name :", powerUp.Exp)
}

func MakeProfile(names string) Profile {
	var name = Profile{}

	if names == "Sasuke" {
		name.Name = names
		name.Health = 200
		name.Power = 100
		name.Exp = 0
	} else if names == "Goku" {
		name.Name = names
		name.Health = 400
		name.Power = 300
		name.Exp = 100
	} else if names == "Naruto" {
		name.Name = names
		name.Health = 150
		name.Power = 200
		name.Exp = 50
	} else {
		fmt.Println("Names not found")
	}

	return name
}

func PowerUP(user Profile, multiplayer int) Profile {
	var powerUp Profile

	powerUp.Name = user.Name
	powerUp.Health = user.Health + (user.Health * multiplayer)
	powerUp.Power = user.Power + (user.Power * multiplayer)
	powerUp.Exp = user.Exp + (user.Exp * multiplayer)

	return powerUp
}
