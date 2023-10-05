#!/usr/bin/env bash

# Remove the aicommit alias
sed -i '/alias gitbird="git-bird"/d' ~/.bashrc

# Remove the go package
sudo rm /usr/local/bin/gitbird

echo "git-bird has been uninstalled."