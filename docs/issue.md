# Weak Isolation Mock DB Proposal

Hello team,

We would like to propose the addition of a new state store that is meant for testing of dapr applications. The store implementation is called [Weak-Isolatation-Mock-DB](https://github.com/microsoft/weak-isolation-mock-db). It has been designed to test applications against _weak behaviors_ of databases, such as those arising when using _eventual consistency_ instead of _strong consistency_. Real-world databases only [rarely](http://www.news.cs.nyu.edu/~jinyang/ds-reading/facebookmeasure.pdf) exhibit weak behaviors, although they do happen, which makes it very difficult for application developers to test against worst-case scenarios. Weak-Isolatation-Mock-DB can generate these worst-case scenarios with a much higher probability. It works as follows: on each read, it computes the set of return values possible under eventual consistency, and then randomly returns a value from this set. We make an argument below showcasing why Weak-Isolatation-Mock-DB is useful. 

### Applications

#### Hello World [[Github](https://github.com/dapr/quickstarts/tree/master/hello-world)]

Eventual consistency does not guarantee _read-your-writes_. That is, if you write a value and then immediately do a read, it need not return the value just written (even in the absence of any concurrent write).

Does it happen in practice? Yes, it does. We modified the hello-world app to check for _read-your-writes_ then connected it to Cassandra with a special setup, and observed a violation of read-your-writes. See appendix for more details. 

#### Dapr-Store [[Github](https://github.com/benc-uk/dapr-store)]

This is a shopping store application built using Dapr. We highlight an _anomaly_ in the application where a deleted iteam reappears in the shopping cart when using eventual consistency. Consider the case when a user is accessing their shopping cart from multiple clients, deleting an item in one session and adding the item in the second session. What can happen here in the following. When the user first looks at their cart, they see the delete having succeeded. After a refresh, reading the cart again, the deleted item comes back and there are two items in the cart this time! The figure below illustrates this example. 

![Shopping Cart Example](shopping_cart_example.png)

This anomaly is valid behavior under eventual consistency. It is, however, difficult to reproduce with a real world database. We confirmed this by running it with Redis and Cassandra state stores and did not observe any violation within reasonable amount of time. However, this anomaly can be reproduced using a modified Cassandra setup where we delibrately injected faults at the right place and time in the Cassandra cluster. 

Weak-Isolatation-Mock-DB makes testing the app against such corner-cases of eventual consistency easy. Just run the dapr application against Weak-Isolatation-Mock-DB. Each run of the app will observe different (randomly chosen) values on store reads. For this example, we catch the anomaly in roughly 20 iterations (< 2 seconds). 

##### Fixing Shopping Cart Anomaly

Once the application developer knows that a violation exist, they have to fix the application. The item reappearing anomaly can be fixed in multiple ways, we describe one possible solution below:

* Use ETags to read and write - In case of concurrent writes within a replica, only one write should go through.
* ETags alone is not sufficient to solve the issue, the writes done by the AddItem and DeleteItem have to be strongly consistent. 

We implemented the above solution and then did not observe any violation when testing with Weak-Isolatation-Mock-DB.

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

