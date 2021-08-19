connection "googleworkspace" {
  plugin = "googleworkspace"
  
  # You may connect to google workspace using more than one option
  # 1. To authenticate using domain-wide delegation, specify service account credential file, and the user email for impersonation
  # credential_file         = "/Users/subhajit/Downloads/test-01-319203-9fa05a847c16.json"
  # impersonated_user_email = "steampipe@turbot-dev.com"


  # 2. To authenticate OAuth 2.0 client, specify client secret file
  # client_secret_file      = "/Users/subhajit/Downloads/client_secret_737900959689-sgmqnbhukh51g9ge1cch7b1dujmcku0d.apps.googleusercontent.com.json"
}