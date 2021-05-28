Migrator
========

Database migration tool using dbmate. Created to be used in the Kubernetes context as a tool to run in an init container to migrate schema files before the actual applications starts. **Note:** Only works with postgresql at the moment.

Configure using environment variables:

| Name                | Description                                                                       |
|---------------------|-----------------------------------------------------------------------------------|
| `MIGRATOR_DB_ADDR`  | Address of the database e.g `mypghost:5432`                                       |
| `MIGRATOR_DB_USER`  | User to access the database                                                       |
| `MIGRATOR_DB_PASS`  | Password to access the database                                                   |
| `MIGRATOR_DB`       | Database to use                                                                   |
| `MIGRATOR_SSL_MODE` | Whether SSL should be enabled or disabled. Choose between `disable` and `enable`. |
