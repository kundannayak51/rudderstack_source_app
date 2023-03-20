# rudderstack_source_app

### Project Details

|                          |                        |
|--------------------------|------------------------|
| **Service name**         | rudderstack_source_app |
| **Language / Framework** | Golang/Gin             |
| **Golang Version**       | 1.18                   |
| **Database**             | MongoDB                |


### Setup in Local
1. Install golang version 1.18
2. Clone the project repo
3. Run `go mod download`
4. Install MongoDB and start the MongoDB server by running `mongod` command


### APIs:
1. API to add source template:
- Endpoint: /source-template
- Method: POST
- cURL:

            
    curl --location 'http://localhost:8080/source-template' \
    --header 'User-Id: 2' \
    --header 'Content-Type: application/json' \
    --data '{
    "type": "source2",
    "fields": {
    "apiKey": {
    "type": "INPUT",
    "label": "API key",
    "regexErrorMessage": "Invalid api key",
    "placeholder": "e.g: 1234asdf",
    "regex": "^[a-zA-Z0-9_]*$",
    "required": true
    },
    "useHTTP": {
    "type": "CHECKBOX",
    "required": false,
    "label": "Enable HTTP"
    },
    "category": {
    "type": "SINGLE_SELECT",
    "label": "Select category",
    "required": true,
    "options": [
    {
    "label": "Android",
    "value": "android"
    },
    {
    "label": "IOS",
    "value": "ios"
    }
    ]
    }
    }
    }'

- Resppnse:


    HTTP status: 201 Created
    Content-Type: application/json

    {
    "success": "successfully added"
    }

2. API to get all source types:
- Endpoint: /all-source-types
- Method: GET
- cURL:


        curl --location 'http://localhost:8080/all-source-types'

- Response:
        

        HTTP status: 200 OK

        [
            {
                "type": "source4"
            },
            {
                "type": "source1"
            },
            {
                "type": "source2"
            }
        ]

3. API to get a source template for a particular source type:
- Endpoint: /source-template/:type
- Method: GET
- cURL:
    
       curl --location 'http://localhost:8080/source-template/source1' \
       --header 'Content-Type: application/json' 
- Response:


        {
            "type": "source1",
            "fields": {
                "apiKey": {
                    "label": "API key",
                    "placeholder": "e.g: 1234asdf",
                    "regex": "^[a-zA-Z0-9_]*$",
                    "regexErrorMessage": "Invalid api key",
                    "required": true,
                    "type": "INPUT"
                },
                "category": {
                    "label": "Select category",
                    "options": [
                        {
                            "label": "Android",
                            "value": "android"
                        },
                        {
                            "label": "IOS",
                            "value": "ios"
                        }
                    ],
                    "required": true,
                    "type": "SINGLE_SELECT"
                },
                "useHTTP": {
                    "label": "Enable HTTP",
                    "required": false,
                    "type": "CHECKBOX"
                }
            }
        }

4. API to add source:
- Endpoint: /source
- Method: POST
- cURL:
    

      curl --location 'http://localhost:8080/source' \
      --header 'User-Id: 2' \
      --header 'Content-Type: application/json' \
      --data '{
      "type": "source1",
      "data": {
      "apiKey": "676f6e",
      "useHTTP": false,
      "category": "android"
      }
      }'

- Response:

      {
          "message": "Source created successfully"
      }      
