#!/bin/zsh
# Update package list
sudo apt-get update

# ----------------- install -----------------
#install zsh
echo "installing zsh and turn to default shell"
sudo apt install zsh -y
chsh -s $(which zsh)
zsh --version
echo "You need to logout"

#install oh-my-zsh
echo "installing oh-my-zsh"
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"


# ----------------- config -----------------
# clear file test
bash ./scripts/clear_file.sh ./test.txt

# install zsh plugins
echo "installing zsh plugins"
git clone https://github.com/fdellwing/zsh-bat.git $ZSH_CUSTOM/plugins/zsh-bat
git clone https://github.com/MichaelAquilina/zsh-you-should-use.git $ZSH_CUSTOM/plugins/you-should-use
git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions
git clone https://github.com/asdf-vm/asdf.git ~/.asdf

# update .zshrc file
zshrcContent="export ZSH="$HOME/.oh-my-zsh"

plugins=(git asdf zsh-autosuggestions zsh-syntax-highlighting you-should-use zsh-bat)

source $ZSH/oh-my-zsh.sh

. $HOME/.asdf/asdf.sh
. $HOME/.asdf/completions/asdf.bash
"
echo "$zshrcContent" > ./test.txt
echo "~/.zshrc updated"
