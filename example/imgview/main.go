// Copyright 2011 The XGB Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// imgview is an image viewer.
package main

import (
	"fmt"
	"image"
	"image/draw"
	"os"

	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s imagefile\n", os.Args[0])
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	m, _, err := image.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	w, err := x11.NewWindow()
	if err != nil {
		fmt.Println(err)
		return
	}
	draw.Draw(w.Screen(), w.Screen().Bounds(), m, image.ZP, draw.Src)
	w.FlushImage()
	for e := range w.EventChan() {
		switch e := e.(type) {
		case ui.KeyEvent:
			if e.Key == ' ' {
				return
			}
		}
	}
}
