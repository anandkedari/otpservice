# otpservice

Golang service to post and fetch otp


#Fetch SMS for all mobile numbers
```
curl --location --request GET 'http://localhost:8000/sms'
```


#Fetch SMS for specific mobile number
```
curl --location --request GET 'http://localhost:8000/sms/1234567890'
```


#Post new sms to service
- This api deletes old sms and adds new sms for specific number
```

  curl --location --request POST 'http://localhost:8000/sms' \
  --header 'Content-Type: application/json' \
  --data-raw '{
  "smsbody": "OTP1",
  "smssender": "1111",
  "smsrecepient": "1234567890"
  }'
```

#Delete SMS for specific mobile number
```
curl --location --request DELETE 'http://localhost:8000/sms/1234567890' \
--data-raw ''
```



# Docker

#Build Docker image
```
docker build -t otpservice . --no-cache
```

#List of Docker images
```
docker images
```

#Run Docker image
```
docker run -p 8000:8000 -tid otpservice
```

#Go inside docker container
```
docker ps
docker exec -it <container id> bash
```
