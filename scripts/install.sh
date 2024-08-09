#!/bin/bash

read -p "Install system development libraries? (y/n) " yn
if [[ $yn == "y" ]]; then
  if [[ "$(uname)" == "Linux" ]]; then
    echo "Installing dependencies for Linux"
    sudo apt install -y libicu-dev
  elif [[ "$(uname)" == "Darwin" ]]; then
    echo "No dependencies for Mac OS"
  else
    echo "No dependencies for your OS"
  fi
fi

read -p "Install go development tools? (y/n) " yn
if [[ $yn == "y" ]]; then
  go install golang.org/x/tools/gopls@latest
  go install github.com/go-delve/delve/cmd/dlv@latest
  go install github.com/cosmtrek/air@latest
  go install github.com/a-h/templ/cmd/templ@latest
fi

