package main

import "fmt"

type IClothes interface {
	setName(name string)
	setSize(size int)
	GetName() string
	GetSize() int
}

type clothes struct {
	name string
	size int
}

func (c *clothes) setName(name string) {
	c.name = name
}

func (c *clothes) GetName() string {
	return c.name
}

func (c *clothes) setSize(size int) {
	c.size = size
}

func (c *clothes) GetSize() int {
	return c.size
}

type ANTA struct {
	clothes
}

func newANTA() IClothes {
	return &ANTA{
		clothes: clothes{
			name: "ANTA clothes",
			size: 4,
		},
	}
}

type PEAK struct {
	clothes
}

func newPEAK() IClothes {
	return &PEAK{
		clothes: clothes{
			name: "PEAK clothes",
			size: 1,
		},
	}
}

func MakeClothes(clothesType string) (IClothes, error) {
	if clothesType == "ANTA" {
		return newANTA(), nil
	}
	if clothesType == "PEAK" {
		return newPEAK(), nil
	}
	return nil, fmt.Errorf("Wrong clothes type passed")
}

func printDetails(c IClothes) {
	fmt.Printf("Clothes: %s\n", c.GetName())
	fmt.Printf("Size: %d\n", c.GetSize())
}

func main() {
	ANTA, _ := MakeClothes("ANTA")
	printDetails(ANTA)

	fmt.Println()

	PEAK, _ := MakeClothes("PEAK")
	printDetails(PEAK)
}
