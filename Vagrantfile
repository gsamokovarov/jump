Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install -y golang rpm gccgo build-essential ruby-dev

    sudo gem install fpm

    mkdir -p /home/vagrant/.go/src/github.com/gsamokovarov
    chown -hR vagrant:vagrant /home/vagrant/.go

    ln -nsf /vagrant /home/vagrant/.go/src/github.com/gsamokovarov/jump

    echo "export GOPATH=/home/vagrant/.go" >> /home/vagrant/.bashrc
  SHELL
end
