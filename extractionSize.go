package magicnumbers

// extractionSize - this is the number of bytes needed to cover up to the largest magic number in Catalog
var extractionSize int = Catalog.GetLongest()
