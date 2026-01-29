# JoPanel UI Wireframes

## 1. Authentication
**Path:** `/login`
- **Center Card:**
  - Logo "JoPanel"
  - Inputs: Username, Password
  - Button: "Login"
  - Link: "Forgot Password?"

## 2. Admin Dashboard (WHM-like)
**Path:** `/admin`
**Layout:** Sidebar (Left), Header (Top), Content (Main)
- **Sidebar:** Users, Packages, Server Status, Settings, Logout.
- **Dashboard Home:**
  - Cards: Total Users, Server Load, Disk Usage, Network Traffic.
  - Recent Activities Table.

### 2.1 User Management
**Path:** `/admin/users`
- **Datatable:** ID, Username, Domain, Package, Status, Actions (Edit, Suspend, Delete, Login As).
- **Button:** "Create New Account" (Modal or Slide-over).
    - Inputs: Domain, Username, Password, Email, Package Select.

## 3. User Dashboard (cPanel-like)
**Path:** `/panel`
**Layout:** Top Navbar (User Info, Logout), Sidebar (Groups: Files, DB, Domains, etc.), Content.

### 3.1 General Information (Sidebar or Right Column)
- Current User: `archozali`
- Primary Domain: `example.com`
- Home: `/home/archozali`
- IP: `192.168.1.1`

### 3.2 Statistics (Sidebar or Widget)
- Disk Usage: [====..] 40%
- Bandwidth: [=.....] 10%
- SQL Usage: 500MB

### 3.3 Main Sections (Grid Icon Layout)
- **Files:** File Manager, Disk Usage, FTP Accounts.
- **Databases:** MySQL Databases, phpMyAdmin.
- **Domains:** Site Publisher, Addon Domains, Subdomains, DNS Zone Editor.
- **Email:** Email Accounts, Forwarders.
- **Security:** SSH Access, SSL/TLS, 2FA.
- **Software:** MultiPHP Manager, NodeJS Selector.

### 3.4 File Manager UI
**Path:** `/panel/file-manager`
- **Layout:** Three panes (Tree View Left, File List Center, Actions Top).
- **Toolbar:** New File, New Folder, Upload, Download, Rename, Edit, Delete, Permissions.
- **File List:** Icon, Name, Size, Type, Permissions, Last Modified.
- **Code Editor (Monaco):** Full screen overlay or tab.
