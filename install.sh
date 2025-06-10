#!/bin/bash

set -e

# Configuration
FORMULA_URL="https://raw.githubusercontent.com/Ansh-Rathod/imgpim/main/Formula/imgpim.rb"
FORMULA_FILE="imgpim.rb"

# Download the formula
echo "Downloading formula from $FORMULA_URL..."
if ! curl -L "$FORMULA_URL" -o "$FORMULA_FILE"; then
    echo "Error: Failed to download formula from $FORMULA_URL."
    exit 1
fi

# Install the formula
echo "Installing imgpim formula..."
if ! brew install "./$FORMULA_FILE"; then
    echo "Error: Failed to install imgpim formula."
    rm -f "$FORMULA_FILE"
    exit 1
fi

# Clean up
rm -f "$FORMULA_FILE"
echo "Installation complete! You can now use 'imgpim'."

# Test the installation
imgpim --help
