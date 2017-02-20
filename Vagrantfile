Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu/trusty64"

  config.vm.provision "shell", inline: <<-SHELL
    sudo apt-get update
    sudo apt-get install -y rpm build-essential ruby-dev

    sudo gem install fpm

    cd /tmp
    wget https://storage.googleapis.com/golang/go1.8.linux-amd64.tar.gz
    tar -C /usr/local -xzf go1.8.linux-amd64.tar.gz

    mkdir -p /home/vagrant/.go/src/github.com/gsamokovarov
    chown -hR vagrant:vagrant /home/vagrant/.go

    ln -nsf /vagrant /home/vagrant/.go/src/github.com/gsamokovarov/jump

    echo "export GOPATH=/home/vagrant/.go" >> /home/vagrant/.bashrc
    echo "export PATH=/usr/local/go/bin:$PATH" >> /home/vagrant/.bashrc
  SHELL
end
