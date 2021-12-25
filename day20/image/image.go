package image

type Image struct {
	algo	[]bool
	// assumed to be a square
	pixels	[][]bool

	// denotes the state of all the other pixels "outside" the image
	outside	bool
}

// create a new image based on old image data
// to be used as a holder for an enhanced image
func NewEmptyEnhancedImage(old Image) Image {
	// the size increases by 1 to each side, so 2 horizontal, 2 vertical
	size := len(old.pixels) + 2

	img := Image{
		algo: old.algo,
		pixels: make([][]bool, size),
		outside: old.outside,
	}

	for i := 0; i < size; i++ {
		img.pixels[i] = make([]bool, size)
	}

	return img
}

func NewImage(algoStr string, pixelStr []string) Image {
	img := Image{
		algo: parseImageAlgo(algoStr),
		pixels: parseImageString(pixelStr),
	}
	return img
}

// process the pixel at i.pixels[row][col] and
// return the new pixel based on p.algo
func (i Image) processPixel(row, col int) bool {
	imgSize := len(i.pixels)
	
	// initialize binary data as outside value
	binaryData := make([]bool, 9)
	for j := range binaryData {
		binaryData[j] = i.outside
	}

	// top left
	if row-1 >= 0 && col-1 >= 0 {
		binaryData[0] = i.pixels[row-1][col-1]
	}
	// top mid
	if row-1 >= 0 && col >= 0 && col <= imgSize-1 {
		binaryData[1] = i.pixels[row-1][col]
	}
	// top right
	if row-1 >= 0 && col+1 <= imgSize-1 {
		binaryData[2] = i.pixels[row-1][col+1]
	}

	// mid left
	if row >= 0 && row <= imgSize-1 && col-1 >= 0 {
		binaryData[3] = i.pixels[row][col-1]
	}
	// mid mid
	if row >= 0 && row <= imgSize-1 && col >= 0 && col <= imgSize-1 {
		binaryData[4] = i.pixels[row][col]
	}
	// mid right
	if row >= 0 && row <= imgSize-1 && col+1 <= imgSize-1 {
		binaryData[5] = i.pixels[row][col+1]
	}

	// bot left
	if row+1 <= imgSize-1 && col-1 >= 0 {
		binaryData[6] = i.pixels[row+1][col-1]
	}
	// bot mid
	if row+1 <= imgSize-1 && col >= 0 && col <= imgSize-1 {
		binaryData[7] = i.pixels[row+1][col]
	}
	// bot right
	if row+1 <= imgSize-1 && col+1 <= imgSize-1 {
		binaryData[8] = i.pixels[row+1][col+1]
	}

	index := bin2dec(binaryData)
	return i.algo[index]
}

func (i *Image) Enhance() {
	// it has to be done simultaneusly for each pixel
	// so we make an entirely new image
	newImg := NewEmptyEnhancedImage(*i)

	// determine the new outside
	// if outside is off then its new value is the binary 000000000
	// which is the first index, 0
	if !i.outside {
		newImg.outside = i.algo[0]
	} else {
		// if outside is on then its new value is the binary 111111111
		// which is the final index, 511
		newImg.outside = i.algo[511]
	}
	
	// process each pixel and put them in the new image
	// the main idea is that we only care about the "main" image
	// and also the "border" surrounding it (border of 1 pixel)
	// since the border are the only pixels affected by our "main" image
	// the border of our border will not be affected by our "main" image
	// since it will be all dark (000000000) or all lit (111111111)
	// which we already took care of using the Image.outside variable
	imgSize := len(i.pixels)
	for j := -1; j <= imgSize; j++ {
		for k := -1; k <= imgSize; k++ {
			newImg.pixels[j+1][k+1] = i.processPixel(j, k)
		}
	}

	// overwrite the old image
	*i = newImg
}

func (i *Image) EnhanceMultiple(times int) {
	for j := 0; j < times; j++ {
		i.Enhance()
	}
}
