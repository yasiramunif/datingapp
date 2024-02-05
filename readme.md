### Step to setup
1. install beego
2. install bee tool
3. install github.com/go-sql-driver/mysql
3. bee migrate -driver=mysql -conn="(username):(password)@tcp(127.0.0.1:3306)/dating_app" 
4. generate rsa key in directory keys
- openssl genrsa -out rsakey.pem 2048
- openssl rsa -in rsakey.pem -pubout > rsakey.pem.pub
5. run seeder sql file

### Details
1. /controllers contains the functions that hit by endpoint
2. /models contains the data struct include=ing request and response data
3. /services contains the logic function 
4. /repositories contains the function to get data from database
5. /middleware contains the function to filter the authentication token