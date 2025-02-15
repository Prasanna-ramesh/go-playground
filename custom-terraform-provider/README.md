# Custom terraform provider

This repository contains a custom Terraform provider for [mock upstream server](./mock_upstream_server) using [plugin framework](https://developer.hashicorp.com/terraform/plugin/framework).
With this provider, user resource can be managed. For more details on user resource, refer [here](./terraform_provider/provider/resource_user.go)


## Local installation of provider

In order to test the provider locally, Terraform allows the plugin developers to use local build of the provider by setting a `dev_overrides` block in a configuration file called `.terraformrc`. This block overrides all other configured installation methods.

This section mainly focuses on creating build of the provider and configuring Terraform to use it and following are the details. If you are stuck in any of the steps, you can refer [here](https://developer.hashicorp.com/terraform/tutorials/providers-plugin-framework/providers-plugin-framework-provider#prepare-terraform-for-local-provider-install)


### Running mockserver locally

- From the root directory, run the following
```shell
make run-mock-upstream
```

### Terraform configuration

- Terraform plugins are based in Go. Follow [this instruction](https://go.dev/doc/install) for Go installation.
- After installing, find the `GOBIN` path where Go installs the binaries. Run the following to find the `GOBIN` path
```sh
go env GOBIN
```
- If the above command does not return any path, `GOBIN` path is set to `/Users/<Username>/go/bin`. Note this path as this will be used in the next steps.
- Create a new file called `.terraformrc` in your home directory `~`, then paste the following by replacing `<PATH>` with `GOBIN` path
```
provider_installation {

  dev_overrides {
      "terraform.local/Prasanna-ramesh/mockupstream" = "<GOBIN-PATH>"
  }
 
}
```
- With this step, the Terraform is configured to use the local build of the provider

### Installing Terraform provider

- Navigate to the root folder and run the following\
```shell
make install-mock-provider
```

### Terraform

- Once above three steps are completed, the user resource can be managed using Terraform
- For Terraform code, refer [here](./terraform)

>[!Note]
> This is only a playground. So bugs are expected.
