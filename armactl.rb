# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Armactl < Formula
  desc "CLI for managing ARMA3 dedicated servers."
  homepage "https://github.com/brittonhayes/armactl"
  version "0.3.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/brittonhayes/armactl/releases/download/v0.3.0/armactl_0.3.0_darwin_arm64.tar.gz"
      sha256 "a7f0e0b0d22c81615639b749bd0a7961df9967aa1cd0533ec01449219c95b721"

      def install
        bin.install "armactl"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/brittonhayes/armactl/releases/download/v0.3.0/armactl_0.3.0_darwin_amd64.tar.gz"
      sha256 "8199c4a7ed63c98f7380d7d2245bcbae966386ebe51b70928c2b89e71e6f2e1c"

      def install
        bin.install "armactl"
      end
    end
  end

  on_linux do
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/brittonhayes/armactl/releases/download/v0.3.0/armactl_0.3.0_linux_arm64.tar.gz"
      sha256 "10e60118930d343b8f5651d4a01dc87029e252bc786684fee8ace57a2409e15d"

      def install
        bin.install "armactl"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/brittonhayes/armactl/releases/download/v0.3.0/armactl_0.3.0_linux_amd64.tar.gz"
      sha256 "a73d777bd1420e4a66b521c6abb2ef8c1573c60d28070efc1301f95d6de8d099"

      def install
        bin.install "armactl"
      end
    end
  end
end
