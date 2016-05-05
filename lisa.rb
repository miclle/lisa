# Documentation: https://github.com/Homebrew/brew/blob/master/share/doc/homebrew/Formula-Cookbook.md
#                http://www.rubydoc.info/github/Homebrew/brew/master/Formula
# PLEASE REMOVE ALL GENERATED COMMENTS BEFORE SUBMITTING YOUR PULL REQUEST!

class Lisa < Formula
  desc "Starting a file system watcher then execute a command"
  homepage "https://github.com/miclle/lisa"
  url "http://o6m233frm.qnssl.com/lisa/v0.0.1.tar.gz"
  sha256 "c9cdeecace9f6d8b48043c776ed3b2ec129e7b328648581fe0a7d9f7825fca9e"

  # depends_on "cmake" => :build
  # depends_on :x11 # if your formula requires any X11/XQuartz components

  def install
    # ENV.deparallelize  # if your formula fails when building in parallel
    libexec.install "lisa"

    chmod 0755, "#{libexec}/lisa"

    bin.install_symlink libexec/"lisa"
  end

  test do
    # `test do` will create, run in and delete a temporary directory.
    #
    # This test will fail and we won't accept that! It's enough to just replace
    # "false" with the main program this formula installs, but it'd be nice if you
    # were more thorough. Run the test with `brew test lisa`. Options passed
    # to `brew install` such as `--HEAD` also need to be provided to `brew test`.
    #
    # The installed folder is not in the path, so use the entire path to any
    # executables being tested: `system "#{bin}/program", "do", "something"`.
    # system "false"
    system bin/"lisa","--help"
  end
end
