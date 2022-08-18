package request

const (
	// Next release add feature to parse PGP messages
	// beginPGPMes = "-----BEGIN PGP SIGNED MESSAGE-----"
	// beginPGPSig = "-----BEGIN PGP SIGNATURE-----"
	// endPGPSig   = "-----END PGP SIGNATURE-----"

	// HTTP client fields
	securityTXT    = "/.well-known/security.txt"
	securityTXTAlt = "/security.txt"
	userAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36"
	cacheHeader    = "no-cache, private, max-age=0"
	statusOK       = 200
	statusIMUsed   = 226

	HTTPError       = "file may not exist - HTTP error code: "
	HTTPtimeoutSecs = 10
)
