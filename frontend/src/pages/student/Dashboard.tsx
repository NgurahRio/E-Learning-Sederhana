import { useEffect, useState } from "react";
import API from "../../lib/api";

type Course = {
  id: number;
  name: string;
  teacher: { name: string };
};

export default function StudentDashboard() {
  const [courses, setCourses] = useState<Course[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        const token = localStorage.getItem("token");
        const res = await API.get("/students/courses", {
          headers: { Authorization: `Bearer ${token}` },
        });
        setCourses(res.data);
      } catch (err) {
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchCourses();
  }, []);

  if (loading) {
    return (
      <div className="flex items-center justify-center min-h-screen bg-gray-100">
        <p>Loading...</p>
      </div>
    );
  }

  return (
    <div className="min-h-screen bg-gray-100 p-8">
      <h1 className="text-2xl font-bold mb-6">🎓 Student Dashboard</h1>

      {courses.length === 0 ? (
        <p className="text-gray-600">Kamu belum mengambil course.</p>
      ) : (
        <ul className="space-y-3">
          {courses.map((course) => (
            <li
              key={course.id}
              className="p-4 bg-white shadow rounded flex justify-between"
            >
              <span>
                {course.name} <br />
                <span className="text-sm text-gray-500">
                  Dosen: {course.teacher?.name}
                </span>
              </span>
              <button className="px-3 py-1 bg-red-500 text-white text-sm rounded hover:bg-red-600">
                Unenroll
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
