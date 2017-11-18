# golang の環境構築はUbuntuを除き自前で行ってください。(Ubuntuもパスの設定は自前でお願いします。)
# http://golang-jp.org/doc/install

"Ubuntuのlts用にgolangをインストールします."
"ltsでなくとも動くと思います。ちなみにgolangの最新版はltsでは動きません."
# sudo apt remove golang-go
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go

mkdir $HOME/go
#
# 以下、bashrc, zshrcに記述
# export GOPATH=$HOME/go
# fishならconfig.fishに
# set -gx GOPATH $HOME/go

mkdir -p ~/dev/src/github.com/pydaa
git clone https://github.com/pydaa/go-saidai-bus.git ~/dev/src/github.com/pydaa

go get github.com/Masterminds/glide
go install github.com/Masterminds/glide

cd ~/dev/src/github.com/pydaa
glide install -v

