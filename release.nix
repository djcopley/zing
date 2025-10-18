let
  nixpkgs = fetchTarball {
    url = "https://github.com/NixOS/nixpkgs/tarball/nixos-25.05";
    sha256 = "sha256:0dz08mm268j34aaxnfp8j776d3nzjjh94yfa5z0ab7ibna4mz1zx";
  };
  pkgs = import nixpkgs {
    overlays = [
      (import ./overlay.nix)
    ];
  };
in
with pkgs;
{
  inherit zing;
  default = zing;

  devShell = pkgs.mkShell {
    inputsFrom = [
      zing
    ];
    packages = [
      zing
      git
      jujutsu
      protobuf
      protoc-gen-go
      go
      gotools
      caddy
      redis
    ];
  };
}
