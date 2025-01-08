Rails.application.configure do
  # Settings specified here will take precedence over those in config/application.rb.

  # In the development environment your application's code is reloaded on
  # every request. This slows down response time but is perfect for development
  # since you don't have to restart the web server when you make code changes.
  config.cache_classes = false

  # Do not eager load code on boot.
  config.eager_load = false

  # Show full error reports.
  config.consider_all_requests_local = true

  # Enable/disable caching. By default caching is disabled.
  # Run rails dev:cache to toggle caching.
  if Rails.root.join('tmp', 'caching-dev.txt').exist?
    config.action_controller.perform_caching = true
    config.action_controller.enable_fragment_cache_logging = true

    config.cache_store = :memory_store
    config.public_file_server.headers = {
      'Cache-Control' => "public, max-age=#{2.days.to_i}"
    }
  else
    config.action_controller.perform_caching = false

    config.cache_store = :null_store
  end

  # Store uploaded files on the local file system (see config/storage.yml for options).
  config.active_storage.service = :local

  # Don't care if the mailer can't send.
  config.action_mailer.raise_delivery_errors = false

  config.action_mailer.perform_caching = false

  # Print deprecation notices to the Rails logger.
  config.active_support.deprecation = :log

  # Raise an error on page load if there are pending migrations.
  config.active_record.migration_error = :page_load

  # Highlight code that triggered database queries in logs.
  config.active_record.verbose_query_logs = true

  # Debug mode disables concatenation and preprocessing of assets.
  # This option may cause significant delays in view rendering with a large
  # number of complex assets.
  config.assets.debug = true

  # Suppress logger output for asset requests.
  config.assets.quiet = true

  # Raises error for missing translations.
  # config.action_view.raise_on_missing_translations = true

  # Use an evented file watcher to asynchronously detect changes in source code,
  # routes, locales, etc. This feature depends on the listen gem.
  config.file_watcher = ActiveSupport::EventedFileUpdateChecker

  config.action_mailer.default_url_options = {host:"localhost", port:3000}
end

# cb0qgvn9KExuwTod4+0CEODP2IlvoDVYEfDzveiRB1JROUF9SPv+g39eDLcbLjlV2T31V+Js2x8ecjaEUOJELSxO60HA+lDp2lJUqRBLvsxMj6n+6ILkAk+FMb3UgnpUCpXlxJNu2E8ExgJ9eoynFA0UT7olS98pLJsk+FLj0dE0OKZALIGPk96aTpp8jV7hN2osH5psXNWfb9zMBXn3U/db48JnfEQ8KLiGYkoiPhDSFgyhBxX2PCZ02wY842K3KoUxwiUgR0DvL1ddXcw/vwyW5704ILbWG7JGx0cCRs7jKJL3okm7js5H4MwdIUHEKiWtGMcGumTHV1UZe0+tOJTCqSgL2EJnUa5dzRNaylPTgzY7fo5L0WiDSGgVoQX1WaTK8zvRosdKflquCt8nA8ICNQZ99A+89PrO--Mh69lbLjvYdSIsBs--X/27xp0mxTPplGD6ZOhRJw==