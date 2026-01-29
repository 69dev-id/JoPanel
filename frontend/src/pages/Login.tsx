import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import axios from 'axios'

export default function Login() {
    const [username, setUsername] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')
    const navigate = useNavigate()

    const handleLogin = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            // In real app, use env var for API URL
            const res = await axios.post('http://localhost:8080/api/auth/login', { username, password })

            localStorage.setItem('token', res.data.access_token)
            localStorage.setItem('role', res.data.user.role)
            localStorage.setItem('user', JSON.stringify(res.data.user))

            if (res.data.user.role === 'admin') {
                navigate('/admin')
            } else {
                navigate('/panel')
            }
        } catch (err: any) {
            setError('Invalid credentials or server error')
        }
    }

    return (
        <div className="flex min-h-screen items-center justify-center bg-gray-100">
            <div className="w-full max-w-sm rounded-lg border bg-white p-8 shadow-lg">
                <h2 className="mb-6 text-center text-2xl font-bold text-gray-800">JoPanel Login</h2>

                {error && <div className="mb-4 rounded bg-red-100 p-2 text-center text-sm text-red-600">{error}</div>}

                <form onSubmit={handleLogin} className="space-y-4">
                    <div>
                        <label className="mb-1 block text-sm font-medium text-gray-700">Username</label>
                        <input
                            type="text"
                            className="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:outline-none"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            required
                        />
                    </div>
                    <div>
                        <label className="mb-1 block text-sm font-medium text-gray-700">Password</label>
                        <input
                            type="password"
                            className="w-full rounded border border-gray-300 p-2 focus:border-blue-500 focus:outline-none"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            required
                        />
                    </div>
                    <button
                        type="submit"
                        className="w-full rounded bg-blue-600 py-2 font-bold text-white hover:bg-blue-700 transition"
                    >
                        Login
                    </button>
                </form>
            </div>
        </div>
    )
}
