generate qr codes from URL


The {{cmd .Name}} command generates QR codes from a given URL. Expects a URL string as the first argument. Optionally accepts an outut filepath string as the second argument. Optionally accepts an image size string as the third argument.

By default, the qr code will be generated at 256x256 pixels	and will be output into the CWD as 'qr.png'.

When a full path is passed as the second argument, the qr code will be output to the full path specified.

Note: All arguments are expected to be strings.
