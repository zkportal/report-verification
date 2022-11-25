# SGX Report breakdown

SGX report is used to verify hat the signing authority is rooted to a trusted authority such as the enclave platform manufacturer. Practically, it means that you can verify that the code is running in a Trusted Execution Environment. It also includes a number of attributes you, as a consumer of the report, need to verify yourself.

## `Data`
The report data that has been included in the report. This can any data up to 64 bytes that the enclave signs and attests to.

### How do I verify it?

Check the developer's documentation for details about the report data in each specific case.

## `SecurityVersion`
Enclaves that represent different versions of a module can have different security version numbers.
The SGX design disallows the migration of secrets from an enclave with a higher `SecurityVersion` to an enclave with a lower `SecurityVersion`. This restriction is intended to assist with the distribution of security patches, as follows.

If a security vulnerability is discovered in an enclave, the developer can release a fixed version with a higher `SecurityVersion`.
As users upgrade, SGX will facilitate the migration of secrets from the vulnerable version of the enclave to the
fixed version.

### How do I verify it?

Check the developer's documentation and release pages about important security updates.

## `Debug`
If true, the report is for a debug enclave. From a practical standpoint, this means that secrets will never be migrated between enclaves that support debugging and production enclaves.

### How do I verify it?

Check the developer's documentation if they're running their enclaves in debug mode. Generally, you want to check that debug mode is **disabled** for production systems.

## `UniqueID`
UniqueID uniquely identifies enclave. It changes if the program changes.

### How do I verify it?
Check the developer's documentation and release pages about the version of the software they're running. You can also verify that there were no changes to the source code compared to the published source code thanks to [Reproducible builds](./REPRODUCIBLE_BUILDS.md).

## `SignerID`
SignerID uniquely identifies enclave's signer. A developer generates a pair of RSA keys, which they use to sign enclaves.

### How do I verify it?

Check the developer's documentation to find the signer ID for the software they're running.

## `ProductID`
ProductID uniquely identifies a Product - changed by the developer to indicate different software modules, which are a part of the same enclave. All the enclaves whose signatures have the same ProductID and are issued by the same RSA key (and therefore have the same MRENCLAVE) are assumed to represent different versions of the same software module.

### How do I verify it?

Check the developer's documentation to find the product ID for the software they're running.

## `TCBStatus`

TODO
