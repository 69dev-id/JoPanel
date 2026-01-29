# JoPanel API Endpoints

## Authentication
- `POST /api/auth/login` - Login and get JWT.
- `POST /api/auth/refresh` - Refresh access token.
- `POST /api/auth/logout` - Invalidate token.

## Admin (WHM-like)
### Account Management
- `POST /api/admin/users` - Create new hosting account (Triggers Agent).
- `GET /api/admin/users` - List users.
- `GET /api/admin/users/:id` - Get user details.
- `PUT /api/admin/users/:id` - Update user package/email.
- `DELETE /api/admin/users/:id` - Terminate account (Triggers Agent).
- `POST /api/admin/users/:id/suspend` - Suspend account.
- `POST /api/admin/users/:id/unsuspend` - Unsuspend.
- `POST /api/admin/users/:id/reset-password` - Reset password.
- `POST /api/admin/users/:id/login-as` - Impersonate user.

### Server Status
- `GET /api/admin/system/load` - Get server load/RAM.
- `GET /api/admin/services/status` - Status of Nginx, MySQL, etc.

## User (cPanel-like)
### Dashboard
- `GET /api/user/stats` - Disk usage, BW usage (Real-time).

### File Manager
- `GET /api/user/files/list?path=/` - List directory.
- `GET /api/user/files/content?path=/file.txt` - Read file.
- `POST /api/user/files/upload` - Upload file(s).
- `POST /api/user/files/mkdir` - Create directory.
- `POST /api/user/files/move` - Move/Rename.
- `DELETE /api/user/files/delete` - Delete.
- `POST /api/user/files/chmod` - Change permissions.
- `POST /api/user/files/zip` - Compress.

### Domains
- `GET /api/user/domains` - List domains.
- `POST /api/user/domains` - Add addon domain/subdomain.
- `DELETE /api/user/domains/:id` - Remove domain.

### Databases
- `GET /api/user/databases` - List DBs.
- `POST /api/user/databases` - Create DB.
