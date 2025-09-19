import { useEffect, useState } from "react";
import API from "../../lib/api";

type Course = {
  id_course: number;
  title: string;
  description: string;
  teacher?: { name: string };
};

export default function EnrollCourse() {
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    API.get("/students/all-courses", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => setCourses(res.data))
      .catch(() => alert("Gagal load courses"));
  }, []);

  const handleEnroll = async (id: number) => {
    const token = localStorage.getItem("token");
    try {
      await API.post(
        "/students/enroll",
        { course_id: id },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert("Berhasil enroll!");
    } catch {
      alert("Enroll gagal");
    }
  };

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">Enroll New Course</h2>
      {courses.length === 0 ? (
        <p className="text-gray-400">Belum ada course tersedia.</p>
      ) : (
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {courses.map((c) => (
            <div key={c.id_course} className="bg-gray-800 p-4 rounded-lg shadow">
              <h3 className="text-lg font-bold mb-2">{c.title}</h3>
              <p className="text-sm mb-2">{c.description}</p>
              {c.teacher && (
                <p className="text-sm italic text-gray-300">
                  by {c.teacher.name}
                </p>
              )}
              <button
                onClick={() => handleEnroll(c.id_course)}
                className="mt-3 w-full bg-yellow-500 py-2 rounded hover:bg-yellow-600 transition"
              >
                Enroll
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
