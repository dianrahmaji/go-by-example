{
  description = "Go By Example project's flake";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
	utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
	utils.lib.eachDefaultSystem (system:
	  let
	    pkgs = import nixpkgs { inherit system; };
	  in {
		devShells.default = pkgs.mkShell {
		  buildInputs = with pkgs; [
			go
		  ];
		  shellHook = ''
			zsh
		  '';
		};
	  }
	);
}
