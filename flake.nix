{
  description = "A flake for the cave CLI tool";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-24.05";
  };

  outputs = { nixpkgs, ... }:
    let
      system = "x86_64-linux";
      pkgs = nixpkgs.legacyPackages.${system};
    in
    {
      packages.${system}.cave = pkgs.stdenv.mkDerivation {
        pname = "cave";
        version = "0.0.1";
        src = ./.;
        buildInputs = [ pkgs.go ];
        
        buildPhase = ''
          export GOCACHE=$(mktemp -d /tmp/go-build-cache-XXXXXX)
          go build -o cave ./main.go
        '';

        installPhase = ''
          mkdir -p $out/bin
          cp cave $out/bin/
        '';

        meta = with pkgs.lib; {
          description = "A CLI tool to manage Nix operations";
          license = licenses.mit;
          maintainers = with maintainers; [ Matteo Cavestri ];
        };
      };
    };
}

