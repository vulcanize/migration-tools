# migration-tools
Tools for migrating a [v2 ipld-eth-db](https://github.com/vulcanize/ipld-eth-db/releases/tag/v0.2.1) schema DB to a
[v3 ipld-eth-db](https://github.com/vulcanize/ipld-eth-db/releases/tag/v0.3.2) schema DB.

Reads data from a v2 database, transforming them into v3 DB models, and writing them to a v3 database.

Can be configured to work over a subset of the tables, and over specific block ranges.

While processing headers, state, or state accounts it checks for gaps in the data and writes these out to a file. It only checks for
gaps for these tables, as every other table can (in theory) be empty for a given range. Whereas even a block without any transactions will
produce a header and state trie and state account updates (for the miner's reward).

## Usage
`./migration-tools migrate --config={path_to_toml_config_file}`

Example TOML config:

```toml
[migrator]
    ranges = [
        [0, 1000]
    ]
    start = 0 # $MIGRATION_START
    stop = 1000 # $MIGRATION_STOP
    tableNames = [ # $MIGRATION_TABLE_NAMES
        "headers",
        "transactions",
        "storage"
    ]
    workersPerTable = 1 # $MIGRATION_WORKERS_PER_TABLE

[log]
    file = "path/to/log/file" # $LOGRUS_FILE
    level = "info" # $LOGRUS_LEVEL

[v2]
    databaseName = "vulcanize_public_v2" # $V2_DATABASE_NAME
    databaseHostName = "localhost" # $V2_DATABASE_HOSTNAME
    databasePort = "5432" # $V2_DATABASE_PORT
    databaseUser = "postgtres" # $V2_DATABASE_USER
    databasePassword = "" # $V2_DATABASE_PASSWORD
    databaseMaxIdleConns = 50 # $V2_DATABASE_MAX_IDLE_CONNECTIONS
    databaseMaxOpenConns = 100 # $V2_DATABASE_MAX_OPEN_CONNECTIONS
    databaseMaxConnLifetime = 0 # $V2_DATABASE_MAX_CONN_LIFETIME

[v3]
    databaseName = "vulcanize_public_v2" # $V3_DATABASE_NAME
    databaseHostName = "localhost" # $V3_DATABASE_HOSTNAME
    databasePort = "5432" # $V3_DATABASE_PORT
    databaseUser = "postgtres" # $V3_DATABASE_USER
    databasePassword = "" # $V3_DATABASE_PASSWORD
    databaseMaxIdleConns = 50 # $V3_DATABASE_MAX_IDLE_CONNECTIONS
    databaseMaxOpenConns = 100 # $V3_DATABASE_MAX_OPEN_CONNECTIONS
    databaseMaxConnLifetime = 0 # $V3_DATABASE_MAX_CONN_LIFETIME
```

The command can be configured through the linked TOML file as shown above, through ENV variable bindings, or through CLI flags.
The names of the ENV variables are listed in the comments next to their corresponding TOML binding. For a list of the CLI flags, run: 

`./migration-tools migrate --help`

The precedence of configuration is ENV > CLI > TOML. For example, if a parameter is configured by both ENV variable
and in the TOML file, the ENV variable value is the one used.

The tableNames options are:

public.nodes  
eth.header_cids  
eth.uncles_cids  
eth.transaction_cids  
eth.access_list_elements  
eth.receipt_cids  
eth.log_cids  
eth.state_cids  
eth.state_accounts  
eth.storage_cids  
eth.log_cids.repair

public.blocks should be migrated using pg_dump and COPY FROM using a foreign table to handle unique constraint conflicts on INSERT.



