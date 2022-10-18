resource "azurerm_mssql_server" "challenge_mssql_server" {
  name                         = "challenge-sqlserver"
  resource_group_name          = azurerm_resource_group.challenge_rg.name
  location                     = azurerm_resource_group.challenge_rg.location
  version                      = "12.0"
  administrator_login          = "challengeAdmin"
  administrator_login_password = "@Challenge@2022"
}

resource "azurerm_mssql_database" "challenge_mssql_database" {
  name         = "StarWarsDB"
  server_id    = azurerm_mssql_server.challenge_mssql_server.id
  collation    = "SQL_Latin1_General_CP1_CI_AS"
  license_type = "LicenseIncluded"
  sku_name     = "Basic"
}
