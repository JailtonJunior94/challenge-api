terraform {
  backend "azurerm" {
    resource_group_name  = "terraform-rg"
    storage_account_name = "storageterraformjj"
    container_name       = "challenge-infra"
    key                  = "terraform.tfstate"
  }
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "=2.71.0"
    }
  }
}

provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "challenge_rg" {
  name     = "challenges-rg"
  location = var.location
}
