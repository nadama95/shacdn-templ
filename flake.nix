{
  description = "shadcn-templ";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";

    templ.url = "github:a-h/templ/e2511cd57e5ecd28ce6e3d944c87f1e31e20b596";
  };

  outputs = inputs@{ ... }:
    let
      system = "x86_64-linux";
      pkgs = inputs.nixpkgs.legacyPackages.${system};
      templ = system: inputs.templ.packages.${system}.templ;
    in
    {

      devShells.${system}.default = pkgs.mkShell {
        packages = with pkgs; [
          delve
          go
          gopls
          gotests
          go-tools
          (templ system)

          nodejs
        ];
      };
    };
}
