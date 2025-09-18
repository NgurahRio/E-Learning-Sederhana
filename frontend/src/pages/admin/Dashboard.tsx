import { useEffect, useState } from "react";
import API from "../../lib/api";
import ProtectedRoute from "../../components/ProtectedRoute";

type User = {
  id_user: number;
  name: string;
  email: string;
  role: { roleName: string };
};

export default function AdminDashboard() {
  const [users, setUsers] = useState<User[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    API.get("/users", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => setUsers(res.data))
      .catch(() => alert("Gagal load users"));
  }, []);

  return (
    <ProtectedRoute allowedRoles={[1]}>
      <div className="p-6">
        <h1 className="text-2xl font-bold mb-4">🛠 Admin Dashboard</h1>
        <table className="w-full border">
          <thead className="bg-gray-200">
            <tr>
              <th className="p-2 border">ID</th>
              <th className="p-2 border">Name</th>
              <th className="p-2 border">Email</th>
              <th className="p-2 border">Role</th>
            </tr>
          </thead>
          <tbody>
            {users.map((u) => (
              <tr key={u.id_user}>
                <td className="p-2 border">{u.id_user}</td>
                <td className="p-2 border">{u.name}</td>
                <td className="p-2 border">{u.email}</td>
                <td className="p-2 border">{u.role?.roleName}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </ProtectedRoute>
  );
}
