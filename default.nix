{ buildGoModule
, nix-gitignore
}:

buildGoModule rec {
  pname = "cgroup-mover";
  version = "0.1.0";

  src = nix-gitignore.gitignoreSource [ ] ./.;

  # The checksum of the Go module dependencies. `vendorSha256` will change if go.mod changes.
  # If you don't know the hash, the first time, set:
  # sha256 = "0000000000000000000000000000000000000000000000000000";
  # then nix will fail the build with such an error message:
  # hash mismatch in fixed-output derivation '/nix/store/m1ga09c0z1a6n7rj8ky3s31dpgalsn0n-source':
  # wanted: sha256:0000000000000000000000000000000000000000000000000000
  # got:    sha256:173gxk0ymiw94glyjzjizp8bv8g72gwkjhacigd1an09jshdrjb4
  vendorSha256 = "0000000000000000000000000000000000000000000000000000";

  buildFlagsArray = ''
    -ldflags=
        -X cdp/version.Version=${version}
  '';
}
