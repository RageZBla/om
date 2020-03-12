# `om` Changelog

Nominally kept up-to-date as we work,
sometimes pushed post-release.

### Versioning

`om` went 1.0.0 on May 7, 2019

As of that release, `om` is [semantically versioned](https://semver.org/).
When consuming `om` in your CI system, 
it is now safe to pin to a particular minor version line (major.minor.patch)
without fear of breaking changes.

#### API Declaration for Semver

Any changes to the `om` commands are considered a part of the `om` API.
Any changes to `om` commands will be released according to the semver versioning scheme defined above.
The exceptions to this rule are any commands marked as "**EXPERIMENTAL**"
- "**EXPERIMENTAL**" commands work, and pull information from the API
  same as any other. The format in which the information is returned, however,
  is subject to change without announcing a breaking change
  by creating a major or minor bump of the semver version. 
  When the `om` team is comfortable enough with the command output,
  the "**EXPERIMENTAL**" mark will be removed.
  
Any changes to the `om` filename as presented in the Github Release page.
  
Changes internal to `om` will _**NOT**_ be included as a part of the om API.
The `om` team reserves the right to change any internal structs or structures
as long as the outputs and behavior of the commands remain the same.

**NOTE**: Additional documentation for om commands 
leveraged by Pivotal Platform Automation 
can be found in [Pivotal Documentation](docs.pivotal.io/platform-automation).

`om` is versioned independently from platform-automation. 

### Tips
* Use environment variables
  to set what Ops Manager `om` is targeting.
  For example:
  ```bash
  $  export OM_PASSWORD=example-password om -e env.yml deployed-products
  ```
  Note the additional space before the `export` command.
  This ensures that commands are not kept in `bash` history.
  The environment variable `OM_PASSWORD` will overwrite the password value in `env.yml`. 

## 3.2.3

### Bug Fixes
- Downloading a stemcell associated with a product will try to download the light or heavy stemcell. 
  If anyone has experienced the recent issue with download-product and the AWS heavy stemcell, 
  this will resolve your issue. Please remove any custom globbing that might've been added to 
  circumvent this issue. For example, `stemcall-iaas: light*aws` should just be `stemcell-iaas: aws` now. 
  
## 3.2.2

### Bug Fixes
- `staged-config` will now work again for Ops Manager versions <= 2.3.
  This solves issue [#419](https://github.com/pivotal-cf/om/issues/419)

## 3.2.1

### Bug Fixes
- `configure-director` now will configure VM Extensions before setting Resource Config.
  This fixes issue [#411](https://github.com/pivotal-cf/om/issues/411)   
  
## 3.2.0

### Features
- `expiring-certificates` command was added.
  This command returns a list of certificates
  from an Ops Manager
  expiring within a specified (`--expires-within/-e`) timeframe. 
  Default: "3m" (3 months)
- `configure-product` and `staged-config` now have support for the `/syslog_configurations` endpoint. 
  This affects tiles, such as the Metrics tile,
  that do not return these properties nested in the `product-properties` section. 
  This provides a solution for issue [331](https://github.com/pivotal-cf/om/issues/331).
  An example of this inside of your product config:
  
    ```yaml
    syslog-properties:
      address: example.com
      custom_rsyslog_configuration: null
      enabled: true
      forward_debug_logs: false
      permitted_peer: null
      port: "4444"
      queue_size: null
      ssl_ca_certificate: null
      tls_enabled: false
      transport_protocol: tcp
    ```
- `generate-certificate` can now accept multiple `--domains | -d` flags.
  Comma separated values can be passed with a single `--domains | -d` flag,
  or using a `--domains | -d` flag for each value. Thanks to @jghiloni for the PR!
  Example:
    ```bash
      om -e env.yml generate-certificate -d "example1.com" --domains "example2.com" \
         -d "example3.com,*.example4.com" --domains "example5.com,*.example6.com"
    ```
- `product-metadata` has been added to replace `tile-metadata`.
  This was done to increase naming consistency.
  Both commands currently exist and do exactly the same thing.
  Thank you @jghiloni for the PR!
- `config-template` now supports the `--exclude-version` flag.
  If provided, the command will exclude the version directory in the `--output-directory` tree.
  The contents will with or without the flag will remain the same.
  Please note including the `--exclude-version` flag
  will make it more difficult to track changes between versions
  unless using a version control system (such as git).
  Thanks to @jghiloni for the PR!
- `config-template` supports `--pivnet-disable-ssl` to skip SSL validation.
- When interacting with an OpsManager, that OpsManager may have a custom CA cert.
  In the global options `--ca-cert` has been added to allow the usage of that custom CA cert.
  The value of `--ca-cert` can be a file or command line string.
  
### Bug Fix
- When using `config-template` or `download-product`,
  the `--pivnet-skip-ssl` is honored when capturing the token. 

### Deprecation Notices
- `tile-metadata` has been deprecated in favor of `product-metadata`.
  This was done to increase naming consistency.
  Both commands currently exist and do exactly the same thing.
  The `tile-metadata` command will be removed in a future release.
  
## 3.1.0

### Features

- TLS v1.2 is the minimum version supported when connecting to an Ops Manager
- `config-template` now will provide required-vars in addition to default-vars.
- `config-template` will define vars with an `_` instead of a `/`.
  This is an aesthetically motivated change.
  Ops files are denoted with `/`,
  so changing the vars separators to `_` makes this easier to differentiate.
- `config-template` output `product-default-vars.yml` has been changed to `default-vars.yml`
- `staged-config` includes the property `max_in_flight` will be included
  in the `resource-config` section of a job.
- `configure-product` can set the property `max_in_flight`
  in the `resource-config` section of a job.

  The legal values are:
  * an integer for the number of VMs (ie `2`)
  * a percentage of 1-100 (ie `20%`)
  * the default value specified in tile (`default`)
  For example,

  ```yaml
  resource-config:
    diego_cells:
      instances: 10
      max_in_flight: 10
  ```

## 3.0.0

### Features

- `pivnet-api-token` is now optional in `download-product`
  if a source is defined. Thanks to @vchrisb for the PR!
- `configure-authentication`, `configure-ldap-authentication`, and `configure-saml-authentication`
  can create a UAA client on the Ops Manager vm.
  The client_secret will be the value provided to this option `precreated-client-secret`.
- add support for NSX and NSXT in Ops Manager 2.7+
  
### Breaking Changes

- remove `--skip-unchanged-products` from `apply-changes`
  This option has had issues with consistent successful behaviour.
  For example, if the apply changes fails for any reason, the subsequent apply changes cannot pick where it left off.
  This usually happens in the case of errands that are used for services.
  
  We are working on scoping a selective deploy feature that makes sense for users.
  We would love to have feedback from users about this.
  
- remove `revert-staged-changes`
  `unstage-product` functionally does the same thing,
  but uses the API. This resolves issue [#399](https://github.com/pivotal-cf/om/issues/399)
  
### Bug Fix
- `apply-changes` will error with _product not found_ if that product has not been staged.
- `upload-stemcell` now accepts `--floating false` in addition to `floating=false`.
  This was done to offer consistency between all of the flags on the command.
- `configure-director` had a bug in which `iaas_configurations` could not be set
  on AWS/GCP/Azure because "POST" was unsupported for these IAASes
  (Multiple IAAS Configurations only work for vSphere and Openstack).
  `configure-director` will now check if the endpoint is supported.
  If it is not supported, it will construct a payload, and selectively configure
  iaas_configuration as if it were nested under `properties-configuration`. 
  _The behavior of this command remains the same._ 
  IAAS Configuration may still be set via `iaas_configurations` OR `properties.iaas_configuration`  
  

## 2.0.1

Was a release to make sure that `brew upgrade` works.

## 2.0.0

### Features
- `configure-ldap-authentication` and `configure-saml-authentication` can create a UAA client on the Ops Manager vm.
  The client_secret will be the value provided to this option `precreated-client-secret`.
  This is supported in OpsManager 2.5+.
- A homebrew formula has been added!
  It should support both linux and mac brew.
  Since, we don't have our own `tap`, we've used a simpler method:
  ```bash
  brew tap pivotal-cf/om https://github.com/pivotal-cf/om
  brew install om
  ```
  
### Bug Fixes
- The order of vm types and resources was being applied in the correct order.
  Now vm types will be applied then resources, so that resource can use the vm type.
- When using `bosh-env`, a check is done to ensure the SSH private key exists.
  If does not the command will exit 1.
- `config-template` will enforce the default value for a property to always be `configurable: false`.
  This is inline with the OpsManager behaviour.
  
### Breaking Change
- The artifacts on the Github Release include `.tar.gz` (for mac and linux) and `.zip` (windows) for compression.
  It also allows support for using `goreleaser` (in CI) to create other package manager artifacts -- `brew`.
  This will break globs that were permissive. For example `*linux*`, will download the binary and the `.tar.gz`, use `*linux*[^.gz]` to just download the binary.
  Our semver API declaration has been updated to reflect this.

## 1.2.0

### Features 
* Both `om configure-ldap-authentication` 
  and `om configure-saml-authentication`
  will now automatically
  create a BOSH UAA admin client as documented [here](https://docs.pivotal.io/pivotalcf/2-5/customizing/opsmanager-create-bosh-client.html#saml).
  This is only supported in OpsManager 2.4 and greater.
  You may specify the flag `skip-create-bosh-admin-client`
  to skip creating this client.
  If the command is run for an OpsManager less than 2.4,
  the client will not be created and a warning will be printed.
  However, it is recommended that you create this client.
  For example, your SAML or LDAP may become unavailable,
  you may need to sideload patches to the BOSH director, etc.
  Further, in order to perform automated operations on the BOSH director,
  you will need this BOSH UAA client.
  After the client has been created,
  you can find the client ID and secret
  by following [steps three and four found here](https://docs.pivotal.io/pivotalcf/2-5/customizing/opsmanager-create-bosh-client.html#-provision-admin-client).
* `om interpolate` now allows for the `-v` flag
  to allow variables to be passed via command line. 
  Command line args > file args > env vars.
  If a user passes a var multiple times via command line,
  the right-most version of that var will
  be the one that takes priority,
  and will be interpolated.
* `om configure-director` now supports custom VM types.
  Thanks for the PR @jghiloni!
  Refer to the [VM Types Bosh documentation](https://bosh.io/docs/cloud-config/#vm-types) for IaaS specific use cases.
  For further info: [`configure-director` readme](https://github.com/pivotal-cf/om/tree/master/docs/configure-director#vmtypes-configuration). 
  Please note this is an advanced feature, and should be used at your own discretion.
* `download-product` will now return a `download-file.json` 
  if `stemcell-iaas` is defined but the product has no stemcell.
  Previously, this would exit gracefully, but not return a file.
  
## 1.1.0

### Features
* (**EXPERIMENTAL**) `pre-deploy-check` has been added as a new command.
  This command can be run at any time. 
  It will scan the director and any staged tiles
  in an Ops Manager environment for invalid or missing properties.
  It displays these errors in a list format 
  for the user to manually (or automatedly) update the configuration.
  This command will also return an `exit status 1`;
  this command can be a gatekeeper in CI 
  before running an `apply-changes`
* `download-product` will now include the `product-version` in `download-file.json`
  Thanks to @vchrisb for the PR on issue [#360](https://github.com/pivotal-cf/om/issues/360)

### Bug Fixes
* Extra values passed in the env file 
  will now fail if they are not recognized properties.
  This closes issue [#258](https://github.com/pivotal-cf/om/issues/258)
* Non-string environment variables can now be read and passed as strings to Ops Manager.
  For example, if your environment variable (`OM_NAME`) is set to `"123"` (with quotes escaped),
  it will be evaluated in your config file with the quotes.
  
    Given `config.yml`
    ```yaml
    value: ((NAME))
    ```
    
    `om interpolate -c config.yml --vars-env OM`
    
    Will evaluate to:
    ```yaml
      value: "123"
    ```
  This closes issue [#352](https://github.com/pivotal-cf/om/issues/352)
* the file outputted by `download-product`
  will now use the `product-name` as defined 
  in the downloaded-product, 
  _not_ from the Pivnet slug.
  This fixes a mismatch between the two
  as documented in issue [#351](https://github.com/pivotal-cf/om/issues/351)
* `bosh-env` will now set `BOSH_ALL_PROXY` without a trailing slash
  if one is provided.
  This closes issue [#350](https://github.com/pivotal-cf/om/issues/350) 

## 1.0.0

### Breaking Changes
* `om` will now follow conventional Semantic Versioning,
  with breaking changes in major bumps,
  non-breaking changes for minor bumps,
  and bug fixes for patches.
* `delete-installation` now has a force flag. 
  The flag is required to run this command quietly, as it was working before.
  The reason behind this is
  it was easy to delete your installation without any confirmation. 
* `staged-director-config` no longer supports `--include-credentials`
  this functionality has been replaced by `--no-redact`.
  This can be paired with `--include-placeholders`
  to return a interpolate-able config
  with all the available secrets from a running OpsMan.
  This closes issue #356. 
  The OpsMan API changed so that IAAS Configurations
  were redacted at the API level. 

### Features
* new command `diagnostic-report`
  returns the full OpsMan diagnostic report
  which holds general information about the
  targeted OpsMan's state.
  Documentation on the report's payload
  can be found [here.](https://docs.pivotal.io/pivotalcf/2-2/opsman-api/#diagnostic-report)
* `om interpolate` now can take input from stdin.
  This can be used in conjunction with the new
  `diagnostic-report` command to extract
  a specific section or value
  from the report, simply by using the pipe operator. For example,
  ```bash
  om -e env.yml diagnostic-report | om interpolate --path /versions
  ```
  This will return the `versions` block of the json payload:
  ```yaml
  installation_schema_version: "2.6"
  javascript_migrations_version: v1
  metadata_version: "2.6"
  release_version: 2.6.0-build.77
  ```
* `staged-director-config` now checks
  `int`s and `bool`s when filtering secrets
* `configure-director` and `staged-director` now support `iaas-configurations`.
  This allows OpsManager 2.2+ to have multiple IAASes configured.
  Please see the API documentation for your version of OpsMan for what IAASes are supported.
  
  If you are using `iaas_configuration` in your `properties-configuration` and use `iaas-configurations`
  you'll receive an error message that only one method of configuration can be used. 

## 0.57.0

### Features
* new command `assign-multi-stemcell` supports the OpsMan 2.6+.
  This allows multiple stemcells to be assgined to a single product.
  For example, for product `foo`,
  you could assign Ubuntu Trusty 3586.96 and Windows 2019 2019.2,
  using the command, `om assign-multi-stemcell --product foo --stemcell ubuntu-trusty:3586.96 --stemcell windows2019:2019.2`.
* `upload-stemcell` will not upload the same stemcell (unless using `--force`) for OpsMan 2.6+.
  The API has changed that list the stemcells associated with a product.
  This command is still backwards compatible with OpsMan 2.1+, just has logic specific for 2.6+.

## NOTES
* https://github.com/graymeta/stow/issues/197 has been merged! This should make `om` `go get`-able again.

## 0.56.0

### Breaking Changes
* the `upload-product` flag `--sha256` has been changed to `--shasum`. `upload-stemcell`
  used the `--shasum` flag, and this change adds parity between the two. Using 
  `--shasum` instead of `--sha256` also future-proofs the flag when sha256 is no longer the
  de facto way of defining shasums.

### Features
* `download-product` now supports skipping ssl validation when specifying `--pivnet-disable-ssl`
* `download-product` ensures sha sum checking when downloading the file from Pivotal Network
* `upload-stemcell` now supports a `--config`(`-c`) flag to define all command line arguments
   in a config file. This gives `upload-stemcell` feature parity with `upload-product`
* Improved info messaging for `download-product` to explicitly state whether downloading
  from pivnet or S3
 

## 0.55.0

### Features
* configure-director now has the option to `ignore-verifier-warnings`.
  ([PR #338](https://github.com/pivotal-cf/om/pull/338) Thanks @Logiraptor!)
  This is an _advanced_ feature
  that should only be used if the user knows how to disable verifiers in OpsManager.
  This flag will only disable verifiers for configure-director,
  and will not disable the warnings for apply-changes.
* There's now a shell-completion script;
  see the readme for details.
* We have totally replaced the code and behavior
  of the _experimental_ `config-template` command.
  It now contains the bones of the [tile-config-generator](https://github.com/pivotalservices/tile-config-generator).
  We expect to further refine
  (and make breaking changes to) this command in future releases.

## 0.54.0

### Breaking Changes
* download-product's prefix format and behavior has changed.
  - the prefix format is now `[example-product,1.2.3]original-filename.pivotal`.
  - the prefix is added to _all_ product files if `s3-bucket` is set in the config when downloading from Pivnet.

### Features
* download-product now supports downloading stemcells from S3, too.
* download-product allows use of an instance iam account when `s3-auth-method: iam` is set.
* apply-changes now has the ability to define errands via a config file when running (as a one-off errand run).
  The [apply-changes readme](https://github.com/pivotal-cf/om/docs/apply-changes/README.md) details how this 
  config file should look.
* pending-changes now supports a `--check` flag, that will return an exit code 0(pass) or 1(fail) when running the command, 
  to allow you to fail in CI if there are pending changes in the deployment. 
* download-product will now create a config file (`assign-stemcell.yml`) that can be fed into `assign-stemcell`. It will have the appropriate
  format with the information it received from download-product


### Bug Fixes
* when trying to delete a product on Ops Manager during a selective deploy (`apply-changes --product-name tile`),
  OpsManager would fail to `apply-changes` due to a change to the version string for 2.5 (would include the build
  number). A change was made to the info service to accept the new semver formatting as well as the old 
  versioning. 
* upload-product (among other things) is no longer sensitive to subdirectories in tile metadata directories
* to support 2.5, @jplebre and @edwardecook submitted a PR to support new semver versioning for 
  OpsManager in addition to supporting the current versioning format.
  
### WARNING

To anyone who is having go install fail, it will fail until graymeta/stow#199 is merged.

Here is the error you are probably seeing.

```
$ go install
# github.com/pivotal-cf/om/commands
commands/s3_client.go:62:3: undefined: s3.ConfigV2Signing
```
to work around, you can include `om` in your project without using `go get` or `go install`. you will need to add the following to your `go.mod`:
```
replace github.com/graymeta/stow => github.com/jtarchie/stow v0.0.0-20190209005554-0bff39424d5b
```

## 0.53.0 

### Bug Fixes

* `download-product` would panic if the product was already downloaded and you asked for a stemcell. This has been fixed to behave appropriately

### WARNING

The behavior of `download-product` in this release is not final. Please hold off on using this feature until a release without this warning.

## 0.52.0
### Breaking changes
* `download-product` will now enforce a prefix of `{product-slug}-{semver-version}` when downloading from pivnet. The original
  filename is preserved after the prefix. If the original filename already matches the intended format, there will be no
  change. Any regexes that strictly enforce the original filename at the _beginning_ of the regex will be broken. Please
  update accordingly. This change was done in order to encourage tile teams to change their file names to be more consistent. 
  Ops Manager itself has already agreed to implement this change in newer versions. 

### Features
* add support for the `selected_option` field when calling `staged-config` to have better support for selectors.
  * this support also extends to `configure-product`, which will accept both `selected_option` and `option_value` as
  the machine readable value. 
* `download-product` now has support for downloading from an external s3 compatible blobstore using the `--blobstore s3`
  flag. 
* `staged-director-config` now supports a `no-redact` flag that will return all of the credentials from an Ops Manager
  director, if the user has proper permissions to do so. It is recommended to use the admin user. 
  
### WARNING

The behavior of `download-product` in this release is not final. Please hold off on using this feature until a release without this warning.

## 0.51.0 

### Features

* `import-installation` provides validation on the installation file to ensure
  * it exists
  * it is a valid zip file
  * it contains the `installation.yml` artifact indicative of an exported installation
  
### Bug Fixes

* Fixed typo in `configure-director` vmextensions

## 0.50.0

### Breaking changes

`configure-director` and `staged-director-config` now include a `properties-configuration`.

  The following keys have recently been removed from the top level configuration: director-configuration, iaas-configuration, security-configuration, syslog-configuration.
  
  To fix this error, move the above keys under 'properties-configuration' and change their dashes to underscores.
  
  The old configuration file would contain the keys at the top level.

```yaml
director-configuration: {}
iaas-configuration: {}
network-assignment: {}
networks-configuration: {}
resource-configuration: {}
security-configuration: {}
syslog-configuration: {}
vmextensions-configuration: []
```

  They'll need to be moved to the new 'properties-configuration', with their dashes turn to underscore.
  For example, 'director-configuration' becomes 'director_configuration'.
  The new configration file will look like.

```yaml
az-configuration: {}
network-assignment: {}
networks-configuration: {}
properties-configuration:
  director_configuration: {}
  security_configuration: {}
  syslog_configuration: {}
  iaas_configuration: {}
  dns_configuration: {}
resource-configuration: {}
vmextensions-configuration: []
```
### Features

* The package manager has been migrated from `dep` to `go mod`. It now requires golang 1.11.4+. For information on go modules usage, see the [golang wiki](https://github.com/golang/go/wiki/Modules).

### Bug Fixes

* `import-installation` will now retry 3 times (it uses the polling interval configuration) if it suspects that nginx has not come up yet. This fixes an issue with opsman if you tried to import an installation with a custom SSL Cert for opsman.
* When using `configure-product` on opsman 2.1, it would fail because the completeness check does not work. To disable add the field `validate-config-complete: false` to your config file.
* fix the nil pointer dereference issue in `staged-products` when `om` cannot reach OpsManager

## 0.49.0

### Features

* `download-product` supports grabbing for a version via a regular expression.
  Using `--product-version-regex` sorts the versions returned by semver and
  returns the highest matching version to the regex. The sort ignores non-semver
  version numbers -- similar to the pivnet resource behaviour.
* `download-product` no longer requires `download-stemcell` to be set when specifying `stemcell-iaas`. It is there for backwards compatibility, but it is a no-op.
* added more copy for the help message of `bosh-env`
* fix documentation for `vm-extensions` usage

## 0.48.0

### Features

* Increased the default connect-timeout from `5` seconds to `10`. This should alleviate reliability issues some have seen in CI.

* Adds several commands (`delete-ssl-certificate`, `ssl-certificate`, `update-ssl-certificate`) around managing the Ops Manager VM's TLS cert. These new commands are courtesy of a PR, and we're still tinkering a bit (especially in terms of how they communicate in the case of a default cert, given that the Ops Manager API doesn't even bother returning a cert in that case).
  There should be a fast-to-follow release with these commands more polished; if we'd planned better we might have marked these as experimental until we were done tinkering with them, but we don't see any reason to delay releasing this time.

## 0.47.0

### Breaking changes

* `stage-product` & `configure-product` & `configure-director`: Now errors if `apply-changes` is running. [a3ebd5241d2aba3b93ec642255e0b9c11686d996]

### Features

* `configure-ldap-authentication`: add the command to configure ldap auth during initial setup

### Bug Fixes

* `assign-stemcell`: fix a message format

## 0.46.0

### Breaking changes

* download-product now outputs `product_path`, `product_slug`, `stemcell_path`, and `stemcell_version` instead
  of just `product` and `stemcell`. This will help compatability with `assign-stemcell`.

## 0.45.0

### Breaking changes

* removed individual configuration flags for configure-director \[[commit](https://github.com/pivotal-cf/om/commit/669eca466ca364e4d7597330e5600a013ab9ffe3)\]
* removed individual configuration flags for configure-product \[[commit](https://github.com/pivotal-cf/om/commit/040651b211b8985879337c86357727546099c46e)\]

### Features

* add more intelligent timeouts for commands
* fail fast if a key is not defined in configuration files for configure-product and configure-director
* add `assign-stemcell` command to associate a specified stemcell to the product

### Bug Fixes

* fix stemcell version check logic in `download-product` command -- stemcells can now be downloaded even if they
don't have a minor version (e.g. version 97)

## 0.44.0

### Bug fixes

* The decryption passphrase check was returning dial timeout errors more frequently. Three HTTP retries were added if dial timeout occurs. [Fixes #283]

## 0.43.0

### Breaking changes

* removed command `configure-bosh`, use command `configure-director` for configuring the bosh directory on OpsMan
* removed command `set-errand-state`, use the `errand-config` with your config with the command `configure-product`

### Features

* add command `download-product`, it can download product and associated stemcell from Pivnet
* add `--path` to command `interpolate` so individual values can be extracted out

### Bug Fixes

* automatic decryption passphrase unlock will only attempt doing so once on the first HTTP call #283
* when using command `configure-product`, collections won't fail when `guid` cannot be associated #274

## 0.42.0

### Breaking changes:
* `config-template` & `staged-config` & `staged-director-config`: pluralize `--include-placeholders` flag
* `import-installation`: removed `decryption-passphrase` from the arguments. Global `decryption-passphrase` flag is required when using this command

### Bug Fixes
* update command documentation to reflect various command flags change.
* `configure-product`: handles collection types correctly by decorate collection with guid
* `staged-director-config`: fix failed api request against azure
* `curl`: close http response body to avoid potential resource leaks

### Features
* `configure-product`: allow `product-name` be read from config file
* `interpolation`: added `--vars-env` support to `interpolation`
* `configure-authentication` & `configure-saml-authentication` & `import-installation`: allow the commandline flag been passed through config file
* `configure-director`: able to add/modify/remove vm extensions
*  `staged-config`: able to get errand state for the product
* `apply-changes`: added `skip-unchanged-products`
* `staged-config`: add `product-name` top-level-key in the returned payload to work better with `configure-product`
* `upload-product`: able to validate `sha256` and `product-version` before uploading
* global: added a `decryption-passphrase` to unlock the opsman vm if it is rebooted (if provided)

## 0.40.0

### Bug Fixes

Fix `tile-metadata` command for some tiles that were failing due to it attempting to parse the metadata directory itself as a file - via @chendrix and @aegershman

## 0.39.0

BACKWARDS INCOMPATIBILITIES:
- `om interpolate` no longer takes `--output-file` flag.
- `om interpolate` takes ops-files with `-o` instead of `--ops`.
- `om --format=json COMMAND` is no longer supported. This flag should not have
  been global as it is only supported on some commands. The flag is now
  supported on certain commands and needs to be called: `om COMMAND
  --format=json`. The commands that output in `table` format will continue to do
  so by default.

FEATURES:
- `om configure-product` accepts ops-files.
