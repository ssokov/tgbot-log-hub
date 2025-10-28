# project service name
NAME := tgbot-log-hub

# local database connection settings
TEST_PGDATABASE ?= test-tgbot-log-hub

PGDATABASE ?= tgbot-log-hub
PGHOST ?= localhost
PGPORT ?= 5432
PGUSER ?= mikhail
PGPASSWORD ?= mikhail

# add -race to GOFLAGS if RACE=1
RACE=0
