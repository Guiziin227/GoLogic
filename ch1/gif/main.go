package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Red
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
	color.RGBA{0x00, 0xFF, 0xFF, 0xFF}, // Cyan
	color.RGBA{0xFF, 0x00, 0xFF, 0xFF}, // Magenta
	color.Black,                        // Black
}

func main() {
	// Gerar uma fonte de números aleatórios com base no tempo atual
	randSource := rand.NewSource(time.Now().UnixNano())
	r := rand.New(randSource)

	// Criar um arquivo para salvar o GIF gerado
	outfile, err := os.Create("lissajous.gif")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao criar o arquivo: %v\n", err)
		return
	}
	defer outfile.Close() // Fechar o arquivo quando terminar

	// Gerar o GIF e salvar no arquivo
	lissajous(outfile, r)
}

func lissajous(out io.Writer, r *rand.Rand) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 200   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := r.Float64() * 3.0 // Frequência aleatória do oscilador y
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // Diferença de fase

	// Gerar os quadros da animação
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		// Gerar a forma Lissajous
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Codificar e salvar o GIF no arquivo
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao codificar o GIF: %v\n", err)
	}
}
