import { useEffect, useState } from "react";
import API from "../../lib/api";

type Course = {
  id_course: number;
  title: string;
  description: string;
  teacher?: { name: string };
};

export default function MyCourse() {
  const [courses, setCourses] = useState<Course[]>([]);

  useEffect(() => {
    const token = localStorage.getItem("token");
    API.get("/students/my-courses", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => setCourses(res.data))
      .catch(() => alert("Gagal load my courses"));
  }, []);

  return (
    <div>
      <h2 className="text-2xl font-bold mb-4">My Courses</h2>
      {courses.length === 0 ? (
        <p className="text-gray-400">Kamu belum enroll course apapun.</p>
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
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
