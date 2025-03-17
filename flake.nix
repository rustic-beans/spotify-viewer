{
  description = "spotify-viewer development environment";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    hl.url = "github:pamburus/hl";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, hl, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        goVersion = "1_23"; # Change this to update the whole stack

        pkgs = nixpkgs.legacyPackages.${system};
      in {
        overlays.default = final: prev: {
          go = final."go_${goVersion}";
        };

        devShells = {
          default = pkgs.mkShell {
            packages = with pkgs; [
              go
              gotools
              golangci-lint
              sqlc
              nodejs_22
              hl.packages.${system}.default
            ];
          };
        };
      }
    );
}

