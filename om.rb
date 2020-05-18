# This file was generated by GoReleaser. DO NOT EDIT.
class Om < Formula
  desc ""
  homepage ""
  version "4.8.0"
  bottle :unneeded

  if OS.mac?
    url "https://github.com/pivotal-cf/om/releases/download/4.8.0/om-darwin-4.8.0.tar.gz"
    sha256 "7502ace9936bf7c68cda93a96054d70eab743529612929857e44e1c5e454b6d8"
  elsif OS.linux?
    if Hardware::CPU.intel?
      url "https://github.com/pivotal-cf/om/releases/download/4.8.0/om-linux-4.8.0.tar.gz"
      sha256 "3fc066bddcb0f4d306bb9f51a3d397e6789b72c09875b5ba4dc6ab696b74d543"
    end
  end

  def install
    bin.install "om"
  end

  test do
    system "#{bin}/om --version"
  end
end
