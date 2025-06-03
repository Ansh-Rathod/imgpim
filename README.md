imgpim
A command-line tool for lossless compression of images using open-source libraries like oxipng, gifsicle, mozjpeg, and libheif. Supports jpg, jpeg, png, gif, heic formats, with planned support for webp, bmp, and tiff.
Installation
Install via Homebrew:
brew install imgpim

Or install manually:

Clone the repository:
git clone https://github.com/Ansh-Rathod/imgpim.git
cd imgpim

Install dependencies:
brew install oxipng gifsicle mozjpeg libheif

Build and install:
go build -o imgpim
sudo mv imgpim /usr/local/bin/

Usage
Compress a single image:
imgpim input.jpg --output compressed.jpg

Compress all images in a directory:
imgpim /path/to/images --output /path/to/output

Supported Formats

JPEG: Losslessly compressed using mozjpeg (jpegtran)
PNG: Losslessly compressed using oxipng
GIF: Losslessly compressed using gifsicle
HEIC: Losslessly compressed using libheif (where supported)
WebP, BMP, TIFF: Planned for future support

License
MIT License
