import { useState } from "react";
import MyCourse from "./MyCourse";
import EnrollCourse from "./EnrollCourse";

export default function StudentDashboard() {
  const [activeTab, setActiveTab] = useState("my");

  return (
    <div className="flex min-h-screen bg-gray-900 text-white">
      {/* Sidebar */}
      <aside className="w-64 bg-gray-800 p-6">
        <h1 className="text-xl font-bold mb-6">🎓 U Learn</h1>
        <ul>
          <li
            className={`cursor-pointer mb-4 ${activeTab === "my" ? "text-yellow-400" : ""}`}
            onClick={() => setActiveTab("my")}
          >
            My Course
          </li>
          <li
            className={`cursor-pointer mb-4 ${activeTab === "enroll" ? "text-yellow-400" : ""}`}
            onClick={() => setActiveTab("enroll")}
          >
            Enroll Course
          </li>
        </ul>
      </aside>

      {/* Content */}
      <main className="flex-1 p-8">
        {activeTab === "my" && <MyCourse />}
        {activeTab === "enroll" && <EnrollCourse />}
      </main>
    </div>
  );
}
