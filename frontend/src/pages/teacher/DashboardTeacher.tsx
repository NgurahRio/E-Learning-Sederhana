import { useState } from "react";
import MyCourseTeacher from "./MyCourseTeacher";
import AddCourseTeacher from "./AddCourseTeacher";
import ManageCoursesTeacher from "./ManageCoursesTeacher";
import { LogOut } from "lucide-react"; // ✅ samain kayak student

export default function TeacherDashboard() {
  const [activeTab, setActiveTab] = useState("my");

  return (
    <div className="flex min-h-screen bg-gray-900 text-white">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-800 p-6 flex flex-col justify-between">
        <div>
          <h1 className="text-xl font-bold mb-6">👨‍🏫 Teacher Panel</h1>
          <ul>
            <li
              className={`cursor-pointer mb-4 ${
                activeTab === "my" ? "text-yellow-400" : ""
              }`}
              onClick={() => setActiveTab("my")}
            >
              My Courses
            </li>
            <li
              className={`cursor-pointer mb-4 ${
                activeTab === "add" ? "text-yellow-400" : ""
              }`}
              onClick={() => setActiveTab("add")}
            >
              Add Course
            </li>
            <li
              className={`cursor-pointer mb-4 ${
                activeTab === "manage" ? "text-yellow-400" : ""
              }`}
              onClick={() => setActiveTab("manage")}
            >
              Manage Courses
            </li>
          </ul>
        </div>

        {/* tombol logout (samain dengan student) */}
        <button
          onClick={() => {
            localStorage.removeItem("token");
            window.location.href = "/";
          }}
          className="flex items-center gap-2 text-gray-300 hover:text-red-400 transition mt-6"
        >
          <LogOut className="w-5 h-5" /> {/* ✅ ikon sama */}
          <span>Logout</span>
        </button>
      </aside>

      {/* Content */}
      <main className="flex-1 p-8">
        {activeTab === "my" && <MyCourseTeacher />}
        {activeTab === "add" && <AddCourseTeacher />}
        {activeTab === "manage" && <ManageCoursesTeacher />}
      </main>
    </div>
  );
}
