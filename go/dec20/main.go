package main

import (
	"fmt"

	u "github.com/einarssons/adventofcode2021/go/utils"
)

func main() {
	lines := u.ReadLinesFromFile("data.txt")
	task(lines)
}

func task(lines []string) {
	iea, img := readData(lines)

	surrChar := "."
	for i := 0; i < 50; i++ {
		img = expand(img, surrChar)
		img = enhance(img, iea, surrChar)
		if iea[0:1] == "#" {
			if surrChar == "." {
				surrChar = "#"
			} else {
				surrChar = "."
			}
		}
		if i == 1 {
			nrLit := countLit(img)
			fmt.Printf("Task 1: nrLit = %d\n", nrLit)
		}

	}
	nrLit := countLit(img)
	fmt.Printf("Task 2: nrLit = %d\n", nrLit)

}

func readData(lines []string) (string, [][]string) {
	imgEnhAlg := ""
	inAlg := true
	var image [][]string
	for _, l := range lines {
		if inAlg {
			if l == "" {
				inAlg = false
				continue
			}
			imgEnhAlg += l
			continue
		}
		image = append(image, u.SplitToChars(l))
	}
	return imgEnhAlg, image
}

func pattern(y, x int, img [][]string, surroundingChar string) string {
	p := ""
	width := len(img[0])
	height := len(img)
	for yn := y - 1; yn <= y+1; yn++ {
		for xn := x - 1; xn <= x+1; xn++ {
			if xn < 0 || yn < 0 || xn >= width || yn >= height {
				p += surroundingChar
				continue
			}
			p += img[yn][xn]
		}
	}
	return p
}

func pattern2nr(pattern string) int {
	n := 0
	for i := 0; i < len(pattern); i++ {
		if pattern[i:i+1] == "." {
			n = n << 1
		} else {
			n = n<<1 + 1
		}
	}
	return n
}

func expand(img [][]string, expChar string) [][]string {
	width := len(img[0])
	height := len(img)
	nrNewLeft := 1
	nrNewTop := 1
	nrNewRight := 1
	nrNewBottom := 1
	newWidth := width + nrNewLeft + nrNewRight
	newHeight := height + nrNewTop + nrNewBottom
	newImg := make([][]string, 0, newHeight)
	for i := 0; i < nrNewTop; i++ {
		newImg = append(newImg, newEmptyLine(newWidth, expChar))
	}
	for _, l := range img {
		nl := make([]string, newWidth)
		copy(nl, newEmptyLine(nrNewLeft, expChar))
		copy(nl[nrNewLeft:], l)
		copy(nl[nrNewLeft+width:], newEmptyLine(nrNewRight, expChar))
		newImg = append(newImg, nl)
	}
	for i := 0; i < nrNewBottom; i++ {
		newImg = append(newImg, newEmptyLine(newWidth, expChar))
	}
	return newImg
}

func printImg(img [][]string) {
	width := len(img[0])
	height := len(img)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			fmt.Printf("%s", img[y][x])
		}
		fmt.Printf("\n")
	}
}

func newEmptyLine(w int, newChar string) []string {
	l := make([]string, w)
	for i := 0; i < w; i++ {
		l[i] = newChar
	}
	return l
}

func enhance(img [][]string, iea string, surroundingChar string) [][]string {
	width := len(img[0])
	height := len(img)
	//fmt.Printf("Enhancing %dx%d\n", width, height)
	newImg := make([][]string, 0, height)
	for i := 0; i < height; i++ {
		newImg = append(newImg, make([]string, width))
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			newImg[y][x] = "."
		}
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := pattern(y, x, img, surroundingChar)
			newNr := pattern2nr(p)
			newPixel := iea[newNr : newNr+1]
			newImg[y][x] = newPixel
		}
	}
	return newImg
}

func countLit(img [][]string) int {
	count := 0
	width := len(img[0])
	height := len(img)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if img[y][x] == "#" {
				count++
			}
		}
	}
	return count
}
