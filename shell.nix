let
  sources = import ./nix/sources.nix;
  pkgs = import sources.nixpkgs {
    overlays = [
      (self: super: {
        buildGoApplication = super.callPackage "${sources.gomod2nix}/builder" { };
        gomod2nix = super.callPackage sources.gomod2nix { };
      })
    ];
  };
in
pkgs.mkShell {
  buildInputs = with pkgs; [
    gomod2nix
  ];
}
