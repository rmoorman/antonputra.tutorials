import Config

config :app, App.Repo,
  username: "postgres",
  password: "postgres",
  hostname: "localhost",
  database: "app_dev",
  stacktrace: true,
  show_sensitive_data_on_connection_error: true,
  pool_size: 10

config :logger, :console, format: "[$level] $message\n"

config :app, App.Router, serve: true, port: 4000
