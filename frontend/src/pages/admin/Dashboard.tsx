import { useNavigate } from 'react-router-dom'

export default function AdminDashboard() {
    const navigate = useNavigate()
    const user = JSON.parse(localStorage.getItem('user') || '{}')

    const handleLogout = () => {
        localStorage.clear()
        navigate('/login')
    }

    return (
        <div className="flex h-screen bg-gray-100">
            {/* Sidebar */}
            <aside className="w-64 bg-gray-900 text-white">
                <div className="p-4 text-xl font-bold">JoPanel Admin</div>
                <nav className="mt-4 px-2">
                    <a href="#" className="block rounded bg-gray-800 px-4 py-2 text-gray-200">Dashboard</a>
                    <a href="#" className="mt-1 block rounded px-4 py-2 hover:bg-gray-800">User Management</a>
                    <a href="#" className="mt-1 block rounded px-4 py-2 hover:bg-gray-800">Packages</a>
                    <a href="#" className="mt-1 block rounded px-4 py-2 hover:bg-gray-800">Server Status</a>
                </nav>
            </aside>

            {/* Main Content */}
            <main className="flex-1 overflow-y-auto">
                <header className="flex items-center justify-between border-b bg-white px-6 py-4 shadow-sm">
                    <h1 className="text-xl font-semibold text-gray-800">Dashboard</h1>
                    <div className="flex items-center gap-4">
                        <span className="text-sm font-medium text-gray-600">Welcome, {user.username}</span>
                        <button
                            onClick={handleLogout}
                            className="text-sm font-medium text-red-600 hover:text-red-800"
                        >
                            Logout
                        </button>
                    </div>
                </header>

                <div className="p-6">
                    <div className="grid grid-cols-1 gap-6 md:grid-cols-4">
                        <div className="rounded-lg bg-white p-6 shadow">
                            <div className="text-gray-500">Total Users</div>
                            <div className="text-3xl font-bold">12</div>
                        </div>
                        <div className="rounded-lg bg-white p-6 shadow">
                            <div className="text-gray-500">Server Load</div>
                            <div className="text-3xl font-bold text-green-600">0.45</div>
                        </div>
                        <div className="rounded-lg bg-white p-6 shadow">
                            <div className="text-gray-500">Disk Used</div>
                            <div className="text-3xl font-bold">24%</div>
                        </div>
                    </div>
                </div>
            </main>
        </div>
    )
}
