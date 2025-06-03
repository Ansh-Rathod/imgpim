class imgpim < Formula
  desc "CLI tool for compressing images using open-source libraries"
  homepage "https://github.com/Ansh-Rathod/image-compressor"
  url "https://github.com/Ansh-Rathod/image-compressor/archive/refs/tags/v0.1.0.tar.gz"
  sha256 "replace_with_actual_sha256"
  license "MIT"

  depends_on "go" => :build
  depends_on "oxipng"
  depends_on "gifsicle"
  depends_on "mozjpeg"
  depends_on "jpegoptim"
  depends_on "libheif"

  def install
    system "go", "build", "-o", bin/"image-compressor", "."
  end

  test do
    system "#{bin}/image-compressor", "--version"
  end
end