{ buildGoModule, lib }:

buildGoModule {
  pname = "zing";
  version = "0.1.0";
  src = lib.cleanSource ./.;

  vendorHash = "sha256-cvNRzsZXyVwbXT1Mp87XSpTTiFFzOJ1EZFWs93R8jCo=";
}
