## File Processing

1. Clone the repository
2. Start Redis server by calling command: 
   - *redis-server*
3. Now run local server on terminal by calling: 
   - *go run cmd/main.go*
4. Open Postman or any other application to run API (POST request):
    - *localhost:8080/upload*
    - In body attach the file to be uploaded.
    - Sample Response for successful upload
   ```
    {
       "message": "File uploaded and processed successfully"
    }
   ```
5. To retrieve the data of the uploaded file call (GET request):
    - *localhost:8080/file/{filename}
    - Sample Response for successful request
    ```
    {
       "filename": "sample.txt",
       "file_size": 23
    }
   ```
6. You can enable redis cli and check the result:
   - *redis-cli* 
   - GET {file_name} for ex: GET sample.txt
   - Which returns the size of the file.
   - Sample response:
        "23"