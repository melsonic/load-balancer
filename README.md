### Load-Balancer in Golang

#### What is a Load-Balancer
A load balancer is a application that sits between external client(internet) and internal servers to manage traffic between the servers.

#### What are the benifits of using a Load-Balancer
- It helps to reduce application downtime.
- It adds a extra security layer to out application.
- It makes the application more scalable since it adapts the application to handle large amound of traffic.

#### How to Load Balance?
- There are different algorithms for Load Balancing. In broader terms the algorithms can be distributed into two categories.
    - Static Load Balancing : In this method, the load-balancer doesn't care weather a server is running or not, it just sends the traffic towards that server.
    - Dynamic Load Balancing : In this method, the load-balancer takes the server health into account, and makes smart decisions.

#### Algorithms under Static Load Balancing
- Round Robin Algorithm : In this algorithm, the load balancer sends traffic to all the available servers periodically in a round robin fashion.
- IP Hash Algorithm : In this algorith, the load balancer hashes the client ip address and based on the result it selects the target server.

#### Algorithms under Dynamic Load Balancing
- Least Active Connection Algorithm : Here the load balancer directs traffic towards that server which has the least active connection.
- Least Response Time Algorithm : Here the load balancer counts the response time as the decisive parameter for picking the target server.

#### Application Overview
This application contains the following items
- 4 servers
- 1 Load Balancer

The load balancer creates a reverseProxy to direct traffic towards the target server. For now, round robin algorithm is included in this application.

#### Future work
- Add more algorithms.
- Add a property to check health of the available servers.
