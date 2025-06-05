# Maintainer: Darwin Cereska <darwin@darwincereska.dev>
pkgname=xswitcher
pkgver=0.1.0
pkgrel=1
pkgdesc=""
arch=('x86_64')
url="https://github.com/darwincereska/xswitcher"
license=('MIT')
makedepends=('go' 'git' 'make')
source=("$pkgname-$pkgver.tar.gz::https://github.com/darwincereska/xswitcher/archive/v$pkgver.tar.gz")
sha256sums=('SKIP')

prepare() {
    cd "$pkgname-$pkgver"
    
    mkdir -p bin
        
    # If you need to modify Makefile variables for packaging
    sed -i 's|/usr/local/bin|$(DESTDIR)/usr/bin|' Makefile
}

build() {
    # Use Makefile but with Arch-friendly flags
    export CGO_CPPFLAGS="${CPPFLAGS}"
    export CGO_CFLAGS="${CFLAGS}"
    export CGO_CXXFLAGS="${CXXFLAGS}"
    export CGO_LDFLAGS="${LDFLAGS}"
    export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
    
    make build
}

check() {
    cd "$pkgname-$pkgver"
    make test
}

package() {
    cd "$pkgname-$pkgver"
    install -d "$pkgdir/usr/bin"
    make DESTDIR="$pkgdir" install
}