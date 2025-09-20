import { useEffect, useState } from "react";
import { UserCircle } from "lucide-react"; // ✅ import icon
import API from "../../lib/api";

type Course = {
  id_course: number;
  title: string;
  description: string;
  teacher?: { name: string };
};

export default function MyCourseTeacher() {
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    API.get("/teachers/my-courses", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => setCourses(res.data))
      .catch(() => alert("Gagal load courses"));
  }, []);

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">My Courses</h2>
      {courses.length === 0 ? (
        <p className="text-gray-400">Kamu belum membuat course apapun.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {courses.map((c) => (
            <div
              key={c.id_course}
              className="bg-gray-800 p-4 rounded-lg shadow"
            >
              <h3 className="text-lg font-bold mb-2">{c.title}</h3>
              <p className="text-sm mb-2 text-gray-300">{c.description}</p>
              {c.teacher && (
                <p className="flex items-center text-sm italic text-gray-400">
                  <UserCircle className="w-4 h-4 mr-1" /> {c.teacher.name}
                </p>
              )}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
