#!/usr/bin/env bash

set -e

echo "Removing Halcyon..."
sudo rm -f /usr/local/bin/halcyon
rm -rf ~/.halcyon

echo "Halcyon has been successfully uninstalled."
