{
  buildGoModule,
  ...
}: 

let
  name = "discord-restore";
  src = ./.;
  pname = "dcrestore";
in
buildGoModule {
  inherit pname name src;

  vendorHash = "sha256-XRSx2Uie9e92REsuvTZOzTtBMcuXCaqd9DeQS7jgfwo=";
}
