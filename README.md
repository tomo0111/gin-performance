# Gin performance
Gin performance verification repository.

# Docker image
Docker pull
```
$ docker pull tomohito/gin-performance
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 8080:8080 tomohito/gin-performance
```

# Build by myself
Git clone.
```
$ git clone git@github.com:tomoyane/gin-performance.git
```

Docker build.
```
$ docker build -t {username}/gin-performance:latest .
```

Set Environment variable.
```
DB_SOURCE={database source}
```

Docker run.
```
$ docker run -p 8080:8080 {username}/gin-performance
```

# Performance verification sql
```
CREATE TABLE items (
  id INT(11) NOT NULL AUTO_INCREMENT,
  name VARCHAR(128) NOT NULL,
  category VARCHAR(128) NOT NULL,
  created_at DATETIME(6),
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```
