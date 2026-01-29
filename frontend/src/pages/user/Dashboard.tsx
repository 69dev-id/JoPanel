import { Link, useNavigate } from 'react-router-dom'
import { Folder, Database, Globe, Key, Settings } from 'lucide-react'

export default function UserDashboard() {
    const navigate = useNavigate()
    const user = JSON.parse(localStorage.getItem('user') || '{}')

    const handleLogout = () => {
        localStorage.clear()
        navigate('/login')
    }

    const features = [
        { name: 'File Manager', icon: Folder, link: '/panel/files' },
        { name: 'Databases', icon: Database, link: '#' },
        { name: 'Domains', icon: Globe, link: '#' },
        { name: 'SSH Access', icon: Key, link: '#' },
        { name: 'Settings', icon: Settings, link: '#' },
    ]

    return (
        <div className="min-h-screen bg-gray-50">
            <nav className="bg-white px-6 py-4 shadow">
                <div className="flex items-center justify-between">
                    <div className="text-xl font-bold text-primary">JoPanel User</div>
                    <div className="flex items-center gap-4">
                        <span>{user.username}</span>
                        <button onClick={handleLogout} className="text-red-600">Logout</button>
                    </div>
                </div>
            </nav>

            <div className="container mx-auto p-6">
                <div className="mb-8 grid grid-cols-1 gap-6 md:grid-cols-3">
                    <div className="rounded-lg bg-white p-4 shadow">
                        <h3 className="mb-2 font-medium text-gray-500">Disk Usage</h3>
                        <div className="text-2xl font-bold">120 MB / 1000 MB</div>
                        <div className="mt-2 h-2 w-full rounded-full bg-gray-200">
                            <div className="h-2 w-[12%] rounded-full bg-blue-600"></div>
                        </div>
                    </div>
                    <div className="rounded-lg bg-white p-4 shadow">
                        <h3 className="mb-2 font-medium text-gray-500">Bandwidth</h3>
                        <div className="text-2xl font-bold">50 MB / 10 GB</div>
                    </div>
                    <div className="rounded-lg bg-white p-4 shadow">
                        <h3 className="mb-2 font-medium text-gray-500">Domains</h3>
                        <div className="text-2xl font-bold">1 / 5</div>
                    </div>
                </div>

                <h2 className="mb-4 text-lg font-semibold text-gray-800">Tools</h2>
                <div className="grid grid-cols-2 gap-4 md:grid-cols-4 lg:grid-cols-6">
                    {features.map((item) => (
                        <Link
                            key={item.name}
                            to={item.link}
                            className="flex flex-col items-center justify-center rounded-lg bg-white p-6 shadow transition hover:bg-gray-50 hover:shadow-md"
                        >
                            <item.icon className="mb-3 h-8 w-8 text-blue-600" />
                            <span className="text-sm font-medium text-gray-700">{item.name}</span>
                        </Link>
                    ))}
                </div>
            </div>
        </div>
    )
}
