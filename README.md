# Cloud Configuration Microservice

A basic microservice for cloud configurations of your project. Uses MySQL for data storage. Multi-threaded and non-blocking, designed for maximum performance.
### 🔧 Building
```
cd cmd
go build -o cloud
./cloud
```

### ⚙️ Environment
    DSN - MySQL Connection Data (Example: log:pass@/dbname?parseTime=true)

### 💴 Metrics
You can see statistics on the use of resources by the microservice at `yourserver.com/metrics`

### TODO
- [x] Basic CRUD Operations
- [ ] Security
- [ ] C++ API Client Example
 
### 🔌 API Docs
#### Create config
```
POST /api/v1/save
```
| Parameter  | Type     | Description                                                              |
|:-----------|:---------|:-------------------------------------------------------------------------|
| `username` | `string` | **Required**. Username                                                   |
| `data`     | `string` | **Required**. Any content you wanna save.                                |

#### Get all items
```
POST /api/v1/getAll
```
| Parameter  | Type     | Description              |
|:-----------| :------- |:-------------------------|
| `username` | `string` | **Required**. Username   |

----
#### Get config data
```
POST /api/v1/get
```
| Parameter  | Type     | Description                            |
|:-----------|:---------|:---------------------------------------|
| `username` | `string` | **Required**. Username                 |
| `uid`      | `string` | **Required**. System configuration ID. |
----
#### Save config
```
POST /api/v1/save
```
| Parameter  | Type     | Description                               |
|:-----------|:---------|:------------------------------------------|
| `username` | `string` | **Required**. Username                    |
| `uid`      | `string` | **Required**. System configuration ID.    |
| `data`     | `string` | **Required**. Any content you wanna save. |
----
#### Delete config
```
POST /api/v1/delete
```
| Parameter  | Type     | Description                                                              |
|:-----------|:---------|:-------------------------------------------------------------------------|
| `username` | `string` | **Required**. Username                                                   |
| `uid`      | `string` | **Required**. System configuration ID. You can get it from getAll method |
