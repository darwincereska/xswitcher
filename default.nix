{
  lib,
  buildGoModule,
  fetchFromGitHub
}:

buildGoModule {
  pname = "xswitcher";
  version = "0.1.2";
  
  src = fetchFromGitHub {
    owner = "darwincereska";
    repo = "xswitcher";
    rev = "v0.1.2";
    hash = "";
  };
  
  vendorHash = null;
  
  meta = with lib; {
    description = "";
    homepage = "https://github.com/darwincereska/xswitcher";
    license = licenses.mit;
    maintainers = [ maintainers.darwincereska ];
  };
}
