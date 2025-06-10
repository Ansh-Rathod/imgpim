class Imgpim < Formula
  desc "CLI tool for lossless image compression using open-source libraries"
  homepage "https://github.com/Ansh-Rathod/imgpim"
  url "https://github.com/Ansh-Rathod/imgpim/archive/refs/tags/v0.1.4.tar.gz"
  sha256 "140e61ece574bd31c1d15eb09dc1c29cc4ac94bc22311b8de8f21d8a04f55d49"
  license "MIT"

  depends_on "go" => :build
  depends_on "oxipng"
  depends_on "gifsicle"
  depends_on "mozjpeg"
  depends_on "libheif"

  def install
    system "go", "build", "-o", bin/"imgpim", "."
  end

  test do
    system "#{bin}/imgpim", "--version"
  end
end