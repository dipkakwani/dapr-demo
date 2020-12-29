# DAPR DEMO

### Modifications done in Dapr code

**Dapr**

1. Add WeakIsolationMockDB (dapr/cmd/daprd/main.go)
2. Use local ../components-contrib (dapr/go.mod)
3. CLI brought some changes in parameters names, undo to make it compatible with cli-0.11.0  (dapr/pkg/runtime/cli.go)

**Components-contrib**

1. Add WeakIsolationMockDB to statestores (components-contrib/state/weakisolationmockdb/weakisolationmockdb.go)
2. Temporarily fix error in Cassandra driver [Bind variables cannot be used for table names] by always using table name as default name - dapr.items (components-contrib/state/cassandra/cassandra.go)

### Applications

#### Hello World

https://github.com/dapr/quickstarts/tree/master/hello-world

Modified application to check for read-your-writes consistency in application.

##### Steps to reproduce consistency violation with Cassandra 

1. Set up Cassandra cluster using https://github.com/riptano/ccm

```bash
ccm create test -v 2.2.18 -n 3 -s # create a new test cluster with 3 nodes
ccm node1 cqlsh # login to node1 cqlsh
```

```cassandra
cqlsh> CREATE KEYSPACE dapr WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 3};
cqlsh> use dapr;
cqlsh> CREATE TABLE items (key text PRIMARY KEY, value blob);
```

2. Add rule to partition nodes

   The following rules will drop all outgoing messages to all the three replicas (see https://thelastpickle.com/blog/2015/10/12/partitioning-cassandra-for-fun-and-timeouts.html for more systematic way to do it)

```bash
sudo iptables  -A INPUT -p tcp --destination localhost --destination-port 7000 -j DROP
sudo iptables  -A INPUT -p tcp --destination 127.0.0.2 --destination-port 7000 -j DROP
sudo iptables  -A INPUT -p tcp --destination 127.0.0.3 --destination-port 7000 -j DROP
```

3. Run hello world Dapr app

```bash
dapr run --app-id nodeapp --app-port 3000 --dapr-http-port 3500 node app.js
dapr run --app-id pythonapp python3 app.py
```



**Steps to reproduce consistency violation with WeakIsolationMockDB** 

TODO

#### Dapr-store

https://github.com/benc-uk/dapr-store

1000 iterations

| Assertion\Store | MonkeyDB | Redis |
| --------------- | -------- | ----- |
| Reappear Item   | 21       | xxx   |
| Extra Order     |          |       |

