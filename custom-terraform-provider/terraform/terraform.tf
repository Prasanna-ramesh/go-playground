terraform {
  required_providers {
    mockupstream = {
      source = "terraform.local/prasanna-ramesh/mockupstream"
    }
  }
}

provider "mockupstream" {
  base_url = "http://localhost:4000"
}
