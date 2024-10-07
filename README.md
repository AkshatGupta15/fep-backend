# FEP - Backend

## Steps to run the project
 Make sure Postgress is installed 
 Run these commands to create databases
 ```
 CREATE DATABASE AUTH;
 CREATE DATABASE PROF;
 CREATE DATABASE STUDENT;
 CREATE DATABASE APPLICATION;
 ```

1. Change secret.yml.template to secret.yml and enter the passwords for your email and db.

*NOTE*: Use iitk email address, and just add your username, for e.g. : bmerchant22, not bmerchant22@iitk.ac.in

2. Run ```go mod tidy```

3. Run the project:
``` go run cmd/main.go cmd/auth.go cmd/project.go cmd/student.go cmd/admin.go cmd/prof.go```
*NOTE*: Configs are set for /backend as root, so don't run ``` go run cmd/main.go cmd/auth.go cmd/project.go cmd/student.go cmd/admin.go``` in cmd/

## Create User:
 1. On postman Enter POST: `http://localhost:8084/api/auth/otp` with payload 
   ```
   {
      "user_id": "akshat23@iitk.ac.in"
   }
   ```
  ### Repeat For RollNo Verification 

   ```
   {
      "user_id": "akshat23@iitk.ac.in"
   }
   ```

   1. Then open psql shell and run `\c auth`
   2. Run `select user_id, otp from otps;`
   3. Copy the otp corresponding to the user_id
   4. Paste it in signup route payload ie, POST `http://localhost:8084/api/auth/signup` 
   Payload 
   ```
   {
    "user_id" : "akshat23@iitk.ac.in",
    "name":"Akshat",
    "roll_no":"230094",
    "roll_no_otp":"otp", //get from step 3
    "password" : "password",
    "user_otp" : "otp" //get from step 3
}
   ```
## Change Role_id : 
   1. Open your psql shell 
   2. Run `\c auth`
   3. Run `UPDATE users set role_id = 100 where user_id = <your_userid> ; `
   
## Services

Currently, if you run the project, there are the following services running:

1. **Auth**: This service will be running on 8084 port of your machine by default, to change it, you can change the config, it has the following routes:
   a. **/api/auth/sendotp (POST)**: If you hit this route with your user_id in payload (which is your roll no.), you will get an otp on your mail.
   b. **/api/auth/signup (POST)**: If you hit this route with user_id, username and password, you will be signed up i.e. your details will be saved in the postgresql db running on your local

2. **Prof Registration(Only via admin side)**: Service running on port : 8081,
   But to access these routes you must have admin role id ie 100, kindly refer how to change role_id
   ### Routes 
    - To register Prof: 

      ```
      POST: /api/admin/prof 
      ```
    - To get all registered Prof: 

      ```
      GET: /api/admin/prof 
      ```
    - To get specific Prof by id: 

      ```
      GET: /api/admin/prof/:pid 
      ```
    - To Update Prof Prof: 

      ```
      PUT: /api/admin/prof 
      ```
    - To Delete Prof: 

      ```
      DELETE: /api/admin/prof/:pid 
      ```


