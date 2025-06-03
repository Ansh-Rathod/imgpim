class Imgpim < Formula
  desc "CLI tool for lossless image compression using open-source libraries"
  homepage "https://github.com/Ansh-Rathod/imgpim"
  url "https://github.com/Ansh-Rathod/imgpim/archive/refs/tags/v0.1.3.tar.gz"
  sha256 "f31c0c17d76cfbdec1b69ea33b090b74f17d9f87168d9ac0c3755723111a2bc8"
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