.PHONY: godoc
godoc:
	@which pkgsite > /dev/null || (echo "pkgsite required for `make godoc`" && exit 1)
	pkgsite -open
