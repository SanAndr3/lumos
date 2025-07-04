{ pkgs ? import ../../../nix { } }:
let lumosd = (pkgs.callPackage ../../../. { });
in
lumosd.overrideAttrs (oldAttrs: {
  patches = oldAttrs.patches or [ ] ++ [
    ./broken-lumosd.patch
  ];
})
