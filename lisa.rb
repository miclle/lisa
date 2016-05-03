require "formula"

class Lisa < Formula
  homepage "https://github.com/miclle/lisa"
  url "http://o6m233frm.qnssl.com/lisa.0.0.1"
  sha256 "d84fe7e07bedb227cffff10009151d96fc944f6a1bd37cff60e8e4626a1eb1c3"

  def install
    libexec.install "lisa.0.0.1"
    bin.write_jar_script libexec/"lisa.0.0.1","lisa"
  end
  
  test do
    system bin/"lisa","--help"
  end
end