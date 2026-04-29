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

  vendorHash = "";
}
