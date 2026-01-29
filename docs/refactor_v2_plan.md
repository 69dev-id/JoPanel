# V2 System-Centric Refactoring Plan

## Goals
Refactor JoPanel to use Linux Users and Filesystem Configuration as the Source of Truth, minimizing Database dependency.

## 1. File Structure Definition
**Location:** `/etc/jopanel/users/<username>/`
- `account.json` (Identity)
- `limits.json` (Quotas)
- `domains.json` (Routing)
- `services.json` (Services config)
- `ssh.json` (SSH Security)
- `meta.json` (Audit/Versioning)

## 2. Component Changes

### Backend (API)
- **Repo Layer:**
    - Create `UserRepositoryFS` (Filesystem) implementation.
    - Keep `UserRepositoryDB` only for indexing/caching (List/Search).
- **Services:**
    - `CreateUser`: Call Agent -> Write Config -> Create Linux User -> Sync to DB.
    - `UpdateUser`: Lock -> Write Config -> Agent Apply -> Unlock -> Sync DB.
    - `GetStats`: Read directly from `limits.json` + Real-time system checks (quota cmd).

### Agent (The Heavy Lifter)
- **New Handlers:**
    - `CreateUserFS`: Validates, creates dirs, writes JSONs, runs useradd.
    - `UpdateUserConfig`: Atomic Write + Apply System Changes.
    - `GetConfig`: Read JSONs safely.

### Database
- **Schema Update:**
    - Remove `package_id`, `status` columns (move to `account.json`).
    - Keep `users` table mainly for `id`, `username`, `password_hash` (Panel Auth), `role`.
    - Add `audit_logs` table.

## 3. Migration Plan (V1 -> V2)
1.  **Backup DB**: Dump existing MySQL data.
2.  **Generate Configs**:
    - Iterate all users in DB.
    - Create `/etc/jopanel/users/<username>/` folder.
    - Write default `.json` files based on DB data.
3.  **System Sync**:
    - Check if Linux user exists (`id -u <username>`).
    - If not, create it.
4.  **Verification**: Login to Panel, verify data matches file config.

## 4. Implementation Steps
- [ ] Update `architecture_v2_system_centric.md` with final specs.
- [ ] Create Go Structs for new JSON Schemas (`backend/models/config/`).
- [ ] Implement `Agent` logic for File Operations (Atomic Write, Lock).
- [ ] Implement `Agent` logic for System User primitives (`useradd`, `usermod`).
- [ ] Refactor `Backend` User Service to call Agent for Source of Truth.
