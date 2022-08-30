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
