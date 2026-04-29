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

  vendorHash = "sha256-8AJ/Big0vlc2iWlkY6k3K0xnP5Nw+9VjzQYcIYsmGTg=";
}
