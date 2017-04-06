package BarPrinter

import (
	"fmt"
)

func New(data []uint) (*barPrinter, error) {
	return &barPrinter{
		data: data,
	}, nil
}

type barPrinter struct {
	data []uint
}

func (bp *barPrinter) Len() int {
	return len(bp.data)
}

func (bp *barPrinter) FindMax() (uint, bool) {
	var max uint = 0
	var found bool = false

	for _, value := range (*bp).data {
		if max < value {
			max = value
			found = true
		}
	}

	return max, found
}

func (bp *barPrinter) PrintBars(showScale bool) {
	var listLength int = bp.Len()

	fmt.Println()
	if max, found := bp.FindMax(); found {
		for line := max; line >= 1; line-- {
			if showScale {
				fmt.Printf("%d ", line)
			}

			bp.printLine(line, listLength)
		}

		if showScale {
			bp.printScale(listLength)
		}
	}

	fmt.Println()
}

func (bp *barPrinter) printScale(listLength int) {
	fmt.Print("  ")
	for index := 0; index < listLength; index++ {
		fmt.Printf("%d ", (*bp).data[index])
	}
	fmt.Println()
}

func (bp *barPrinter) printLine(line uint, listLength int) {
	for index := 0; index < listLength; index++ {
		if line <= (*bp).data[index] {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		fmt.Print(" ")
	}
	fmt.Println()
}
