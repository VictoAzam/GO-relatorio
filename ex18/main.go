package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"os"
)

func main() {
	var largura, altura int
	fmt.Print("Largura e altura (ex: 200 100): ")
	fmt.Scan(&largura, &altura)
	if largura <= 0 || altura <= 0 {
		fmt.Println("Dimensões inválidas")
		return
	}
	R := make([][]uint8, altura)
	G := make([][]uint8, altura)
	B := make([][]uint8, altura)
	for y := 0; y < altura; y++ {
		R[y] = make([]uint8, largura)
		G[y] = make([]uint8, largura)
		B[y] = make([]uint8, largura)
		for x := 0; x < largura; x++ {
			R[y][x] = uint8((x * 255) / (largura - 1))
			G[y][x] = uint8((y * 255) / (altura - 1))
			B[y][x] = uint8(((x + y) * 255) / (largura + altura - 2))
		}
	}
	img := image.NewRGBA(image.Rect(0, 0, largura, altura))
	for y := 0; y < altura; y++ {
		for x := 0; x < largura; x++ {
			img.Set(x, y, color.RGBA{R[y][x], G[y][x], B[y][x], 255})
		}
	}
	f, err := os.Create("imagem_rgb.jpg")
	if err != nil {
		fmt.Println("Erro criando arquivo:", err)
		return
	}
	defer f.Close()
	if err := jpeg.Encode(f, img, &jpeg.Options{Quality: 90}); err != nil {
		fmt.Println("Erro ao salvar JPEG:", err)
		return
	}
	fmt.Println("Arquivo salvo: imagem_rgb.jpg")
}
