{ buildGoModule, lib }:

buildGoModule {
  pname = "zing";
  version = "0.1.0";
  src = lib.cleanSource ./.;

  vendorHash = "sha256-EK/BH6fIZ3pLhKbrHkBCVoWc/ilmhAq6x1HXlr3nla0=";
}
