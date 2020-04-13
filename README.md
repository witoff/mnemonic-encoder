Mnemonic Encoder
================

**!!! WORK IN PROGRESS, THIS DOES NOT WORK SO DO NOT USE IT!!!**

Given a BIP-39 Mnemonic, convert into a PKCS#8 formatted file suitable for loading into an HSM.

**Usage**

```shell
mnemonic-encoder <mnemonic-file>    <derivation-path> <format>
mnemonic-encoder ./test/mnemonic-1   m/44'/0'/0'/0     pkcs8

> Mnemonic:  earth couch laugh second health submit annual shove eagle asset duck expire weather control truly decline uncover crash birth trip omit lumber emotion diagram
> -----BEGIN PRIVATE KEY-----
> MFQCAQEEINmVh+4jGrglZuetonhi9ZW2MOVGTLgvhAKj3ZHQ/qWnoAcGBSuBBAAK
> oSQDIgcCjuHyrTVlbADxbLOtYkYL3WyZo+5qvyazG7ij0GmKyZM=
> -----END PRIVATE KEY-----
```

**TODO**

* Fix derivation path
* Pack files into the correct pkcs8 format
* Verify output pkcs8 file can be reloaded
* Test cases
* Pretty much everything else...