
Tasks :
1. Run simple WEB server with std lib usage only
endpoint GET /random
input - query parameter min and max - integers
response random number between given min and max parameters inclusively
if POST request is made - return error with "wrong http method"

2. Using Gorilla (https://github.com/gorilla/mux)
3. REST
4. Create GET /coins endpoint using Chi (https://github.com/go-chi/chi)

   GET /coins
   POST /coins
   GET /coins/234/pages/25
   PATCH /coins/234/pages/25
   PUT /coins/234/pages/25
   DELETE /coins/234/pages/25

   DELETE /coins - not allowed
   PATCH /coins - not allowed
   PUT /coins - not allowed

