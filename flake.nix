{
  description = "Go devshell";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
        commonTools = with pkgs; [
          git gh jq yq
          ripgrep fd fzf just
        ];
        goTools = with pkgs; [
          go
          gopls
          golangci-lint
          delve
          pkg-config
        ];
        # Add native libraries needed by cgo packages here.
        # Examples: openssl zlib sqlite libgit2 postgresql
        cgoLibraries = with pkgs; [
          alsa-lib
          libglvnd
          libxkbcommon
          wayland
          libx11
          libxcursor
          libxext
          libxi
          libxinerama
          libxrandr
          libxxf86vm
        ];
      in {
        devShells.default = pkgs.mkShell {
          packages = commonTools ++ goTools;
          buildInputs = cgoLibraries;

          shellHook = ''
            export EDITOR=nvim
            export CGO_ENABLED=1
            export GOPATH="$PWD/.go:$PWD"
            export GOBIN="$GOPATH/bin"
            export PATH="$GOBIN:$PATH"
            if [ -d /usr/lib ]; then
              export LD_LIBRARY_PATH="/usr/lib:${pkgs.lib.makeLibraryPath cgoLibraries}:$LD_LIBRARY_PATH"
            else
              export LD_LIBRARY_PATH="${pkgs.lib.makeLibraryPath cgoLibraries}:$LD_LIBRARY_PATH"
            fi
            mkdir -p "$GOBIN"
          '';
        };
      });
}
