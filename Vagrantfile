# frozen_string_literal: true

Vagrant.configure(2) do |config|
  config.vm.box = 'hashicorp/bionic64'

  config.vm.provision 'shell', inline: <<-SHELL
    sudo apt-add-repository -y ppa:brightbox/ruby-ng
    sudo apt-get update
    sudo apt-get install -y rpm build-essential ruby3.2 ruby3.2-dev git

    sudo apt install flatpak

    sudo gem install fpm
    sudo gem install ronn

    GOVERSION=1.22.4

    cd /tmp
    wget https://storage.googleapis.com/golang/go$GOVERSION.linux-amd64.tar.gz
    tar -C /usr/local -xzf go$GOVERSION.linux-amd64.tar.gz

    mkdir -p /home/vagrant/.go/src/github.com/gsamokovarov
    chown -hR vagrant:vagrant /home/vagrant/.go

    ln -nsf /vagrant /home/vagrant/.go/src/github.com/gsamokovarov/jump

    echo "export GOPATH=/home/vagrant/.go" >> /home/vagrant/.bashrc
    echo "export PATH=/usr/local/go/bin:$PATH" >> /home/vagrant/.bashrc
  SHELL
end
