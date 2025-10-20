{ buildGoModule, lib }:

buildGoModule {
  pname = "zing";
  version = "0.1.0";
  src = lib.cleanSource ./.;

  vendorHash = "sha256-Uy0zROdQSZb833XCjUj86iNftz12QgmyHBoQvICEpV0=";
}
