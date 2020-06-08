---
title: Konnectd
geekdocRepo: https://github.com/owncloud/ocis-konnectd
geekdocEditPath: edit/master/docs
geekdocFilePath: _index.md
---

# Abstract

OCIS needs to authenticate users. Without an identity, there is no ownership of files. `ocis-konnectd` wraps kopano konnectd to embed an OpenID Connect Identity Provider in the ocis single binary. It provides OpenID Connect based authentication out of the box or can be replaced with an existing IdP.
