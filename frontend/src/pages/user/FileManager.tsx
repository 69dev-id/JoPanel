import { useState, useEffect } from 'react'
import axios from 'axios'
import { Folder, File, Upload, Trash2, FolderPlus, ArrowLeft } from 'lucide-react'

// Types
interface FileInfo {
    name: string
    size: number
    is_dir: boolean
    mod_time: string
}

export default function FileManager() {
    const [path, setPath] = useState('/')
    const [files, setFiles] = useState<FileInfo[]>([])
    const [loading, setLoading] = useState(false)
    const [token] = useState(localStorage.getItem('token'))

    const fetchFiles = async (currentPath: string) => {
        setLoading(true)
        try {
            const res = await axios.get(`http://localhost:8080/api/user/files/list?path=${currentPath}`, {
                headers: { Authorization: `Bearer ${token}` }
            })
            setFiles(res.data || [])
            setPath(currentPath)
        } catch (err) {
            console.error('Failed to load files', err)
            alert('Failed to load files')
        } finally {
            setLoading(false)
        }
    }

    useEffect(() => {
        fetchFiles('/')
    }, [])

    const handleNavigate = (folderName: string) => {
        const newPath = path === '/' ? `/${folderName}` : `${path}/${folderName}`
        fetchFiles(newPath)
    }

    const handleUp = () => {
        if (path === '/') return
        const newPath = path.substring(0, path.lastIndexOf('/')) || '/'
        fetchFiles(newPath)
    }

    const handleDelete = async (fileName: string) => {
        if (!confirm(`Delete ${fileName}?`)) return
        try {
            const filePath = path === '/' ? `/${fileName}` : `${path}/${fileName}`
            await axios.delete(`http://localhost:8080/api/user/files/delete?path=${filePath}`, {
                headers: { Authorization: `Bearer ${token}` }
            })
            fetchFiles(path)
        } catch (err) {
            alert('Failed to delete')
        }
    }

    const handleMkdir = async () => {
        const name = prompt('New Folder Name:')
        if (!name) return
        try {
            const fullPath = path === '/' ? `/${name}` : `${path}/${name}`
            await axios.post('http://localhost:8080/api/user/files/mkdir', { path: fullPath }, {
                headers: { Authorization: `Bearer ${token}` }
            })
            fetchFiles(path)
        } catch (err) {
            alert('Failed to create folder')
        }
    }

    // TODO: Implement Upload Logic
    const handleUpload = () => {
        alert("Upload feature not fully implemented in UI mock")
    }

    return (
        <div className="p-6">
            <div className="mb-6 flex items-center justify-between">
                <div className="flex items-center gap-4">
                    <button onClick={handleUp} disabled={path === '/'} className="rounded bg-gray-200 p-2 hover:bg-gray-300 disabled:opacity-50">
                        <ArrowLeft className="h-5 w-5" />
                    </button>
                    <h2 className="text-xl font-bold">File Manager: {path}</h2>
                </div>
                <div className="flex gap-2">
                    <button onClick={handleMkdir} className="flex items-center gap-2 rounded bg-green-600 px-4 py-2 text-white hover:bg-green-700">
                        <FolderPlus className="h-4 w-4" /> New Folder
                    </button>
                    <button onClick={handleUpload} className="flex items-center gap-2 rounded bg-blue-600 px-4 py-2 text-white hover:bg-blue-700">
                        <Upload className="h-4 w-4" /> Upload
                    </button>
                </div>
            </div>

            <div className="rounded-lg border bg-white shadow">
                <table className="w-full text-left">
                    <thead className="bg-gray-50 text-sm uppercase text-gray-500">
                        <tr>
                            <th className="px-6 py-3">Type</th>
                            <th className="px-6 py-3">Name</th>
                            <th className="px-6 py-3">Size</th>
                            <th className="px-6 py-3">Actions</th>
                        </tr>
                    </thead>
                    <tbody className="divide-y divide-gray-200">
                        {loading && <tr><td colSpan={4} className="p-4 text-center">Loading...</td></tr>}
                        {!loading && files.length === 0 && <tr><td colSpan={4} className="p-4 text-center">Empty Directory</td></tr>}

                        {files.map((file) => (
                            <tr key={file.name} className="hover:bg-gray-50">
                                <td className="px-6 py-4">
                                    {file.is_dir ? <Folder className="h-5 w-5 text-yellow-500" /> : <File className="h-5 w-5 text-gray-400" />}
                                </td>
                                <td className="px-6 py-4 font-medium text-gray-900">
                                    {file.is_dir ? (
                                        <button onClick={() => handleNavigate(file.name)} className="hover:underline hover:text-blue-600">
                                            {file.name}
                                        </button>
                                    ) : (
                                        file.name
                                    )}
                                </td>
                                <td className="px-6 py-4 text-gray-500">{file.size} B</td>
                                <td className="px-6 py-4">
                                    <button onClick={() => handleDelete(file.name)} className="text-red-500 hover:text-red-700">
                                        <Trash2 className="h-4 w-4" />
                                    </button>
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            </div>
        </div>
    )
}
