# tinyplm
Lightweight Cloud Native PLM Running with MongoDB, Redis, Go, React

![Screenshot 2022-08-30 at 10 46 04 AM](https://user-images.githubusercontent.com/741952/187355727-9f739fab-3f45-4846-8759-b01e04bfabd7.png)


### Commands

    # start detached
    docker-compose up -d

    # Stop services only
    docker-compose stop

    # Stop and remove containers, networks..
    docker-compose down 

    # Down and remove volumes
    docker-compose down --volumes 

    # Down and remove images
    docker-compose down --rmi <all|local> 


| Feature      | Description | GitHub Link     |
| :---        |    :----:   |          ---: |
| Rest API   | Text        | https://github.com/gin-gonic/gin     |
| Config Mgmt      | Title       | https://github.com/spf13/viper  |
| Database   | Text        | https://github.com/mongodb/mongo-go-driver      |
| Redis Cache   | Text        | https://github.com/go-redis/redis      |
| Validation   | Text        | https://github.com/go-playground/validator      |
| Health Check   | Text        | https://codewithyury.com/how-to-create-health-check-for-restful-microservice-in-golang/      |
| Readiness   | Text        | And more      |
| Liveliness   | Text        | And more      |
| Prometheus   | Text        | https://github.com/penglongli/gin-metrics      |
| Open API 3   | Text        | And more      |
| CORS   | Text        | And more      |
| JWT Auth   | Text        | And more      |
| Graceful Shutdown   | Text        | And more      |
| Logging   | Text        | And more      |


Explore
    
    https://github.com/memphisdev/memphis-broker
    nats

