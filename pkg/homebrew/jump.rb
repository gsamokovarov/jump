require "language/go"

class Jump < Formula
  desc "Like autojump or z, but fuzzy."
  homepage "https://github.com/gsamokovarov/jump"
  url "https://github.com/gsamokovarov/jump/archive/0.1.0.tar.gz"
  sha256 "584baf12a1073c349c70e854a43b19592c170cd7"
  head "https://github.com/gsamokovarov/jump.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    mkdir_p buildpath/"src/github.com/gsamokovarov"
    ln_s buildpath, buildpath/"src/github.com/gsamokovarov/jump"
    Language::Go.stage_deps resources, buildpath/"src"

    system "go", "build"
    bin.install "jump"
  end

  def caveats; <<-EOS.undent
    To install shell integration, follow the instructions at:
      #{prefix}/jump shell
    EOS
  end

  test do
    mkdir_p testpath/"test_dir"
    ENV["JUMP_HOME"] = testpath.to_s
    system "jump chdir #{testpath/"test_dir"}"

    assert_equal testpath/"test_dir", shell_output("jump cd tdir").chomp
  end
end
