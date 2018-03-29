CREATE KEYSPACE IF NOT EXISTS monastery WITH REPLICATION = {'class': 'org.apache.cassandra.locator.SimpleStrategy', 'replication_factor': '1'};

CREATE TABLE IF NOT EXISTS monastery.files (
    name varchar,
    path frozen<list<varchar>>,
    tags set<varchar>,
    metadata map<varchar, varchar>,
    created timestamp,
    updated timestamp,
    PRIMARY KEY (name, path)
);