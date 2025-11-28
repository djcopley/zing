{ buildGoModule, lib }:

buildGoModule {
  pname = "zing";
  version = "0.1.0";
  src = lib.cleanSource ./.;

  vendorHash = "sha256-1aVzQVp2/Tb99Ai226NMTSQN1u7iL9veHkZrgrf3QUc=";
}
