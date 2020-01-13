GITHUB_REPO="simplecatsoftware/terraform-provider-remotefile"
DESTINATION_DIR="$HOME/.terraform.d/plugins/"
DESTINATION_FILE="terraform-provider-remotefile"
OS=$(uname | tr '[:upper:]' '[:lower:]')


[[ $OS == "darwin" ]] && JQ_INSTALL_HELP="brew install jq"
[[ $OS == "linux" ]] && JQ_INSTALL_HELP="brew install jq"

hash jq 2>/dev/null || { echo >&2 "This install script requires jq to work. ($JQ_INSTALL_HELP) Aborting."; exit 1; }

echo "Finding latest release binaries from $GITHUB_REPO for $OS"
ASSETS_URL=$(curl -s "https://api.github.com/repos/$GITHUB_REPO/releases/latest" | grep "assets_url" | cut -d : -f 2,3 | tr -d \" | tr -d \,)

ASSET_URL=$(curl -L -s $ASSETS_URL | jq -r .[].browser_download_url | grep $OS)

echo "Found a binary at $ASSET_URL"
mkdir -p "$DESTINATION_DIR"
curl -o "$DESTINATION_DIR/$DESTINATION_FILE" $ASSET_URL

echo "Finished installing $GITHUB_REPO to $DESTINATION_DIR"
