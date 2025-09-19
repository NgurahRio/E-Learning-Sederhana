import { useState } from "react";
import MyCourse from "./MyCourseStudent";
import EnrollCourse from "./EnrollCourse";
import { LogOut } from "lucide-react"; // ✅ ikon logout

export default function StudentDashboard() {
  const [activeTab, setActiveTab] = useState("my");

  const handleLogout = () => {
    localStorage.removeItem("token");
    window.location.href = "/"; // arahkan ke halaman login/home
  };

  return (
    <div className="flex min-h-screen bg-gray-900 text-white">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-800 p-6 flex flex-col justify-between">
        <div>
          <h1 className="text-xl font-bold mb-6">🎓 U Learn</h1>
          <ul>
            <li
              className={`cursor-pointer mb-4 ${
                activeTab === "my" ? "text-yellow-400" : ""
              }`}
              onClick={() => setActiveTab("my")}
            >
              My Course
            </li>
            <li
              className={`cursor-pointer mb-4 ${
                activeTab === "enroll" ? "text-yellow-400" : ""
              }`}
              onClick={() => setActiveTab("enroll")}
            >
              Enroll Course
            </li>
          </ul>
        </div>

        {/* Tombol Logout di bawah */}
        <button
          onClick={handleLogout}
          className="flex items-center gap-2 text-gray-300 hover:text-red-400 transition mt-6"
        >
          <LogOut className="w-5 h-5" /> {/* ikon di kiri */}
          <span>Logout</span>
        </button>
      </aside>

      {/* Content */}
      <main className="flex-1 p-8">
        {activeTab === "my" && <MyCourse />}
        {activeTab === "enroll" && <EnrollCourse />}
      </main>
    </div>
  );
}
