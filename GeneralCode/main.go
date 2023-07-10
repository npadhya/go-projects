package main

import (
    "os"
    "fmt"
    "unsafe"
    "image"
    "image/png"
)
import "github.com/dveselov/go-libreofficekit"

func main() {
    office, _ := libreofficekit.NewOffice("/usr/bin/libreoffice")
    document, _ := office.LoadDocument("/home/nikul/Documents/Important Documents Synced/Financial documents/Robinhood Documents/June2023.ods")

    rectangles := document.GetPartPageRectangles()
    canvasWidth := libreofficekit.TwipsToPixels(rectangles[0].Dx(), 120)
    canvasHeight := libreofficekit.TwipsToPixels(rectangles[0].Dy(), 120)

    m := image.NewRGBA(image.Rect(0, 0, canvasWidth, canvasHeight))

    for i, rectangle := range rectangles {
        document.PaintTile(unsafe.Pointer(&m.Pix[0]), canvasWidth, canvasHeight, rectangle.Min.X, rectangle.Min.Y, rectangle.Dx(), rectangle.Dy())
        libreofficekit.BGRA(m.Pix)
        out, _ := os.Create(fmt.Sprintf("page_%v.png", i))
        png.Encode(out, m)
        out.Close()
    }
}