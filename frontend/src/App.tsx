import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom'
import Login from './pages/Login'
import AdminDashboard from './pages/admin/Dashboard'
import UserDashboard from './pages/user/Dashboard'
import FileManager from './pages/user/FileManager'

// Simple Auth Guard Mock
const ProtectedRoute = ({ children, role }: { children: JSX.Element, role?: string }) => {
    const token = localStorage.getItem('token')
    const userRole = localStorage.getItem('role')

    if (!token) return <Navigate to="/login" replace />
    if (role && userRole !== role) return <Navigate to="/login" replace />

    return children
}

function App() {
    return (
        <Router>
            <Routes>
                <Route path="/login" element={<Login />} />

                {/* Admin Routes */}
                <Route path="/admin" element={
                    <ProtectedRoute role="admin">
                        <AdminDashboard />
                    </ProtectedRoute>
                } />

                {/* User Routes */}
                <Route path="/panel" element={
                    <ProtectedRoute>
                        <UserDashboard />
                    </ProtectedRoute>
                } />

                <Route path="/panel/files" element={
                    <ProtectedRoute>
                        <FileManager />
                    </ProtectedRoute>
                } />

                <Route path="*" element={<Navigate to="/login" replace />} />
            </Routes>
        </Router>
    )
}

export default App
