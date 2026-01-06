{
  pkgs ? import <nixpkgs> { },
}:

pkgs.mkShell {
  name = "snippetbox-shell";
  # Tools you need to run/build (Go, GCC, tools)
  nativeBuildInputs = with pkgs; [
    air # Live reload for Go apps
    # golangci-lint # Fast linters Runner for Go
    gnumake
    go
    gcc
    gopls # Go language server
    delve # Debugger
    pkg-config # Helper to find C libraries
  ];

  # C Libraries your project depends on
  # Add any library here that gives you "missing header" errors
  buildInputs = with pkgs; [
    # Examples (uncomment if needed):
    # openssl
    # sqlite
    # xorg.libX11
    # libGL
    # vulkan-headers
  ];

  # Environment Variables
  shellHook = ''
    export CGO_ENABLED=1
    echo "❄️  Development environment loaded!"
    # Only switch to fish if it exists AND the user is King
    if [[ -z "$IN_NIX_SHELL_FISH" && "$(which fish 2>/dev/null)" != "" && "$USER" == "king" ]]; then
        export IN_NIX_SHELL_FISH=1
        export STARSHIP_CONFIG=~/.config/starship-nix.toml
    exec fish
    fi
  '';
}
