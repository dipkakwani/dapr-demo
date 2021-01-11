# DAPR Store Anomaly

Dapr currently supports only two consistency levels: eventual and strong consistency, based on quorum configuration.  When using eventual consistency, the store only waits for only one replica. In theory, this could lead to store returning stale data, but in practice most of the state stores perform well even with eventual consistency. 

In the below experiment, we tested what it takes to break an existing Dapr application with a Dapr-supported state store.

### Applications

#### Hello World

https://github.com/dapr/quickstarts/tree/master/hello-world

This is a simple hello world application to which we added an additional check for read-your-writes consistency. We use Cassandra as a state store and show that even with such a simple application, it is possible to violate read-your-writes. See appendix for details on how to reproduce this anomaly.

#### Dapr-store

https://github.com/benc-uk/dapr-store

This is a shopping store application built using Dapr. As with other shopping cart applications, there are few anomalies which exist in such application. A well-known anomaly is that of reappearing of a deleted item. When a user is accessing cart from multiple clients, with concurrently delete and add of an item might lead to the user first observing delete being successful, but in subsequent read, the deleted item reappearing in the cart. This anomaly can also be reproduced using the similar setup as that of the above application.





### Appendix 

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

2. Create cassandra.yaml for Dapr and place it in Dapr config folder (~/.dapr/components/)

   ```yaml
   apiVersion: dapr.io/v1alpha1
   kind: Component
   metadata:
     name: statestore
   spec:
     type: state.cassandra
     metadata:
     - name: hosts
       value: 127.0.0.1,127.0.0.2,127.0.0.3
     - name: consistency
       value: "One"
     - name: replicationFactor
       value: "3"
   ```

   

3. Add rule to partition nodes

   The following rules will drop all outgoing messages to all the three replicas (see https://thelastpickle.com/blog/2015/10/12/partitioning-cassandra-for-fun-and-timeouts.html for more systematic way to do it)

```bash
sudo iptables -A INPUT -p tcp --destination localhost --destination-port 7000 -j DROP
sudo iptables -A INPUT -p tcp --destination 127.0.0.2 --destination-port 7000 -j DROP
sudo iptables -A INPUT -p tcp --destination 127.0.0.3 --destination-port 7000 -j DROP
```

3. Run hello world Dapr app

```bash
dapr run --app-id nodeapp --app-port 3000 --dapr-http-port 3500 node app.js
dapr run --app-id pythonapp python3 app.py
```

