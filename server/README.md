# PERSONABOX BACKEND

## Project Structure
```
├── cmd/server/main.go    # Server entry point
├── internal/
│   ├── handlers/         # HTTP request handlers
│   ├── jwtx/             # JWT validation
│   ├── dynamodbx/        # DynamoDB operations
│   │   └── models/       # Models for the tables
│   ├── httpx/            # HTTP helpers
│   ├── openaiapapter/    # OpenAI operations  
│   └── middleware/       # CORS
├── deploy.sh             # Doploy using AWS CLI
```

## Environment Variables

- `GOOGLE_CLIENT_ID`  - Google OAuth client ID
- `PERSONABOX_SECRET` - JWT secret
- `OPENAI_SECRET_KEY` - OpenAI secret key
- `AWS_REGION`        - AWS region to create DynamoDB client

## API Endpoints

**Auth**
- `POST /auth/google` - Google OAuth login, returns JWT access token
- `GET /user` - Get current user info (requires auth)

**Personas**
- `POST /personas` - Create new AI persona (requires auth)
- `GET /personas` - List all user's personas (requires auth)
- `DELETE /personas/{id}` - Delete persona by ID (requires auth)

**Chat**
- `POST /chat` - Creates new chat session if needed. Sends user message and returns AI response. Updates chat session after each message. (requires auth)

**Sessions**
- `GET /sessions/{id}` - Get chat sessions for a persona (requires auth)
- `DELETE /sessions/{pid}/{sid}` - Delete chat session (pid=PersonaID, sid=SessionID) (requires auth)

**Messages**
- `GET /messages/{id}` - Get all messages for a session (requires auth)



## Deployment
Deployed on AWS Lambda with API Gateway.

`./deploy.sh` Creates a new go build and deploys using AWS CLI

---

### Open the relevant files to see detailed inline documentation.

### Thank You