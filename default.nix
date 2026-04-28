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

  vendorHash = "sha256-dZy/c/HBxgho0OucFzQbdiDsw9NtCe0GFaZnp3LqSnE=";
}
