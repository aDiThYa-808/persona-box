# PERSONABOX CLIENT

## Project Structure
```
├── src/
│   ├── hooks.server.ts      # Auth middleware (JWT verification)
│   ├── lib/
│   │   ├── assets/          # Icons, images, screenshots
│   │   ├── components/      # Reusable Svelte components
│   │   └── types/           # TypeScript type definitions
│   ├── routes/
│   │   ├── api/             # BFF proxy endpoints
│   │   │   ├── auth/        # Login, logout, user info
│   │   │   ├── chat/        # Chat message handling
│   │   │   ├── personas/    # Persona CRUD operations
│   │   │   └── sessions/    # Chat session & message retrieval
│   │   ├── chat/            # Chat UI pages
│   │   ├── login/           # Login page
│   │   └── +page.svelte     # Landing page
│   └── stores/              # Svelte stores (state management)
└── static/                  # Static assets
```

## Environment Variables

- `PUBLIC_GOOGLE_CLIENT_ID` - Google OAuth client ID
- `PUBLIC_AWS_INVOKE_URL` - Backend API URL
- `PUBLIC_ENVIRONMENT` - `local` or `production`
- `PERSONABOX_SECRET` - JWT secret (matches backend)

## Architecture

The SvelteKit app also acts as a BFF (Backend-for-Frontend), proxying requests between the browser and the Go backend. This allows for:
- HTTP-only cookie management (secure token storage)
- Request transformation and validation
- Simplified client-side code

---

### Open the relevant files to see detailed inline documentation.
### Thank You