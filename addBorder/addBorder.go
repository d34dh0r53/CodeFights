package addBorder

func addBorder(picture []string) []string {
	var borderedPicture []string
	var borderStrings string
	pictureWidth := len(picture[0])
	pictureHeight := len(picture)

	//Create border strings
	for x := 0; x < pictureWidth+2; x++ {
		borderStrings = borderStrings + "*"
	}

	//Add top border
	borderedPicture = append(borderedPicture, borderStrings)

	//Add side borders
	for y := 0; y < pictureHeight; y++ {
		picture[y] = "*" + picture[y] + "*"
		borderedPicture = append(borderedPicture, picture[y])
	}

	//Add bottom border
	borderedPicture = append(borderedPicture, borderStrings)

	return borderedPicture
}
