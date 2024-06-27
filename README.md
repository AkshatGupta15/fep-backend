# FEP - Backend

## Steps to run the project

1. Change secret.yml.template to secret.yml and enter the passwords for your email and db.

*NOTE*: Use iitk email address, and just add your username, for e.g. : bmerchant22, not bmerchant22@iitk.ac.in

2. Run ```go mod tidy```

3. Run the project:
``` go run cmd/main.go cmd/auth.go cmd/project.go cmd/student.go```
*NOTE*: Configs are set for /backend as root, so don't run ```go run main.go auth.go``` in cmd/

## Services

Currently, if you run the project, there are the following services running:

1. **Auth**: This service will be running on 8080 port of your machine by default, to change it, you can change the config, it has the following routes:
   a. **/api/auth/sendotp (POST)**: If you hit this route with your user_id in payload (which is your roll no.), you will get an otp on your mail.
   b. **/api/auth/signup (POST)**: If you hit this route with user_id, username and password, you will be signed up i.e. your details will be saved in the postgresql db running on your local
