package magicnumbers

var Catalog = Map{
	"avi":   []byte(AviMagicNumber),
	"bmp":   []byte(BmpMagicNumber),
	"exe":   []byte(ExeMagicNumber),
	"gif":   []byte(GifMagicNumber),
	"gif89": []byte(Gif89MagicNumber),
	"Kbx":   []byte(GnuPgKbx),
	"gzip":  []byte(GzipMagicNumber),
	"jpeg":  []byte(JpegMagicNumber),
	"mkv":   []byte(MkvMagicNumber),
	"mp3":   []byte(Mp3MagicNumber),
	"png":   []byte(PngMagicNumber),
	"pdf":   []byte(PdfMagicNumber),
	"rar":   []byte(RarMagicNumber),
	"rar5":  []byte(Rar5MagicNumber),
	"tar":   []byte(TarMagicNumber),
	"wav":   []byte(WavMagicNumber),
	"zip":   []byte(ZipMagicNumber),
}
