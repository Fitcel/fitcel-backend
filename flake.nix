{
  description = "com.awesomesuite.org projects";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-22.11";
    devenv.url = "github:cachix/devenv";
  };
  outputs = { self, nixpkgs, devenv, ... }@inputs:
    let pkgs = nixpkgs.legacyPackages."x86_64-linux";
    in {
      devShell.x86_64-linux = devenv.lib.mkShell {
        inherit inputs pkgs;
        modules = [
          (_: {
            languages.go.enable = true;
            services.mysql.enable = true;
          })
        ];
      };
    };
}
