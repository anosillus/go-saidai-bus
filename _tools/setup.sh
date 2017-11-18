# golang の環境構築はUbuntuを除き自前で行ってください。(Ubuntuもパスの設定は自前でお願いします。)
# http://golang-jp.org/doc/install

if [ "$(uname)" == 'Darwin' ]; then
  echo "リンゴは知らん"
elif [ "$(expr substr $(uname -s) 1 5)" == 'Linux' ]; then
  "Ubuntuのlts用にgolangをインストールします."
  "ltsでなくとも動くと思います。ちなみにgolangの最新版はltsでは動きません."
  sudo apt remove golang-go
  sudo add-apt-repository ppa:longsleep/golang-backports
  sudo apt-get update
  sudo apt-get install golang-go
elif [ "$(expr substr $(uname -s) 1 10)" == 'MINGW32_NT' ]; then                                                                                           
  
  echo "Binaryをダウンロードしろ"
else
  echo "Your platform ($(uname -a)) is not supported."
  echo "勝手にやって"
  exit 1
fi


mkdir $HOME/go
#
# 以下、bashrc, zshrcに記述
# export GOPATH=$HOME/go
# fishならconfig.fishに
# set -gx GOPATH $HOME/go

mkdir -p ~/dev/src/github.com/pydaa
git clone https://github.com/pydaa/go-saidai-bus.git ~/dev/src/github.com/pydaa
