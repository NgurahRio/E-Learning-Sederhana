import { useState } from "react";
import API from "../../lib/api";
import type { LoginResponse } from "../../types/auth";
import FormInput from "../../components/FormInput";

export default function AdminLogin() {
  const [form, setForm] = useState({ email: "", password: "" });

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const res = await API.post<LoginResponse>("/users/login", form);
      localStorage.setItem("token", res.data.token);

      if (res.data.user.role_id !== 1) {
        alert("Akun ini bukan Admin!");
        return;
      }

      window.location.href = "/admin/dashboard";
    } catch (err: any) {
      alert(err.response?.data?.error || "Login gagal");
    }
  };

  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-100">
      <form onSubmit={handleSubmit} className="bg-white p-6 rounded-lg shadow-md w-96">
        <h1 className="text-2xl font-bold text-center mb-6">🛠 Admin Login</h1>

        <FormInput
          label="Email"
          type="email"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.target.value })}
        />

        <FormInput
          label="Password"
          type="password"
          value={form.password}
          onChange={(e) => setForm({ ...form, password: e.target.value })}
        />

        <button className="w-full bg-red-600 text-white py-2 rounded hover:bg-red-700 transition">
          Login
        </button>
      </form>
    </div>
  );
}
