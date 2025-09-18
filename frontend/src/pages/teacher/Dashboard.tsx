import { useEffect, useState } from "react";
import API from "../../lib/api";

type Course = {
  id: number;
  name: string;
};

export default function TeacherDashboard() {
  const [courses, setCourses] = useState<Course[]>([]);
  const [newCourse, setNewCourse] = useState("");
  const [loading, setLoading] = useState(true);

  const fetchCourses = async () => {
    try {
      const token = localStorage.getItem("token");
      const res = await API.get("/teachers/my-courses", {
        headers: { Authorization: `Bearer ${token}` },
      });
      setCourses(res.data);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  const addCourse = async () => {
    try {
      const token = localStorage.getItem("token");
      await API.post(
        "/teachers/course",
        { name: newCourse },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      setNewCourse("");
      fetchCourses();
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
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
      <h1 className="text-2xl font-bold mb-6">👨‍🏫 Teacher Dashboard</h1>

      {/* Form tambah course */}
      <div className="flex gap-2 mb-6">
        <input
          type="text"
          placeholder="Nama course baru"
          value={newCourse}
          onChange={(e) => setNewCourse(e.target.value)}
          className="flex-1 border px-3 py-2 rounded"
        />
        <button
          onClick={addCourse}
          className="px-4 py-2 bg-yellow-500 text-white rounded hover:bg-yellow-600"
        >
          Tambah
        </button>
      </div>

      {courses.length === 0 ? (
        <p className="text-gray-600">Belum ada course yang kamu buat.</p>
      ) : (
        <ul className="space-y-3">
          {courses.map((course) => (
            <li
              key={course.id}
              className="p-4 bg-white shadow rounded flex justify-between"
            >
              <span>{course.name}</span>
              <div className="space-x-2">
                <button className="px-3 py-1 bg-blue-500 text-white text-sm rounded hover:bg-blue-600">
                  Edit
                </button>
                <button className="px-3 py-1 bg-red-500 text-white text-sm rounded hover:bg-red-600">
                  Hapus
                </button>
              </div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
